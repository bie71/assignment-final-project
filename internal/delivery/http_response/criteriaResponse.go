package http_response

import entity "assigment-final-project/domain/entity/criteria"

type CriteriaResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func DomainCriteriaToCriteriaResponse(criteria *entity.Criteria) *CriteriaResponse {
	return &CriteriaResponse{
		Id:   criteria.Id(),
		Name: criteria.CriteriaName(),
	}
}

func ListDomainToListCriteriaResponse(list []*entity.Criteria) []*CriteriaResponse {
	listResponse := make([]*CriteriaResponse, 0)
	for _, criteria := range list {
		response := DomainCriteriaToCriteriaResponse(criteria)
		listResponse = append(listResponse, response)
	}
	return listResponse
}
