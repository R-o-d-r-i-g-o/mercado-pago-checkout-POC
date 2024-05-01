package models

import (
	"gorm.io/gorm"
)

const PURCHASE_TABLE_NAME string = "purchases"

type Purchase struct {
	gorm.Model
	UserID                  uint              `gorm:"not null"`
	MercadoPagoClientID     string            `gorm:"not null"`
	MercadoPagoCollectorID  uint              `gorm:"not null"`
	MercadoPagoPreferenceId string            `gorm:"not null"`
	Purchases               []PurchaseProduct `gorm:"foreignKey:PurchaseID"`
	Payment                 Payment           `gorm:"foreignKey:PurchaseID"`
}

func (Purchase) TableName() string {
	return PURCHASE_TABLE_NAME
}
