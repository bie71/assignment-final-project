package usecase

import (
	entity "assigment-final-project/domain/entity/coupons"
	entity2 "assigment-final-project/domain/entity/transactions"
	coupon "assigment-final-project/domain/repository/coupons"
	customer "assigment-final-project/domain/repository/customers"
	transactionItems "assigment-final-project/domain/repository/transactions"
	"assigment-final-project/helper"
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"context"
	"errors"
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
	validation           *validator.Validate
}

func NewTransactionServiceImpl(repoTransaction transactionItems.TransactionRepo, repoCustomer customer.RepoCustomer, repoTransactionItems transactionItems.TransactionItemsRepo,
	repoCoupons coupon.CouponsRepo, repoInitialCoupon coupon.CouponsPrefix,
	validation *validator.Validate) *TransactionServiceImpl {
	return &TransactionServiceImpl{repoTransaction: repoTransaction, repoCustomer: repoCustomer, repoTransactionItems: repoTransactionItems,
		repoCoupons: repoCoupons, repoInitialCoupon: repoInitialCoupon, validation: validation}
}

func (t *TransactionServiceImpl) AddTransaction(ctx context.Context, transactionRequest *http_request.TransactionRequest) (string, error) {
	wg := sync.WaitGroup{}
	listItems := make([]*entity2.TransactionItems, 0)
	var totalPriceProduct uint64 = 0
	var totalAfterDiscount uint64 = 0
	var discount uint64 = 0

	errValidation := t.validation.Struct(transactionRequest)
	if errValidation != nil {
		return "", errValidation
	}

	for _, item := range transactionRequest.PurchaseItems {
		result, err := t.repoTransaction.GetProductJoinCategory(ctx, item.ProductId)
		log.Println("error GetProductJoinCategory", err)
		if result.Stock < item.Quantity {
			return "", errors.New("product out of stock or product is empty")
		}
		totalPriceProduct += uint64(result.Price * item.Quantity)
	}

	if transactionRequest.CouponCode != "" {
		wg.Add(1)
		result, err := t.repoCoupons.FindCouponByCustomerIdAndCode(ctx, transactionRequest.CouponCode, transactionRequest.CustomerId)
		if result == nil && err != nil {
			log.Println("result repoCoupons", result, "error repoCoupons", err)
			return "", errors.New("coupon code not valid")
		}
		if result.IsUsed() {
			return "", errors.New("coupon code is used")
		}
		if !result.ExpireDate().After(time.Now()) {
			return "", errors.New("coupon code is expire")
		}

		go func() {
			defer wg.Done()
			prefixs, err2 := t.repoInitialCoupon.GetPrefixs(ctx)
			log.Println("error repoInitialCoupon", err2)
			for _, prefix := range prefixs {
				for _, item := range transactionRequest.PurchaseItems {
					result, err := t.repoTransaction.GetProductJoinCategory(ctx, item.ProductId)
					log.Println("error GetProductJoinCategory", err)
					log.Println("stock ", result.Stock)
					if totalPriceProduct >= uint64(prefix.MinimumPrice()) && strings.EqualFold(prefix.Criteria(), result.CategoryId.CategoryName()) {
						t.repoCoupons.UpdateStatusCoupon(ctx, transactionRequest.CouponCode, transactionRequest.CustomerId)
						discount = (totalPriceProduct / 100) * uint64(prefix.Discount())
						totalAfterDiscount = totalPriceProduct - discount
					}
				}
			}
		}()

		wg.Wait()
	}
	entityTransaction := entity2.NewTransaction(&entity2.DTOTransaction{
		TransactionId:      transactionRequest.TransactionId,
		CustomerId:         transactionRequest.CustomerId,
		CouponCode:         transactionRequest.CouponCode,
		TotalPrice:         totalPriceProduct,
		Discount:           discount,
		TotalAfterDiscount: totalAfterDiscount,
		PurchaseDate:       time.Now(),
	})
	errTrx := t.repoTransaction.CreateTransaction(ctx, entityTransaction)
	if errTrx != nil {
		return "", errTrx
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
		return "", errItems
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		prefixs, err := t.repoInitialCoupon.GetPrefixs(ctx)
		log.Println("error GetPrefixs ", err)
		for _, prefix := range prefixs {
			if totalPriceProduct >= uint64(prefix.MinimumPrice()) {
				nameCoupon := strings.ToUpper(prefix.PrefixName()) + "-" + helper.RandomString(16)
				coupon := entity.NewCoupons(&entity.DTOCoupons{
					CouponCode: nameCoupon,
					ExpireDate: prefix.ExpireDate(),
					CustomerId: transactionRequest.CustomerId,
				})
				err := t.repoCoupons.InsertCoupon(ctx, coupon)
				log.Println("error InsertCoupon ", err)
			}

		}
	}()

	wg.Wait()

	return "Transaction Success Id " + transactionRequest.TransactionId, nil
}

func (t *TransactionServiceImpl) FindTransaction(ctx context.Context, transactionId string) (*http_response.TransactionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TransactionServiceImpl) GetTransaction(ctx context.Context) ([]*http_response.TransactionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TransactionServiceImpl) DeleteTransaction(ctx context.Context, transactionId string) (string, error) {
	//TODO implement me
	panic("implement me")
}
