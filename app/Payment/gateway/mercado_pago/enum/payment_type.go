package enum

type PaymentType string

const (
	ATM              PaymentType = "atm"
	TICKET           PaymentType = "ticket"
	DEBIT_CARD       PaymentType = "debit_card"
	CREDIT_CARD      PaymentType = "credit_card"
	VOUCHER_CARD     PaymentType = "voucher_card"
	PREPAID_CARD     PaymentType = "prepaid_card"
	ACCOUNT_MONEY    PaymentType = "account_money"
	BANK_TRANSFER    PaymentType = "bank_transfer"
	DIGITAL_WALLET   PaymentType = "digital_wallet"
	CRYPTO_TRANSFER  PaymentType = "crypto_transfer"
	DIGITAL_CURRENCY PaymentType = "digital_currency"
)
