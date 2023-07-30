package service

import (
	"context"
	"errors"

	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/helpers"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/model"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/repository"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/auth"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/common"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repository.UserRepository
	*auth.AuthManager
}

func NewUserService(userRepository repository.UserRepository, authManager *auth.AuthManager) *userService {
	return &userService{UserRepository: userRepository, AuthManager: authManager}

}

func (serv *userService) FindAll(ctx context.Context) (responses []entity.User) {
	users := serv.UserRepository.FindAll(ctx)
	responses = append(responses, users...)
	return responses
}

func (serv *userService) SignUp(ctx context.Context, authModel model.UserSignupInput) entity.User {

	common.Validate(authModel)
	roles := []string{"ADMINISTRATOR"}

	password, err := bcrypt.GenerateFromPassword([]byte(authModel.Password), 6)
	exception.PanicLogging(err)

	user := serv.UserRepository.Create(authModel.Username, string(password), authModel.Email, roles)

	return user

}

func (serv *userService) Authenticate(ctx context.Context, authModel model.UserLoginInput) entity.User {
	common.Validate(authModel)
	invalidError := errors.New("invalid credentials")
	user, err := serv.UserRepository.FindByEmail(ctx, authModel.Email)
	if err != nil {
		exception.PanicBadRequest(invalidError)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authModel.Password))
	if err != nil {
		exception.PanicBadRequest(invalidError)
	}

	return user

}

func (serv *userService) FindLoggedUsed(ctx context.Context) entity.User {
	userContext := helpers.GetUserFromContext(ctx, serv.AuthManager)
	user, err := serv.UserRepository.FindById(ctx, userContext.UserId)
	exception.PanicUnauthorized(err)
	return user

}
