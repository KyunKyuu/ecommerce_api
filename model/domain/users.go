package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID    int    `gorm:"column:user_id;primaryKey;autoIncrement"`
	Name      string `gorm:"column:name"`
	Role      string `gorm:"column:role"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
