package http_request

type CriteriaRequest struct {
	Name string `json:"name" validate:"required"`
}
