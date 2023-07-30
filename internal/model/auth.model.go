package model

type Auth struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}
