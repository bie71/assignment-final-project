package usecase

import (
	entity "assigment-final-project/domain/entity/coupons"
	entity2 "assigment-final-project/domain/entity/transactions"
	coupon "assigment-final-project/domain/repository/coupons"
	transactionItems "assigment-final-project/domain/repository/transactions"
	"assigment-final-project/helper"
	mysql_connection "assigment-final-project/internal/config/database/mysql"
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type TransactionServiceImpl struct {
	repoTransaction      transactionItems.TransactionRepo
	repoTransactionItems transactionItems.TransactionItemsRepo
	repoCoupons          coupon.CouponsRepo
	repoInitialCoupon    coupon.CouponsPrefix
	validation           *validator.Validate
}

func NewTransactionServiceImpl(repoTransaction transactionItems.TransactionRepo,
	repoTransactionItems transactionItems.TransactionItemsRepo,
	repoCoupons coupon.CouponsRepo, repoInitialCoupon coupon.CouponsPrefix,
	validation *validator.Validate) *TransactionServiceImpl {
	return &TransactionServiceImpl{repoTransaction: repoTransaction,
		repoTransactionItems: repoTransactionItems,
		repoCoupons:          repoCoupons, repoInitialCoupon: repoInitialCoupon, validation: validation}
}

func (t *TransactionServiceImpl) AddTransaction(ctx context.Context, transactionRequest *http_request.TransactionRequest) (*http_response.TransactionResult, error) {
	var (
		wg                         = sync.WaitGroup{}
		totalPriceProduct  float64 = 0
		totalAfterDiscount float64 = 0
		discount           float32 = 0
		codeChan                   = make(chan []string)
		listItems                  = make([]*entity2.TransactionItems, 0)
	)

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
		totalAfterDiscount = totalPriceProduct
	}

	if transactionRequest.CouponCode != "" {
		result, err := t.repoCoupons.FindCouponByCustomerIdAndCode(ctx, transactionRequest.CouponCode, transactionRequest.CustomerId)
		helper.PrintIfError(err)
		if result == nil && err != nil {
			return nil, errors.New("coupon code not valid")
		}
		if result.IsUsed() {
			return nil, errors.New("coupon code is used")
		}
		if !result.ExpireDate().After(time.Now()) {
			return nil, errors.New("coupon code is expire")
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			prefixName := strings.Split(result.CouponCode(), "-")[0]
			for _, item := range transactionRequest.PurchaseItems {
				dataJoin, err := t.repoTransaction.GetProductJoinCategory(ctx, item.ProductId)
				prefix, err := t.repoInitialCoupon.FindCouponPrefix(ctx, prefixName, dataJoin.CategoryId.CategoryName())
				helper.PrintIfError(err)
				if prefix != nil {
					discount += float32((dataJoin.Price*item.Quantity)/100) * float32(prefix.Discount())
				}
			}

			totalAfterDiscount = totalPriceProduct - float64(discount)
			if discount != 0 {
				statusCoupon, _ := t.repoCoupons.UpdateStatusCoupon(ctx, transactionRequest.CouponCode, transactionRequest.CustomerId)
				log.Println(statusCoupon)
			}
		}()
		wg.Wait()
	}

	dataEntity := helper.TransactionRequestToEntity(transactionRequest, totalPriceProduct, totalAfterDiscount, discount)
	errTrx := t.repoTransaction.CreateTransaction(ctx, dataEntity)
	if errTrx != nil {
		return nil, errTrx
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, item := range transactionRequest.PurchaseItems {
			listItems = append(listItems, helper.TransactionItemRequestToEntity(item, dataEntity.TransactionId()))
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
		var (
			strList       = make([]string, 0)
			entityCoupons = make([]*entity.Coupons, 0)
		)
		prefixs, err := t.repoInitialCoupon.GetPrefixMinimumPrice(ctx, totalPriceProduct)
		helper.PrintIfError(err)
		for _, prefix := range prefixs {
			couponCode := strings.ToUpper(prefix.PrefixName()) + "-" + helper.RandomString(16)
			coupon := helper.CouponsRequestToEntity(couponCode, transactionRequest.CustomerId, prefix.ExpireDate())
			entityCoupons = append(entityCoupons, coupon)
			strList = append(strList, coupon.CouponCode())
		}
		if len(entityCoupons) != 0 {
			err = t.repoCoupons.InsertCoupons(ctx, entityCoupons)
			helper.PanicIfError(err)
		}
		codeChan <- strList
	}()
	resultCode := <-codeChan
	wg.Wait()

	result := &http_response.TransactionResult{
		TransactionId: dataEntity.TransactionId(),
		CouponCode:    resultCode,
	}
	return result, nil
}
func (t *TransactionServiceImpl) GetTransaction(ctx context.Context, page int) ([]*http_response.TransactionResponse, int, error) {
	var (
		limit, _            = strconv.Atoi(os.Getenv("LIMIT"))
		offset              = limit * (page - 1)
		wg                  = sync.WaitGroup{}
		chanListTransaction = make(chan []*http_response.TransactionResponse)
	)

	transactions, err := t.repoTransaction.GetTransactions(ctx, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	wg.Add(len(transactions))
	defer close(chanListTransaction)
	go func() {
		listTransaction := make([]*http_response.TransactionResponse, 0)
		for _, transaction := range transactions {
			listItem := make([]*http_response.TransactionItemsResponse, 0)
			result, err := t.repoTransaction.GetItemsProduct(ctx, transaction.TransactionId())
			resultCustomer, err := t.repoTransaction.GetTransactionCustomers(ctx, transaction.TransactionId())
			helper.PrintIfError(err)
			for _, itemsProduct := range result {
				if itemsProduct.TransactionId == transaction.TransactionId() {
					listItem = append(listItem, helper.ToTransactionItemsResponse(itemsProduct))
				}
			}
			listTransaction = append(listTransaction, helper.ToTransactionsResponse(transaction, resultCustomer, listItem))

			wg.Done()
		}
		chanListTransaction <- listTransaction
	}()
	resultTransactions := <-chanListTransaction
	wg.Wait()
	rows := helper.CountTotalRows(ctx, mysql_connection.InitMysqlDB(), "transaction")
	return resultTransactions, rows.TotalRows, nil
}

func (t *TransactionServiceImpl) FindTransaction(ctx context.Context, transactionId string) (*http_response.TransactionResponse, error) {
	var (
		wg              = sync.WaitGroup{}
		chanTransaction = make(chan *http_response.TransactionResponse)
	)

	transaction, err := t.repoTransaction.FindTransaction(ctx, transactionId)
	if err != nil || transaction == nil {
		return nil, errors.New("transaction not found")
	}

	wg.Add(1)
	defer close(chanTransaction)
	go func() {
		defer wg.Done()
		listItem := make([]*http_response.TransactionItemsResponse, 0)
		result, err2 := t.repoTransaction.GetItemsProduct(ctx, transaction.TransactionId())
		resultCustomer, err2 := t.repoTransaction.GetTransactionCustomers(ctx, transaction.TransactionId())
		helper.PrintIfError(err2)
		for _, itemsProduct := range result {
			if itemsProduct.TransactionId == transaction.TransactionId() {
				listItem = append(listItem, helper.ToTransactionItemsResponse(itemsProduct))
			}
		}

		chanTransaction <- helper.ToTransactionsResponse(transaction, resultCustomer, listItem)
	}()
	transactionResult := <-chanTransaction
	wg.Wait()

	return transactionResult, nil
}

func (t *TransactionServiceImpl) DeleteTransaction(ctx context.Context, transactionId string) (string, error) {
	result, errfFind := t.repoTransaction.FindTransaction(ctx, transactionId)
	if result == nil && errfFind != nil {
		return "", errors.New("transaction not found")
	}
	err := t.repoTransaction.DeleteTransaction(ctx, transactionId)
	helper.PrintIfError(err)
	return "Success Delete Transaction", nil
}
