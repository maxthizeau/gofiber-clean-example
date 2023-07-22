package impl

import (
	"context"
	"errors"

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

func (repo *userRepositoryImpl) Create(username string, password string, email string, roles []string) entity.User {
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
		Email:     email,
		IsActive:  true,
		UserRoles: userRoles,
	}

	err := repo.DB.Create(&user).Error
	exception.PanicLogging(err)

	return user

}

func (repo *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (entity.User, error) {

	var userResult entity.User
	result := repo.DB.WithContext(ctx).
		Where("email = ?", email).First(&userResult)

	if result.RowsAffected == 0 {
		return entity.User{}, errors.New("user not found")
	}

	return userResult, nil
}

func (repo *userRepositoryImpl) FindById(ctx context.Context, id uuid.UUID) (entity.User, error) {

	var userResult entity.User
	result := repo.DB.WithContext(ctx).
		Where("user_id = ?", id).First(&userResult)

	if result.RowsAffected == 0 {
		return entity.User{}, errors.New("user not found")
	}

	return userResult, nil
}

func (repo *userRepositoryImpl) FindAll(ctx context.Context) []entity.User {
	var users []entity.User
	repo.DB.Find(&users)
	return users
}
