package usecases

import (
	mercadopago "code-space-backend-api/app/Payment/gateway/mercado_pago"
	"code-space-backend-api/app/Payment/gateway/mercado_pago/dto"
	"code-space-backend-api/common/errors"
	"encoding/json"
)

const searchPaymentInfoContext = "usecase/create-payment-link"

var (
	failedToRetrievePaymentInfoFromService = errors.Unknown.
						WithContext(searchPaymentInfoContext).
						WithName("failed_to_retrieve_payment_link_from_service").
						WithTemplate("something went wrong when retrieving the payment link: %v")

	failedToMarshalResponse = errors.Unknown.
				WithContext(searchPaymentInfoContext).
				WithName("failed_to_marshal_response").
				WithTemplate("something went wrong when marshalling the response: %v")

	failedToUnmarshalResponse = errors.Unknown.
					WithContext(searchPaymentInfoContext).
					WithName("failed_to_unmarshal_response").
					WithTemplate("something went wrong when unmarshalling the response: %v")
)

type SearchPaymentInfo interface {
	Execute(paymentHash string) (paymentSeach dto.PaymentSearchOutputDTO, err error)
}

type searchPaymentInfo struct {
	paymentService mercadopago.MercadoPagoService
}

func NewSearchPaymentInfo(paymentService mercadopago.MercadoPagoService) SearchPaymentInfo {
	return &searchPaymentInfo{
		paymentService: paymentService,
	}
}

func (usecase *searchPaymentInfo) Execute(paymentHash string) (paymentSearch dto.PaymentSearchOutputDTO, err error) {
	paymentInfo, err := usecase.paymentService.SearchPaymentByHash(paymentHash)
	if err != nil {
		return paymentSearch, failedToRetrievePaymentInfoFromService.WithArgs(err.Error())
	}

	body, err := json.Marshal(&paymentInfo)
	if err != nil {
		return paymentSearch, failedToMarshalResponse.WithArgs(err.Error())
	}

	if err = json.Unmarshal(body, &paymentSearch); err != nil {
		return paymentSearch, failedToUnmarshalResponse.WithArgs(err.Error())
	}

	return
}
