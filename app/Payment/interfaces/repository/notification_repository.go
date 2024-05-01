package repository

import (
	"code-space-backend-api/infra/database/models"
	"database/sql"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type notificationRepository struct {
	db *gorm.DB
}

type NotificationRepository interface {
	CreateNotification(notificationJSON string) error
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{
		db: db,
	}
}

func (n *notificationRepository) CreateNotification(notificationJSON string) error {
	if !n.isJSON(notificationJSON) {
		return fmt.Errorf("invalid notification format")
	}

	var notification = models.Notification{
		Data: sql.NullString{
			String: notificationJSON,
			Valid:  true,
		},
	}

	return n.db.
		Model(models.Notification{}).
		Create(&notification).
		Error
}

func (n *notificationRepository) isJSON(payload string) bool {
	var singlePayload map[string]interface{}
	if err := json.Unmarshal([]byte(payload), &singlePayload); err == nil {
		return true
	}

	var multiplePayload []map[string]interface{}
	if err := json.Unmarshal([]byte(payload), &multiplePayload); err == nil {
		return true
	}

	return false
}
