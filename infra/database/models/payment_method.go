package models

const _PAYMENT_METHOD_TABLE_NAME string = "payment_methods"

type PaymentMethod struct {
	ID      uint    `gorm:"primaryKey"`
	Name    string  `gorm:"unique"`
	Payment Payment `gorm:"foreignKey:MethodID"`
}

func (PaymentMethod) TableName() string {
	return _PAYMENT_METHOD_TABLE_NAME
}
