package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
	"gorm.io/gorm"
)

type userRepository struct {
	*gorm.DB
}

func NewUserRepository(DB *gorm.DB) *userRepository {
	return &userRepository{
		DB: DB,
	}
}

func (repo *userRepository) Create(username string, password string, email string, roles []string) entity.User {
	var userRoles []entity.UserRole

	if len(roles) != 0 {
		for _, r := range roles {
			userRoles = append(userRoles, entity.UserRole{
				Id:   uuid.New(),
				Role: r,
			})
		}
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

func (repo *userRepository) FindByEmail(ctx context.Context, email string) (entity.User, error) {

	var userResult entity.User
	result := repo.DB.WithContext(ctx).
		Table("tb_user").
		// Joins("JOIN tb_user_role ON tb_user_role.user_id = tb_user.user_id").
		Preload("UserRoles").
		Where("email = ?", email).First(&userResult)

	if result.RowsAffected == 0 {
		return entity.User{}, errors.New("user not found")
	}

	return userResult, nil
}

func (repo *userRepository) FindById(ctx context.Context, id uuid.UUID) (entity.User, error) {

	var userResult entity.User
	result := repo.DB.WithContext(ctx).
		Table("tb_user").
		Joins("JOIN tb_user_role ON tb_user_role.user_id = tb_user.user_id").
		Preload("UserRoles").
		Where("tb_user.user_id = ?", id).First(&userResult)

	if result.RowsAffected == 0 {
		return entity.User{}, errors.New("user not found")
	}

	return userResult, nil
}

func (repo *userRepository) FindAll(ctx context.Context) []entity.User {
	var users []entity.User
	repo.DB.Find(&users)
	return users
}
