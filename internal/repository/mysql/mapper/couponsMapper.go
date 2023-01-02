package mapper

import (
	entity "assigment-final-project/domain/entity/coupons"
	"assigment-final-project/internal/repository/mysql/models"
	"github.com/rocketlaunchr/dbq/v2"
)

func DomainCouponsToCouponsModel(domain *entity.Coupons) *models.CouponsModel {
	return &models.CouponsModel{
		Id:         domain.Id(),
		CouponCode: domain.CouponCode(),
		IsUsed:     domain.IsUsed(),
		ExpireDate: domain.ExpireDate(),
		CustomerId: domain.CustomerId(),
	}
}

func ModelCouponsToDomainCoupons(model *models.CouponsModel) *entity.Coupons {
	return entity.NewCoupons(&entity.DTOCoupons{
		Id:         model.Id,
		CouponCode: model.CouponCode,
		IsUsed:     model.IsUsed,
		ExpireDate: model.ExpireDate,
		CustomerId: model.CustomerId,
	})
}

func ListModelToListDomainCoupons(listModel []*models.CouponsModel) []*entity.Coupons {
	listEntity := make([]*entity.Coupons, 0)
	for _, model := range listModel {
		coupons := ModelCouponsToDomainCoupons(model)
		listEntity = append(listEntity, coupons)
	}
	return listEntity
}

func ListEntityCouponToListInterface(listEntity []*entity.Coupons) []interface{} {
	list := make([]interface{}, 0)
	for _, coupons := range listEntity {
		list = append(list, dbq.Struct(DomainCouponsToCouponsModel(coupons)))
	}
	return list
}
