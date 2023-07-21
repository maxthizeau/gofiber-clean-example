package impl

import (
	"context"

	model "github.com/maxthizeau/gofiber-clean-boilerplate/models"
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
			Password: user.Password,
		})
	}

	if len(users) == 0 {
		return []model.UserModel{}
	}

	return responses
}

func (serv *userServiceImpl) SignUp(ctx context.Context, userModel model.UserModel) model.UserModel {
	roles := []string{"member"}
	serv.UserRepository.Create(userModel.Username, userModel.Password, roles)

	return userModel
}
