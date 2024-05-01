package enum

type PaymentMethod string

const (
	Pix           PaymentMethod = "pix"
	AccountMoney  PaymentMethod = "account_money"
	DebinTransfer PaymentMethod = "debin_transfer"
	TED           PaymentMethod = "ted"
	CVU           PaymentMethod = "cvu"
	PSE           PaymentMethod = "pse"
)
