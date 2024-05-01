package models

const PURCHASE_PRODUCT_TABLE_NAME string = "purchase_products"

type PurchaseProduct struct {
	PurchaseID uint `gorm:"primaryKey;autoIncrement:false"`
	ProductID  uint `gorm:"primaryKey;autoIncrement:false"`
}

func (PurchaseProduct) TableName() string {
	return PURCHASE_PRODUCT_TABLE_NAME
}
