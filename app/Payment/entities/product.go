package entities

import "code-space-backend-api/infra/database/models"

type (
	Product struct {
		Title       string
		ProductHash string
		Description string
		ImageUrl    string
		Quantity    uint
		CurrencyId  string
		UnitPrice   float64
		FullPrice   float64
	}
)

func (p Product) WithTitle(title string) Product {
	p.Title = title
	return p
}

func (p Product) WithProductHash(productHash string) Product {
	p.ProductHash = productHash
	return p
}

func (p Product) WithDescription(description string) Product {
	p.Description = description
	return p
}

func (p Product) WithImageUrl(imageUrl string) Product {
	p.ImageUrl = imageUrl
	return p
}

func (p Product) WithQuantity(quantity uint) Product {
	p.Quantity = quantity
	return p
}

func (p Product) WithCurrencyId(currencyId string) Product {
	p.CurrencyId = currencyId
	return p
}

func (p Product) WithBrazilianCurrency() Product {
	p.CurrencyId = "BRL"
	return p
}

func (p Product) WithUnitPrice(unitPrice float64) Product {
	p.UnitPrice = unitPrice
	p.FullPrice = float64(p.Quantity) * p.UnitPrice
	return p
}

func (p *Product) ToModel() models.Product {
	return models.Product{
		Title:       p.Title,
		Description: p.Description,
		CurrencyID:  p.CurrencyId,
		Quantity:    p.Quantity,
		UnitPrice:   p.UnitPrice,
	}
}
