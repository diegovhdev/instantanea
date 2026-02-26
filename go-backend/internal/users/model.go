package users


type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" validate:"required,min=3,max=14"`
	Password string `json:"password" validate:"required,min=5,max=10"`
	Email    string `json:"email" validate:"omitempty,email"`
}

type UserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=14"`
	Password string `json:"password" validate:"required,min=5,max=10"`
}