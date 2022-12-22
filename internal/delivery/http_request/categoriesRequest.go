package http_request

type CategoryRequest struct {
	Name string `json:"name,omitempty" validate:"required"`
}
