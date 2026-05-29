package model


type User struct {
	UserId            int    `json:"id"`
	Username          string `json:"username" validate:"required,min=4,max=20"`
	Email             string `json:"email" validate:"required,email"`
	Password          string `json:"password" validate:"required,min=5,max=15"`
	ProfilePictureUrl string `json:"profilePictureUrl"`
	IsActive          bool
    UserRole          string `json:"userRole"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required,min=4,max=20"`
	Password string `json:"password" validate:"required,min=5,max=15"`
}


type UserRegisterRequest struct {
	Username string `json:"username" validate:"required,min=4,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5,max=15"`
}


type UserResponse struct {
	UserId            int    `json:"userId"`
	Username          string `json:"username" validate:"required,min=4,max=20"`
	ProfilePictureUrl string `json:"profilePictureUrl"`
	UserRole          string `json:"userRole"`
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

type RequestDeleteUser struct {
	Password string `json:"password" validate:"required,min=5,max=10"`
}

type UserFollowingResponse struct {
	UserResponse
	Following bool `json:"following"`
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
		ProfilePictureUrl: u.ProfilePictureUrl,
		UserRole: u.UserRole,
	}
}


