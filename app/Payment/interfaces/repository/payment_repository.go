package repository

import (
	"code-space-backend-api/app/Payment/gateway/mercado_pago/dto"
	"code-space-backend-api/app/Payment/interfaces/mapper"
	"code-space-backend-api/infra/database/models"

	"gorm.io/gorm"
)

type paymentRepository struct {
	db *gorm.DB
}

type PaymentRepository interface {
	CreatePayment(paymentDTO dto.PaymentSearchOutputDTO) error
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}

func (p *paymentRepository) CreatePayment(paymentDTO dto.PaymentSearchOutputDTO) error {
	var purchaseID uint

	err := p.db.Model(models.Purchase{}).
		Select("id").
		Where("mercado_pago_collector_id = ?", paymentDTO.CollectorID).
		First(&purchaseID).
		Error

	if err != nil {
		return err
	}

	var paymentModel models.Payment = mapper.GatewayPaymentToModel(paymentDTO)
	paymentModel.PurchaseID = purchaseID

	return p.db.
		Model(models.Payment{}).
		Create(&paymentModel).
		Error
}
