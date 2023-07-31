package model

import "github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"

type (
	User struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email"`
	}

	UserSignupInput struct {
		Username string `json:"username" validate:"required,min=4,max=30"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8,max=30"`
	}

	UserLoginInput struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
)

type UserModel interface {
	NewUserFromEntity(userEntity entity.User) User
}

func NewUserFromEntity(userEntity entity.User) User {
	var u User
	u.Email = userEntity.Email
	u.Username = userEntity.Username
	return u
}
func NewUserArrayFromEntities(userEntities []entity.User) []User {
	var users []User

	for _, u := range userEntities {
		users = append(users, NewUserFromEntity(u))
	}

	return users
}
