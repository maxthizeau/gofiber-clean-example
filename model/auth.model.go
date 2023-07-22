package model

type AuthModel struct {
	User  UserModel `json:"user"`
	Token string    `json:"token"`
}
