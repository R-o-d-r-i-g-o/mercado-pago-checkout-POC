package mercadopago

import (
	"code-space-backend-api/app/Payment/gateway/mercado_pago/dto"
	"code-space-backend-api/app/Payment/interfaces/dto/input"
	"net/http"
)

// MercadoPagoService provides an interface for interacting with the Mercado Pago API.
type MercadoPagoService interface {
	SearchPaymentByHash(paymentHash string) (paymentSearch dto.PaymentSearchOutputDTO, err error)
	CreatePaymentLink(paymentLinkInputDTO input.PaymentLinkInputDTO) (paymentLink dto.PaymentLinkOutputDTO, err error)
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type mercadoPagoService struct {
	baseUrl       string
	authorization string
	client        httpClient
	routes        routes
}

type routes struct {
	checkoutPreferences  string
	searchPaymentDetails string
}

func NewMercadoPagoService(baseUrl, authorization string, client httpClient) MercadoPagoService {
	return &mercadoPagoService{
		baseUrl:       baseUrl,
		authorization: authorization,
		client:        client,
		routes: routes{
			checkoutPreferences:  baseUrl + "/checkout/preferences",
			searchPaymentDetails: baseUrl + "/v1/payments",
		},
	}
}
