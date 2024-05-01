package enums

type PaymentTypeEnum uint

const (
	ATM              PaymentTypeEnum = 1
	TICKET           PaymentTypeEnum = 2
	DEBIT_CARD       PaymentTypeEnum = 3
	CREDIT_CARD      PaymentTypeEnum = 4
	VOUCHER_CARD     PaymentTypeEnum = 5
	PREPAID_CARD     PaymentTypeEnum = 6
	ACCOUNT_MONEY    PaymentTypeEnum = 7
	BANK_TRANSFER    PaymentTypeEnum = 8
	DIGITAL_WALLET   PaymentTypeEnum = 9
	CRYPTO_TRANSFER  PaymentTypeEnum = 10
	DIGITAL_CURRENCY PaymentTypeEnum = 11
)
