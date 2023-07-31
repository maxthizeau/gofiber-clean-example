package helpers

import (
	"github.com/google/uuid"
	"github.com/maxthizeau/gofiber-clean-boilerplate/internal/entity"
)

func UserIsInArray(arr []entity.User, userId uuid.UUID) bool {
	for _, a := range arr {
		if a.Id == userId {
			return true
		}
	}
	return false
}
