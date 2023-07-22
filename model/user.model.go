package model

type UserModel struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserAuthenticationModel struct {
	Username string `json:"username" validate:"required,min=4,max=30"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=30"`
}
