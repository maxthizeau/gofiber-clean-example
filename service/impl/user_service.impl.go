package impl

import (
	"context"

	"github.com/maxthizeau/gofiber-clean-boilerplate/common"
	"github.com/maxthizeau/gofiber-clean-boilerplate/model"
	"github.com/maxthizeau/gofiber-clean-boilerplate/repository"
	"github.com/maxthizeau/gofiber-clean-boilerplate/service"
)

type userServiceImpl struct {
	repository.UserRepository
}

func NewUserServiceImpl(userRepository *repository.UserRepository) service.UserService {
	return &userServiceImpl{UserRepository: *userRepository}
}

func (serv *userServiceImpl) FindAll(ctx context.Context) (responses []model.UserModel) {
	users := serv.UserRepository.FindAll(ctx)
	for _, user := range users {
		responses = append(responses, model.UserModel{
			Username: user.Username,
			Email:    user.Email,
		})
	}

	if len(users) == 0 {
		return []model.UserModel{}
	}

	return responses
}

func (serv *userServiceImpl) SignUp(ctx context.Context, authModel model.UserAuthenticationModel) model.AuthModel {

	common.Validate(authModel)
	roles := []string{"member"}
	user := serv.UserRepository.Create(authModel.Username, authModel.Password, authModel.Email, roles)

	return model.AuthModel{
		User: model.UserModel{
			Username: user.Username,
			Email:    user.Email,
		},
		Token: "token todo",
	}
}
