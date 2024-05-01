package entities

import "code-space-backend-api/infra/database/models"

type (
	Payer struct {
		Id       uint
		Hash     string
		Name     string
		Email    string
		Phone    string
		Purchase Purchase
	}
)

func (p Payer) WithID(id uint) Payer {
	p.Id = id
	return p
}

func (p Payer) WithHash(hash string) Payer {
	p.Hash = hash
	return p
}

func (p Payer) WithName(name string) Payer {
	p.Name = name
	return p
}

func (p Payer) WithEmail(email string) Payer {
	p.Email = email
	return p
}

func (p Payer) WithPhoneNumber(phone string) Payer {
	p.Phone = phone
	return p
}

func (p Payer) WithPurchase(purchase Purchase) Payer {
	p.Purchase = purchase
	return p
}

func (p *Payer) ToModel() models.User {

	return models.User{
		Name:     p.Name,
		Hash:     p.Hash,
		Phone:    p.Phone,
		Email:    p.Email,
		Purchase: models.Purchase{},
	}
}
