package models

import "gorm.io/gorm"

const VIDEO_CONTENT_TABLE_NAME string = "video_content"

type VideoContent struct {
	gorm.Model
	Url         string `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	ContentID   uint   `gorm:"not null"`
}

func (VideoContent) TableName() string {
	return VIDEO_CONTENT_TABLE_NAME
}
