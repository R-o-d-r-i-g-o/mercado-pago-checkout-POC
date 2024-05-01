package repository

import (
	"code-space-backend-api/app/Payment/entities"
	"code-space-backend-api/infra/database/models"

	"gorm.io/gorm"
)

type purchaseRepository struct {
	db *gorm.DB
}

type PurchaseRepository interface {
	CreatePurchaseWithProducts(payer entities.Payer, purchase entities.Purchase) error
}

func NewPurchaseRepository(db *gorm.DB) PurchaseRepository {
	return &purchaseRepository{
		db: db,
	}
}

func (p *purchaseRepository) CreatePurchaseWithProducts(payer entities.Payer, purchase entities.Purchase) error {
	return p.db.Transaction(func(tx *gorm.DB) (err error) {
		var purchaseID uint
		if purchaseID, err = p.CreatePurchase(purchase, payer.Email); err != nil {
			return
		}

		if err = p.CreateProducts(purchaseID, purchase.Products...); err != nil {
			return
		}

		return
	})
}

func (p *purchaseRepository) CreatePurchase(purchase entities.Purchase, payerEmail string) (uint, error) {
	var purchaseModel models.Purchase = purchase.ToModel()
	var payerID uint

	err := p.db.Model(models.User{}).
		Select("id").
		Where("email = ?", payerEmail).
		First(&payerID).
		Error

	if err != nil {
		return purchase.Id, err
	}

	purchaseModel.UserID = payerID

	return purchaseModel.ID, p.db.Model(models.Purchase{}).
		Create(&purchaseModel).
		Error
}

func (p *purchaseRepository) CreateProducts(purchaseID uint, products ...entities.Product) error {
	return p.db.Transaction(func(tx *gorm.DB) (err error) {
		for _, product := range products {
			var productModel models.Product = product.ToModel()
			var courseID uint

			err = p.db.Model(models.Course{}).Select("id").
				Where("hash = ?", product.ProductHash).
				First(&courseID).
				Error

			if err != nil {
				return
			}

			productModel.CourseID = courseID

			err = p.db.
				Model(models.Product{}).
				Create(&productModel).
				Error

			if err != nil {
				return
			}

			var purchaseProduct = models.PurchaseProduct{
				PurchaseID: purchaseID,
				ProductID:  productModel.ID,
			}

			err = p.db.Model(models.PurchaseProduct{}).Create(&purchaseProduct).Error
			if err != nil {
				return
			}
		}
		return
	})
}
