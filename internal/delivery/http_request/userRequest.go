package http_request

type UserRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=200"`
	Username string `json:"username" validate:"required,min=1,max=100"`
	Password string `json:"password" validate:"required,min=1,max=200"`
	UserType string `json:"user_type" validate:"required,eq=owner|eq=employee"`
}

type UserLogin struct {
	Username string `json:"username" validate:"required,min=1,max=100"`
	Password string `json:"password" validate:"required,min=1,max=200"`
}
