package entity

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID  `gorm:"primaryKey;column:user_id;type:varchar(36)"`
	Username  string     `gorm:"index;column:username;type:varchar(100)"`
	Email     string     `gorm:"column:email"`
	Password  string     `gorm:"column:password;type:varchar(200)"`
	IsActive  bool       `gorm:"column:is_active;type:boolean"`
	UserRoles []UserRole `gorm:"ForeignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (User) TableName() string {
	return "tb_user"
}
