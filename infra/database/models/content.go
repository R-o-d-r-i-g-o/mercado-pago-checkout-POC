package models

import (
	"code-space-backend-api/infra/database/enums"

	"gorm.io/gorm"
)

const CONTENT_TABLE_NAME string = "contents"

type Content struct {
	gorm.Model
	ModuleID     uint              `gorm:"not null"`
	TypeID       enums.ContentType `gorm:"not null"`
	Hash         string            `gorm:"unique; index; not null"`
	VideoContent VideoContent      `gorm:"foreignKey:ContentID"`
	Progresses   []UserProgress    `gorm:"foreignKey:ContentID"`
	Comments     []Comment         `gorm:"foreignKey:ContentID"`
}

func (Content) TableName() string {
	return CONTENT_TABLE_NAME
}
