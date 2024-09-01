package web

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=30"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required,max=100"`
	Phone    string `json:"phone" validate:"required,min=10"`
	Photo    string `json:"photo"`
}
