package model

type UserModel struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserSignupModel struct {
	Username string `json:"username" validate:"required,min=4,max=30"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=30"`
}

type UserLoginModel struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
