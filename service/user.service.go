package service

import (
	"context"

	"github.com/maxthizeau/gofiber-clean-boilerplate/model"
)

type UserService interface {
	FindAll(ctx context.Context) (responses []model.UserModel)
	SignUp(ctx context.Context, authModel model.UserSignupModel) model.AuthModel
}
