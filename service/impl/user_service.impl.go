package impl

import (
	"context"

	"github.com/maxthizeau/gofiber-clean-boilerplate/common"
	"github.com/maxthizeau/gofiber-clean-boilerplate/configuration"
	"github.com/maxthizeau/gofiber-clean-boilerplate/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/exception"
	"github.com/maxthizeau/gofiber-clean-boilerplate/helpers"
	"github.com/maxthizeau/gofiber-clean-boilerplate/model"
	"github.com/maxthizeau/gofiber-clean-boilerplate/repository"
	"github.com/maxthizeau/gofiber-clean-boilerplate/service"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	repository.UserRepository
	configuration.Config
}

func NewUserServiceImpl(userRepository *repository.UserRepository, config configuration.Config) service.UserService {
	return &userServiceImpl{UserRepository: *userRepository, Config: config}
}

func (serv *userServiceImpl) generateTokenForUser(user entity.User) string {
	rolesMap := helpers.GetJwtRoleFromUserEntity(user)
	return common.GenerateJwtToken(user.Id, rolesMap, serv.Config)
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

func (serv *userServiceImpl) SignUp(ctx context.Context, authModel model.UserSignupModel) model.AuthModel {

	common.Validate(authModel)
	roles := []string{"ADMINISTRATOR"}

	password, err := bcrypt.GenerateFromPassword([]byte(authModel.Password), 6)
	exception.PanicLogging(err)

	user := serv.UserRepository.Create(authModel.Username, string(password), authModel.Email, roles)

	return model.AuthModel{
		User: model.UserModel{
			Username: user.Username,
			Email:    user.Email,
		},
		Token: serv.generateTokenForUser(user),
	}
}

func (serv *userServiceImpl) Authenticate(ctx context.Context, authModel model.UserLoginModel) model.AuthModel {
	common.Validate(authModel)
	user, err := serv.UserRepository.FindByEmail(ctx, authModel.Email)
	exception.PanicLogging(err)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authModel.Password))
	exception.PanicUnauthorized(err)

	return model.AuthModel{
		User: model.UserModel{
			Username: user.Username,
			Email:    user.Email,
		},
		Token: serv.generateTokenForUser(user),
	}

}
