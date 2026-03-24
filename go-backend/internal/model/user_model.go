package model


type User struct {
	UserId      int `json:"id"`
	Username string `json:"username" validate:"required,min=3,max=14"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5,max=10"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=14"`
	Password string `json:"password" validate:"required,min=5,max=10"`
}


type UserRegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=14"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5,max=10"`
}


func (u UserRegisterRequest) ToUser() User {
	return User {
		Username: u.Username,
		Email: u.Email,
		Password: u.Password,
	}
}