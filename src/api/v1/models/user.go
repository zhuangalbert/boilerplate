package models

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"primary_key"`
	Username  string     `gorm:"column:username" json:"username"`
	Password  string     `gorm:"column:password" json:"password"`
	Phone     string     `gorm:"column:phone" json:"phone"`
	Email     string     `gorm:"column:email" json:"email"`
	CreatedAt *time.Time `sql:"index" json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `sql:"index" json:"updated_at" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `sql:"index"`
}

func (User) TableName() string {
	return "users"
}
