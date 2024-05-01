package models

import (
	"code-space-backend-api/infra/database/enums"

	"gorm.io/gorm"
)

const _PAYMENT_TABLE_NAME string = "payments"

type Payment struct {
	gorm.Model
	TypeID               enums.PaymentTypeEnum   `gorm:"not null"`
	StatusID             enums.PaymentStatusEnum `gorm:"not null"`
	MethodID             enums.PaymentMethodEnum `gorm:"not null"`
	PurchaseID           uint                    `gorm:"not null"`
	TransactionAmount    float64                 `gorm:"not null"`
	MercadoPagoPaymentID uint                    `gorm:"not null"`
}

func (Payment) TableName() string {
	return _PAYMENT_TABLE_NAME
}
