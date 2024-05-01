package mercadopago

import (
	"bytes"
	"code-space-backend-api/common/errors"
	"fmt"
	"io"
	"net/http"
)

const mercadoPagoContext string = "adapter/mercado-pago"

var (
	couldNotCreateRequest = errors.InvalidField.
				WithContext(mercadoPagoContext).
				WithName("error_to_create_http_request").
				WithTemplate("failed to create http request: %v")

	requestWithErrorStatus = errors.BusinessRule.
				WithContext(mercadoPagoContext).
				WithName("error_response_status").
				WithTemplate("response status is an error: %v")
)

func (m *mercadoPagoService) createRequestAndExecute(method, url string, payload ...byte) (responseBody []byte, err error) {
	request, err := m.createRequest(method, url, payload...)
	if err != nil {
		return
	}

	response, err := m.client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= http.StatusBadRequest {
		err = requestWithErrorStatus.WithArgs(response.Status)
		return
	}

	responseBody, err = io.ReadAll(response.Body)

	return
}

func (m *mercadoPagoService) createRequest(method, url string, payload ...byte) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, couldNotCreateRequest.WithArgs(err.Error())
	}

	m.setupRequestHeaders(req)

	return req, nil
}

func (m *mercadoPagoService) setupRequestHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", m.authorization))
}
