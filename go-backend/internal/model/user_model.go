package model


type User struct {
	UserId          int    `json:"id"`
	Username        string `json:"username" validate:"required,min=4,max=14"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=5,max=10"`
	ProfileImageUrl string `json:"profileImageUrl"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required,min=4,max=14"`
	Password string `json:"password" validate:"required,min=5,max=10"`
}


type UserRegisterRequest struct {
	Username string `json:"username" validate:"required,min=4,max=14"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5,max=10"`
}


type UserResponse struct {
	UserId          int    `json:"id"`
	Username        string `json:"username" validate:"required,min=4,max=14"`
	ProfileImageUrl string `json:"profileImageUrl"`
}

type RequestUpdateUsername struct {
	Username string `json:"username" validate:"required,min=4,max=14"`
}

type RequestUpdateEmail struct {
	Email    string `json:"email" validate:"required,email"`
	NewEmail string `json:"newEmail" validate:"required,email"`
}

type RequestUpdatePassword struct {
	Password    string `json:"password" validate:"required,min=5,max=10"`
	NewPassword string `json:"newPassword" validate:"required,min=5,max=10"`
}

type ResponseUpdateProfilePicture struct {
	PictureURL string `json:"pictureURL"`
}

func (u UserRegisterRequest) ToUser() User {
	return User {
		Username: u.Username,
		Email: u.Email,
		Password: u.Password,
	}
}

func (u User) ToUserResponse() UserResponse {
	return UserResponse{
		UserId: u.UserId,
		Username: u.Username,
		ProfileImageUrl: u.ProfileImageUrl,
	}
}