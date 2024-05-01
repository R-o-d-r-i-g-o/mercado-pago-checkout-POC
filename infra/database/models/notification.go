package models

import (
	"database/sql"

	"gorm.io/gorm"
)

const _NOTIFICATION_TABLE_NAME string = "notifications"

type Notification struct {
	gorm.Model
	Data sql.NullString `gorm:"type:json"`
}
