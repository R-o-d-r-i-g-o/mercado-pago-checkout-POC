package dto

import "time"

type PaymentLinkOutputDTO struct {
	ClientID     string    `json:"client_id"`
	InitPoint    string    `json:"init_point"`
	CollectorID  uint      `json:"collector_id"`
	PreferenceID string    `json:"id"`
	DateCreated  time.Time `json:"date_created"`
}
