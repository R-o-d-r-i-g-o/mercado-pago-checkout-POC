package models

import (
	"time"

	"gorm.io/gorm"
)

const USER_PROGRESS_TABLE_NAME string = "user_progress"

type UserProgress struct {
	gorm.Model
	UserID      uint `gorm:"not null"`
	ContentID   uint `gorm:"not null"`
	CompletedAt time.Time
}

func (UserProgress) TableName() string {
	return USER_PROGRESS_TABLE_NAME
}
