package entities

import "code-space-backend-api/infra/database/models"

type (
	Purchase struct {
		Id           uint
		Status       string
		ClientID     string
		PaymentID    uint
		CollectorID  uint
		PreferenceId string
		Products     []Product
	}
)

func (p Purchase) WithID(id uint) Purchase {
	p.Id = id
	return p
}

func (p Purchase) WithStatus(status string) Purchase {
	p.Status = status
	return p
}

func (p Purchase) WithClientID(clientId string) Purchase {
	p.ClientID = clientId
	return p
}

func (p Purchase) WithPaymentID(paymentId uint) Purchase {
	p.PaymentID = paymentId
	return p
}

func (p Purchase) WithCollectorID(collectorId uint) Purchase {
	p.CollectorID = collectorId
	return p
}

func (p Purchase) WithPreferenceID(preferenceId string) Purchase {
	p.PreferenceId = preferenceId
	return p
}

func (p Purchase) WithProduct(product ...Product) Purchase {
	p.Products = append(p.Products, product...)
	return p
}

func (p *Purchase) ToModel() models.Purchase {
	return models.Purchase{
		MercadoPagoClientID:     p.ClientID,
		MercadoPagoCollectorID:  p.CollectorID,
		MercadoPagoPreferenceId: p.PreferenceId,
	}
}
