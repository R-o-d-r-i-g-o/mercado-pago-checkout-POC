package mapper

import (
	"code-space-backend-api/app/Payment/gateway/mercado_pago/dto"
	"code-space-backend-api/infra/database/models"
)

func GatewayPaymentToModel(paymentDTO dto.PaymentSearchOutputDTO) models.Payment {
	return models.Payment{
		TypeID:               PaymentTypeEnumToModel(paymentDTO.PaymentTypeID),
		StatusID:             PaymentStatusEnumToModel(paymentDTO.Status),
		MethodID:             PaymentMethodEnumToModel(paymentDTO.PaymentMethodID),
		TransactionAmount:    paymentDTO.TransactionAmount,
		MercadoPagoPaymentID: uint(paymentDTO.ID),
	}
}
