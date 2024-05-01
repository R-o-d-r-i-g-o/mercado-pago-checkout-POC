package models

const _PAYMENT_STATUS_TABLE_NAME string = "payment_status"

type PaymentStatus struct {
	ID      uint    `gorm:"primaryKey"`
	Name    string  `gorm:"unique"`
	Payment Payment `gorm:"foreignKey:StatusID"`
}

func (PaymentStatus) TableName() string {
	return _PAYMENT_STATUS_TABLE_NAME
}
