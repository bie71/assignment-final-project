package usecase

import (
	entity "assigment-final-project/domain/entity/coupons"
	entity2 "assigment-final-project/domain/entity/transactions"
	repository2 "assigment-final-project/domain/repository/categories"
	coupon "assigment-final-project/domain/repository/coupons"
	customer "assigment-final-project/domain/repository/customers"
	repository "assigment-final-project/domain/repository/products"
	transactionItems "assigment-final-project/domain/repository/transactions"
	"assigment-final-project/helper"
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"strings"
	"sync"
	"time"
)

type TransactionServiceImpl struct {
	repoTransaction      transactionItems.TransactionRepo
	repoCustomer         customer.RepoCustomer
	repoTransactionItems transactionItems.TransactionItemsRepo
	repoCoupons          coupon.CouponsRepo
	repoInitialCoupon    coupon.CouponsPrefix
	repoProduct          repository.ProductRepo
	repoCategory         repository2.CategoryRepo
	validation           *validator.Validate
}

func NewTransactionServiceImpl(repoTransaction transactionItems.TransactionRepo, repoCustomer customer.RepoCustomer, repoTransactionItems transactionItems.TransactionItemsRepo,
	repoCoupons coupon.CouponsRepo, repoInitialCoupon coupon.CouponsPrefix, repoProduct repository.ProductRepo, repoCategory repository2.CategoryRepo,
	validation *validator.Validate) *TransactionServiceImpl {
	return &TransactionServiceImpl{repoTransaction: repoTransaction, repoCustomer: repoCustomer, repoTransactionItems: repoTransactionItems,
		repoCoupons: repoCoupons, repoInitialCoupon: repoInitialCoupon, repoProduct: repoProduct, repoCategory: repoCategory,
		validation: validation}
}

func (t *TransactionServiceImpl) AddTransaction(ctx context.Context, transactionRequest *http_request.TransactionRequest) (*http_response.TransactionResult, error) {
	wg := sync.WaitGroup{}
	var totalPriceProduct float64 = 0
	var totalAfterDiscount float64 = 0
	var discount float32 = 0
	codeChan := make(chan []string)
	listItems := make([]*entity2.TransactionItems, 0)

	errValidation := t.validation.Struct(transactionRequest)
	if errValidation != nil {
		return nil, errValidation
	}

	for _, item := range transactionRequest.PurchaseItems {
		result, err := t.repoTransaction.GetProductJoinCategory(ctx, item.ProductId)
		helper.PanicIfError(err)
		if result.Stock < item.Quantity {
			return nil, errors.New("product out of stock or product is empty")
		}
		totalPriceProduct += float64(result.Price * item.Quantity)
	}

	if transactionRequest.CouponCode != "" {
		wg.Add(1)
		result, err := t.repoCoupons.FindCouponByCustomerIdAndCode(ctx, transactionRequest.CouponCode, transactionRequest.CustomerId)
		log.Println(err)
		if result == nil && err != nil {
			return nil, errors.New("coupon code not valid")
		}
		if result.IsUsed() {
			return nil, errors.New("coupon code is used")
		}
		if !result.ExpireDate().After(time.Now()) {
			return nil, errors.New("coupon code is expire")
		}

		go func() {
			defer wg.Done()
			prefixs, err2 := t.repoInitialCoupon.GetPrefixs(ctx)
			helper.PanicIfError(err2)
			for _, prefix := range prefixs {
				for _, item := range transactionRequest.PurchaseItems {
					result, err := t.repoTransaction.GetProductJoinCategory(ctx, item.ProductId)
					log.Println("stock ", result.Stock)
					helper.PanicIfError(err)
					if totalPriceProduct >= float64(prefix.MinimumPrice()) && strings.EqualFold(prefix.Criteria(), result.CategoryId.CategoryName()) {
						_, err := t.repoCoupons.UpdateStatusCoupon(ctx, transactionRequest.CouponCode, transactionRequest.CustomerId)
						helper.PanicIfError(err)
						discount = float32(totalPriceProduct/100) * float32(prefix.Discount())
						totalAfterDiscount = totalPriceProduct - float64(discount)
						fmt.Println(totalAfterDiscount, prefix.Discount(), totalPriceProduct/100)
					}
				}
			}
		}()

		wg.Wait()
	}
	entityTransaction := entity2.NewTransaction(&entity2.DTOTransaction{
		TransactionId:      "transaction-" + helper.RandomString(16),
		CustomerId:         transactionRequest.CustomerId,
		CouponCode:         transactionRequest.CouponCode,
		TotalPrice:         totalPriceProduct,
		Discount:           discount,
		TotalAfterDiscount: totalAfterDiscount,
		PurchaseDate:       time.Now(),
	})
	errTrx := t.repoTransaction.CreateTransaction(ctx, entityTransaction)
	if errTrx != nil {
		return nil, errTrx
	}

	wg.Add(1)

	go func() {
		defer wg.Done()
		for _, item := range transactionRequest.PurchaseItems {
			items := entity2.NewTransactionItems(&entity2.DTOTransactionItems{
				TransactionId: entityTransaction.TransactionId(),
				ProductId:     item.ProductId,
				Quantity:      item.Quantity,
			})
			listItems = append(listItems, items)
		}
	}()

	wg.Wait()
	errItems := t.repoTransactionItems.InsertItems(ctx, listItems)
	if errItems != nil {
		return nil, errItems
	}

	wg.Add(1)
	defer close(codeChan)
	go func() {
		defer wg.Done()
		strList := make([]string, 0)

		prefixs, err := t.repoInitialCoupon.GetPrefixs(ctx)
		helper.PanicIfError(err)
		for _, prefix := range prefixs {
			if totalPriceProduct >= float64(prefix.MinimumPrice()) {
				nameCoupon := strings.ToUpper(prefix.PrefixName()) + "-" + helper.RandomString(16)
				coupon := entity.NewCoupons(&entity.DTOCoupons{
					CouponCode: nameCoupon,
					ExpireDate: prefix.ExpireDate(),
					CustomerId: transactionRequest.CustomerId,
				})
				err := t.repoCoupons.InsertCoupon(ctx, coupon)
				helper.PanicIfError(err)
				strList = append(strList, coupon.CouponCode())
			}
		}
		codeChan <- strList
	}()
	resultCode := <-codeChan
	wg.Wait()

	result := &http_response.TransactionResult{
		TransactionId: entityTransaction.TransactionId(),
		CouponCode:    resultCode,
	}

	return result, nil
}
func (t *TransactionServiceImpl) GetTransaction(ctx context.Context) ([]*http_response.TransactionResponse, error) {
	wg := sync.WaitGroup{}
	chanListItem := make(chan []*http_response.TransactionItemsResponse)
	transactions, err := t.repoTransaction.GetTransactions(ctx)
	if err != nil {
		return nil, err
	}
	wg.Add(1)
	defer close(chanListItem)
	go func() {
		defer wg.Done()
		for _, transaction := range transactions {
			result, err2 := t.repoTransaction.GetItemsProduct(ctx, transaction.TransactionId())
			helper.PanicIfError(err2)
		}
	}()
}

func (t *TransactionServiceImpl) FindTransaction(ctx context.Context, transactionId string) (*http_response.TransactionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TransactionServiceImpl) DeleteTransaction(ctx context.Context, transactionId string) (string, error) {
	//TODO implement me
	panic("implement me")
}
