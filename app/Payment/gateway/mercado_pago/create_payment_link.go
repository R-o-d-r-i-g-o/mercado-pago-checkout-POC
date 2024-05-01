package mercadopago

import (
	"code-space-backend-api/app/Payment/gateway/mercado_pago/dto"
	"code-space-backend-api/app/Payment/interfaces/dto/input"
	"encoding/json"
	"net/http"
)

func (m *mercadoPagoService) CreatePaymentLink(paymentLinkInputDTO input.PaymentLinkInputDTO) (paymentLink dto.PaymentLinkOutputDTO, err error) {
	if err = paymentLinkInputDTO.ValidateDTO(); err != nil {
		return
	}

	payload, err := json.Marshal(&paymentLinkInputDTO)
	if err != nil {
		return
	}

	responseBody, err := m.createRequestAndExecute(http.MethodPost, m.routes.checkoutPreferences, payload...)
	if err != nil {
		return
	}

	return paymentLink, json.Unmarshal(responseBody, &paymentLink)
}
