package usecases

import (
	"code-space-backend-api/app/Payment/interfaces/dto/input"
	"code-space-backend-api/app/Payment/interfaces/repository"
	"code-space-backend-api/common/errors"
	"encoding/json"
	"time"
)

const waitForPaymentNotificationContext string = "usecase/wait-for-payment-notification"

var (
	invalidPaymentNotificationField = errors.InvalidField.
		WithContext(waitForPaymentNotificationContext).
		WithName("invalid_field_send").
		WithTemplate("field send is either required or invalid: %v")
)

type WaitForPaymentStatus interface {
	Execute(paymentNotification input.PaymentNotificationInputDTO) error
}

type waitForPaymentStatus struct {
	notificationRepository repository.NotificationRepository
	purchaseRepository     repository.PurchaseRepository
	paymentRepository      repository.PaymentRepository
	searchPaymentInfo      SearchPaymentInfo
	retryer
}

type retryer struct {
	maxAttempts   int
	retryInterval time.Duration
}

func NewWaitForPaymentStatus(notificationRepository repository.NotificationRepository, purchaseRepository repository.PurchaseRepository, paymentRepository repository.PaymentRepository, seachPaymentInfo SearchPaymentInfo, maxAttempts int, retryInterval time.Duration) WaitForPaymentStatus {
	return &waitForPaymentStatus{
		notificationRepository: notificationRepository,
		purchaseRepository:     purchaseRepository,
		paymentRepository:      paymentRepository,
		searchPaymentInfo:      seachPaymentInfo,
		retryer: retryer{
			maxAttempts:   maxAttempts,
			retryInterval: retryInterval,
		},
	}
}

func (usecase *waitForPaymentStatus) Execute(paymentNotification input.PaymentNotificationInputDTO) error {
	if err := paymentNotification.ValidateDTO(); err != nil {
		if err = usecase.savePaymentNotification(paymentNotification); err != nil {
			return invalidPaymentNotificationField.WithArgs(err)
		}
	}

	go func(notification input.PaymentNotificationInputDTO) {
		if err := usecase.savePaymentNotification(notification); err != nil {
			return
		}

		if err := usecase.processPaymentNotificationAsyncWithRetry(notification); err != nil {
			return
		}
	}(paymentNotification)

	return nil
}

func (usecase *waitForPaymentStatus) savePaymentNotification(paymentNotification input.PaymentNotificationInputDTO) error {
	payload, err := json.Marshal(&paymentNotification)
	if err != nil {
		return err
	}

	return usecase.notificationRepository.CreateNotification(string(payload))
}

func (usecase *waitForPaymentStatus) processPaymentNotificationAsyncWithRetry(notificationData input.PaymentNotificationInputDTO) error {
	var err error
	for attempt := 1; attempt <= usecase.maxAttempts; attempt++ {
		paymentReport, err := usecase.searchPaymentInfo.Execute(notificationData.Data.ID)
		if err != nil {
			continue
		}

		if err = usecase.paymentRepository.CreatePayment(paymentReport); err == nil {
			return nil
		}
		time.Sleep(usecase.retryInterval)
	}
	return err
}
