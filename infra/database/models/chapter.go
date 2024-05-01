package models

import "gorm.io/gorm"

const _CHAPTER_TABLE_NAME string = "chapters"

type Chapter struct {
	gorm.Model
	Name     string    `gorm:"not null"`
	Hash     string    `gorm:"unique; index; not null"`
	CourseID uint      `gorm:"not null"`
	Contents []Content `gorm:"foreignKey:ModuleID"`
}

func (Chapter) TableName() string {
	return _CHAPTER_TABLE_NAME
}
