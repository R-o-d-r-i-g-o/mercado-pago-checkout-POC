package enums

type PaymentStatusEnum uint

const (
	APPROVED     PaymentStatusEnum = 1
	PENDING      PaymentStatusEnum = 2
	REFUNDED     PaymentStatusEnum = 3
	REJECTED     PaymentStatusEnum = 4
	CANCELLED    PaymentStatusEnum = 5
	CHARGEBACK   PaymentStatusEnum = 6
	AUTHORIZED   PaymentStatusEnum = 7
	IN_PROCESS   PaymentStatusEnum = 8
	IN_MEDIATION PaymentStatusEnum = 9
)
