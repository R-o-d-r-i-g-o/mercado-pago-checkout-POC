package models

import (
	"gorm.io/gorm"
)

const USER_TABLE_NAME string = "users"

type User struct {
	gorm.Model
	Name                  string         `gorm:"not null"`
	Hash                  string         `gorm:"unique; index; not null"`
	Phone                 string         `gorm:"not null; unique"`
	Email                 string         `gorm:"unique; index; not null"`
	Password              string         `gorm:"not null"`
	PolicyPrivacyAccepted bool           `gorm:"not null; default:false"`
	UserCourse            UserCourse     `gorm:"foreignKey:UserID"`
	Purchase              Purchase       `gorm:"foreignKey:UserID"`
	Progresses            []UserProgress `gorm:"foreignKey:UserID"`
	Comments              []Comment      `gorm:"foreignKey:AuthorID"`
}

func (User) TableName() string {
	return USER_TABLE_NAME
}
