package usecases

import (
	"code-space-backend-api/app/Payment/entities"
	mercadopago "code-space-backend-api/app/Payment/gateway/mercado_pago"
	"code-space-backend-api/app/Payment/interfaces/dto/input"
	"code-space-backend-api/app/Payment/interfaces/repository"
	"code-space-backend-api/common/constants/str"
	"code-space-backend-api/common/errors"
	"context"
)

const createPaymentLinkContext = "usecase/create-payment-link"

var (
	invalidPaymentLinkInput = errors.InvalidField.
				WithContext(createPaymentLinkContext).
				WithName("invalid_field_send").
				WithTemplate("field send is either required or invalid: %v")

	failedToRetrievePaymentLinkFromService = errors.Unknown.
						WithContext(createPaymentLinkContext).
						WithName("failed_to_retrieve_payment_link_from_service").
						WithTemplate("something went wrong when retrieving the payment link: %v")

	failedToSavePaymentInfo = errors.Unknown.
				WithContext(createPaymentLinkContext).
				WithName("error_when_storing_payment_in_database").
				WithTemplate("something went wrong when saving payment in database: %v")
)

type CreatePaymentLink interface {
	Execute(ctx context.Context, paymentLinkDTO input.PaymentLinkInputDTO) (string, error)
}

type createPaymentLink struct {
	repository     repository.PurchaseRepository
	paymentService mercadopago.MercadoPagoService
}

func NewCreatePaymentLink(repository repository.PurchaseRepository, paymentService mercadopago.MercadoPagoService) CreatePaymentLink {
	return &createPaymentLink{
		repository:     repository,
		paymentService: paymentService,
	}
}

func (usecase *createPaymentLink) Execute(ctx context.Context, paymentLinkDTO input.PaymentLinkInputDTO) (string, error) {
	if err := paymentLinkDTO.ValidateDTO(); err != nil {
		return str.EMPTY_STRING, invalidPaymentLinkInput.WithArgs(err.Error())
	}

	paymentLink, err := usecase.paymentService.CreatePaymentLink(paymentLinkDTO)
	if err != nil {
		return str.EMPTY_STRING, failedToRetrievePaymentLinkFromService.WithArgs(err.Error())
	}

	var payer entities.Payer = paymentLinkDTO.ToDomain()

	payer.Purchase = payer.Purchase.
		WithStatus("pending").
		WithClientID(paymentLink.ClientID).
		WithCollectorID(paymentLink.CollectorID).
		WithPreferenceID(paymentLink.PreferenceID)

	if err = usecase.repository.CreatePurchaseWithProducts(payer, payer.Purchase); err != nil {
		return str.EMPTY_STRING, failedToSavePaymentInfo.WithArgs(err.Error())
	}

	return paymentLink.InitPoint, nil
}
