package enums

type PaymentMethodEnum uint

const (
	Pix           PaymentMethodEnum = 1
	AccountMoney  PaymentMethodEnum = 2
	DebinTransfer PaymentMethodEnum = 3
	TED           PaymentMethodEnum = 4
	CVU           PaymentMethodEnum = 5
	PSE           PaymentMethodEnum = 6
)
