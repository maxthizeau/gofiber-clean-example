package impl

import (
	"context"

	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/exception"
	"github.com/maxthizeau/gofiber-clean-boilerplate/repository"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	*gorm.DB
}

func NewUserRepositoryImpl(DB *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{
		DB: DB,
	}
}

func (repo *userRepositoryImpl) Create(username string, password string, roles []string) entity.User {
	var userRoles []entity.UserRole

	for _, r := range roles {
		userRoles = append(userRoles, entity.UserRole{
			Id:   uuid.New(),
			Role: r,
		})
	}

	user := entity.User{
		Username:  username,
		Password:  password,
		IsActive:  true,
		UserRoles: userRoles,
	}

	err := repo.DB.Create(&user).Error
	exception.PanicLogging(err)

	return user

}

func (repo *userRepositoryImpl) FindAll(ctx context.Context) []entity.User {
	var users []entity.User
	repo.DB.Find(&users)
	return users
}
