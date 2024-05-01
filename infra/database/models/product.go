package models

import "gorm.io/gorm"

const _PRODUCT_TABLE_NAME string = "products"

type Product struct {
	gorm.Model
	Title       string            `gorm:"not null"`
	Description string            `gorm:"not null"`
	CurrencyID  string            `gorm:"not null"`
	Quantity    uint              `gorm:"not null"`
	UnitPrice   float64           `gorm:"not null"`
	CourseID    uint              `gorm:"not null"`
	Purchases   []PurchaseProduct `gorm:"foreignKey:ProductID"`
}

func (Product) TableName() string {
	return _PRODUCT_TABLE_NAME
}
