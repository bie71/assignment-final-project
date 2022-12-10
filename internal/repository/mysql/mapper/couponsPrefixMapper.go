package mapper

import (
	entity "assigment-final-project/domain/entity/coupons"
	"assigment-final-project/internal/repository/mysql/models"
)

func DomainCounponsPrefixToModel(domain *entity.CouponsPrefix) *models.CouponsPrefixModel {
	return &models.CouponsPrefixModel{
		Id:           domain.Id(),
		PrefixName:   domain.PrefixName(),
		MinimumPrice: domain.MinimumPrice(),
		Discount:     domain.Discount(),
		ExpireDate:   domain.ExpireDate(),
		Criteria:     domain.Criteria(),
		CreatedAt:    domain.CreatedAt(),
	}
}

func ModelToDomainCounponsPrefix(model *models.CouponsPrefixModel) *entity.CouponsPrefix {
	return entity.NewCouponsPrefix(&entity.DTOCouponsPrefix{
		Id:           model.Id,
		PrefixName:   model.PrefixName,
		MinimumPrice: model.MinimumPrice,
		Discount:     model.Discount,
		ExpireDate:   model.ExpireDate,
		Criteria:     model.Criteria,
		CreatedAt:    model.CreatedAt,
	})
}

func ListModelToListDomainCouponsPrefix(list []*models.CouponsPrefixModel) []*entity.CouponsPrefix {
	listDomain := make([]*entity.CouponsPrefix, 0)
	for _, model := range list {
		prefix := ModelToDomainCounponsPrefix(model)
		listDomain = append(listDomain, prefix)
	}
	return listDomain
}
