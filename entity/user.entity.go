package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `gorm:"primaryKey;column:user_id;type:uuid;default:gen_random_uuid()"`
	Username  string    `gorm:"unique_index;index;column:username;type:varchar(100)"`
	Email     string    `gorm:"unique;column:email"`
	Password  string    `gorm:"column:password;type:varchar(200)"`
	IsActive  bool      `gorm:"column:is_active;type:boolean"`
	CreatedAt time.Time
	UserRoles []UserRole `gorm:"ForeignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Games     []*Game    `gorm:"many2many:user_games;"`
}

func (User) TableName() string {
	return "tb_user"
}
