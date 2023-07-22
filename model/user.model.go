package model

type UserModel struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserAuthenticationModel struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
