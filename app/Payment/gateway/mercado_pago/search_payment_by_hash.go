package mercadopago

import (
	"code-space-backend-api/app/Payment/gateway/mercado_pago/dto"
	"encoding/json"
	"fmt"
	"net/http"
)

func (m *mercadoPagoService) SearchPaymentByHash(paymentHash string) (paymentSearch dto.PaymentSearchOutputDTO, err error) {
	var paymentSeachUrl string = fmt.Sprintf("%s/%s", m.routes.searchPaymentDetails, paymentHash)

	payload, err := m.createRequestAndExecute(http.MethodGet, paymentSeachUrl)
	if err != nil {
		return
	}

	return paymentSearch, json.Unmarshal(payload, &paymentSearch)
}
