package service

import (
	"context"

	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/model"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/repository"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/common"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{UserRepository: userRepository}

}

func (serv *userService) FindAll(ctx context.Context) (responses []model.UserModel) {
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

func (serv *userService) SignUp(ctx context.Context, authModel model.UserSignupModel) entity.User {

	common.Validate(authModel)
	roles := []string{"ADMINISTRATOR"}

	password, err := bcrypt.GenerateFromPassword([]byte(authModel.Password), 6)
	exception.PanicLogging(err)

	user := serv.UserRepository.Create(authModel.Username, string(password), authModel.Email, roles)

	return user

}

func (serv *userService) Authenticate(ctx context.Context, authModel model.UserLoginModel) entity.User {
	common.Validate(authModel)
	user, err := serv.UserRepository.FindByEmail(ctx, authModel.Email)
	exception.PanicLogging(err)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authModel.Password))
	exception.PanicLogging(err)

	return user

}
