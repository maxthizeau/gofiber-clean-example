package service

import (
	"context"

	model "github.com/maxthizeau/gofiber-clean-boilerplate/models"
)

type UserService interface {
	FindAll(ctx context.Context) (responses []model.UserModel)
	SignUp(ctx context.Context, userModel model.UserModel) model.UserModel
}
