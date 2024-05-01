package models

import (
	"code-space-backend-api/common/uuid"

	"gorm.io/gorm"
)

const COURSE_TABLE_NAME string = "courses"

type Course struct {
	gorm.Model
	Hash        *uuid.V4   `gorm:"unique; index"`
	Name        string     `gorm:"not null"`
	Description string     `gorm:"not null"`
	UserCourse  UserCourse `gorm:"foreignKey:CourseID"`
	Modules     []Chapter  `gorm:"foreignKey:CourseID"`
	Product     Product    `gorm:"foreignKey:CourseID"`
}

func (Course) TableName() string {
	return COURSE_TABLE_NAME
}
