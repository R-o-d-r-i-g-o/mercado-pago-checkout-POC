package enum

type PaymentStatus string

const (
	APPROVED     PaymentStatus = "approved"
	PENDING      PaymentStatus = "pending"
	REFUNDED     PaymentStatus = "refunded"
	REJECTED     PaymentStatus = "rejected"
	CANCELLED    PaymentStatus = "cancelled"
	CHARGEBACK   PaymentStatus = "charged_back"
	AUTHORIZED   PaymentStatus = "authorized"
	IN_PROCESS   PaymentStatus = "in_process"
	IN_MEDIATION PaymentStatus = "in_mediation"
)
