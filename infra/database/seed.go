package database

import "code-space-backend-api/infra/database/models"

func SeedTables() {
	createContentType()
	createPaymentType()
	createPaymentStatus()
	createPaymentMethod()
}

func createContentType() {
	contentType := []models.ContentType{
		{
			ID:   1,
			Name: "Video",
		},
	}

	instance.Create(&contentType)
}

func createPaymentType() {
	paymentTypeList := []models.PaymentType{
		{ID: 1, Name: "atm"},
		{ID: 2, Name: "ticket"},
		{ID: 3, Name: "debit_card"},
		{ID: 4, Name: "credit_card"},
		{ID: 5, Name: "voucher_card"},
		{ID: 6, Name: "prepaid_card"},
		{ID: 7, Name: "account_money"},
		{ID: 8, Name: "bank_transfer"},
		{ID: 9, Name: "digital_wallet"},
		{ID: 10, Name: "crypto_transfer"},
		{ID: 11, Name: "digital_currency"},
	}

	instance.Create(&paymentTypeList)
}

func createPaymentStatus() {
	paymentStatusList := []models.PaymentStatus{
		{ID: 1, Name: "approved"},
		{ID: 2, Name: "pending"},
		{ID: 3, Name: "refunded"},
		{ID: 4, Name: "rejected"},
		{ID: 5, Name: "cancelled"},
		{ID: 6, Name: "charged_back"},
		{ID: 7, Name: "authorized"},
		{ID: 8, Name: "in_process"},
		{ID: 9, Name: "in_mediation"},
	}

	instance.Create(&paymentStatusList)
}

func createPaymentMethod() {
	paymentMethodList := []models.PaymentMethod{
		{ID: 1, Name: "pix"},
		{ID: 2, Name: "account_money"},
		{ID: 3, Name: "debin_transfer"},
		{ID: 4, Name: "ted"},
		{ID: 5, Name: "cvu"},
		{ID: 6, Name: "pse"},
	}

	instance.Create(&paymentMethodList)
}
