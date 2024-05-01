package input

import (
	"code-space-backend-api/app/Payment/entities"
	"time"

	"github.com/go-playground/validator/v10"
)

// Input Mercado Pago Internal Application DTOs
type (
	PaymentLinkInputDTO struct {
		Payer              PayerInputDTO     `json:"payer"`
		Products           []ProductInputDTO `json:"items"`
		BackUrls           BackUrlInputDTO   `json:"back_urls"`
		NotificationUrl    string            `json:"notification_url"     validate:"required"`
		AutoReturn         string            `json:"auto_return"`
		ExpirationDateTo   time.Time         `json:"expiration_date_to"   validate:"required"`
		ExpirationDateFrom time.Time         `json:"expiration_date_from" validate:"required"`
		Expires            bool              `json:"expires"`
	}

	ProductInputDTO struct {
		ID          string  `json:"id"           validate:"required"`
		Title       string  `json:"title"        validate:"required"`
		CourseHash  string  `json:"course_hash"  validate:"required"`
		Description string  `json:"description"  validate:"required"`
		PictureUrl  string  `json:"picture_url"  validate:"required"`
		Quantity    uint    `json:"quantity"     validate:"required,min=1,max=1"`
		CurrencyId  string  `json:"currency_id"  validate:"required"`
		UnitPrice   float64 `json:"unit_price"   validate:"required"`
	}

	PayerInputDTO struct {
		Name        string        `json:"name"  validate:"required"`
		Email       string        `json:"email" validate:"required,email"`
		PhoneNumber PhoneInputDTO `json:"phone" validate:"required"`
	}

	PhoneInputDTO struct {
		Number string `json:"number" validate:"required"`
	}

	BackUrlInputDTO struct {
		Success string `json:"success"  validate:"required"`
		Pending string `json:"pending"`
		Failure string `json:"failure"  validate:"required"`
	}
)

func (p *PaymentLinkInputDTO) ToDomain() entities.Payer {
	var payerDomain entities.Payer = p.Payer.ToDomain()

	for _, product := range p.Products {
		payerDomain.Purchase.Products = append(payerDomain.Purchase.Products, product.ToDomain())
	}

	return payerDomain
}

func (p *PayerInputDTO) ToDomain() entities.Payer {
	return entities.Payer{
		Name:     p.Name,
		Email:    p.Email,
		Phone:    p.PhoneNumber.Number,
		Purchase: entities.Purchase{},
	}
}

func (p *ProductInputDTO) ToDomain() entities.Product {
	return entities.Product{
		Title:       p.Title,
		ProductHash: p.CourseHash,
		Description: p.Description,
		ImageUrl:    p.PictureUrl,
		Quantity:    p.Quantity,
		CurrencyId:  p.CurrencyId,
		UnitPrice:   p.UnitPrice,
		FullPrice:   float64(p.Quantity) * p.UnitPrice,
	}
}

func (p *PaymentLinkInputDTO) ValidateDTO() error {
	return validator.New().Struct(p)
}

func (p *PayerInputDTO) ValidateDTO() error {
	return validator.New().Struct(p)
}

func (p *ProductInputDTO) ValidateDTO() error {
	return validator.New().Struct(p)
}

func (b *BackUrlInputDTO) ValidateDTO() error {
	return validator.New().Struct(b)
}
