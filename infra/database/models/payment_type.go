package models

const _PAYMENT_TYPE_TABLE_NAME string = "payment_types"

type PaymentType struct {
	ID      uint    `gorm:"primaryKey"`
	Name    string  `gorm:"unique"`
	Payment Payment `gorm:"foreignKey:TypeID"`
}

func (PaymentType) TableName() string {
	return _PAYMENT_TYPE_TABLE_NAME
}
