package enum

type NotificationType string

const (
	PAYMENT_CREATED NotificationType = "payment.created"
	PAYMENT_UPDATED NotificationType = "payment.updated"
)
