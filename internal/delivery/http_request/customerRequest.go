package http_request

type CustomerRequest struct {
	Name  string `json:"name" validate:"required,min=3,max=200"`
	Phone string `json:"phone" validate:"required,min=4,max=100,numeric"`
}
