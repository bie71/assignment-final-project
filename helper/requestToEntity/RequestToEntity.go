package requestToEntity

import (
	categories "assigment-final-project/domain/entity/categories"
	coupons "assigment-final-project/domain/entity/coupons"
	customer "assigment-final-project/domain/entity/customers"
	product "assigment-final-project/domain/entity/products"
	user "assigment-final-project/domain/entity/users"
	"assigment-final-project/helper"
	"assigment-final-project/internal/delivery/http_request"
	"time"
)

func CategoryRequestToEntity(request *http_request.CategoryRequest, categoryId string) *categories.Categories {
	return categories.NewCategories(&categories.DTOCategories{
		CategoryId: categoryId,
		Name:       request.Name,
	})
}

func CouponPrefixRequestToDomainEntity(request *http_request.CouponsPrefixRequest, expireDate time.Time) *coupons.CouponsPrefix {
	return coupons.NewCouponsPrefix(&coupons.DTOCouponsPrefix{
		PrefixName:   request.PrefixName,
		MinimumPrice: request.MinimumPrice,
		Discount:     request.Discount,
		ExpireDate:   expireDate,
		Criteria:     request.Criteria,
		CreatedAt:    time.Now(),
	})
}

func CustomerRequestToEntity(request *http_request.CustomerRequest, customerId string) *customer.Customers {
	data, _ := customer.NewCustomer(&customer.DTOCustomers{
		CustomerId: customerId,
		Name:       request.Name,
		Contact:    request.Phone,
		CreatedAt:  time.Now(),
	})
	return data
}

func ProductRequestToEntity(request *http_request.ProductsRequest, productId string) *product.Products {
	return product.NewProducts(&product.DTOProducts{
		ProductId:  productId,
		Name:       request.Name,
		Price:      request.Price,
		CategoryId: request.CategoryId,
		Stock:      request.Stock,
	})
}

func UserRequestToEntity(request *http_request.UserRequest, userId string) (*user.Users, error) {
	data, err := user.NewUsers(&user.DTOUsers{
		UserId:    userId,
		Name:      request.Name,
		Username:  request.Username,
		Password:  helper.HashAndSalt([]byte(request.Password)),
		UserType:  request.UserType,
		CreatedAt: time.Now(),
	})
	return data, err
}
