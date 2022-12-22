package http_request

type CustomerRequest struct {
	Name  string `json:"name,omitempty" validate:"required,min=3,max=200"`
	Phone string `json:"phone,omitempty" validate:"required,min=4,max=100,numeric"`
}
