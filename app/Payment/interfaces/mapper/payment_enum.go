package mapper

import (
	"code-space-backend-api/app/Payment/gateway/mercado_pago/enum"
	"code-space-backend-api/infra/database/enums"
)

func PaymentStatusEnumToModel(status enum.PaymentStatus) enums.PaymentStatusEnum {
	switch status {
	case enum.APPROVED:
		return enums.APPROVED
	case enum.PENDING:
		return enums.PENDING
	case enum.REFUNDED:
		return enums.REFUNDED
	case enum.REJECTED:
		return enums.REJECTED
	case enum.CANCELLED:
		return enums.CANCELLED
	case enum.CHARGEBACK:
		return enums.CHARGEBACK
	case enum.AUTHORIZED:
		return enums.AUTHORIZED
	case enum.IN_PROCESS:
		return enums.IN_PROCESS
	case enum.IN_MEDIATION:
		return enums.IN_MEDIATION
	default:
		return enums.PENDING
	}
}

func PaymentMethodEnumToModel(method enum.PaymentMethod) enums.PaymentMethodEnum {
	switch method {
	case enum.Pix:
		return enums.Pix
	case enum.AccountMoney:
		return enums.AccountMoney
	case enum.DebinTransfer:
		return enums.DebinTransfer
	case enum.TED:
		return enums.TED
	case enum.CVU:
		return enums.CVU
	case enum.PSE:
		return enums.PSE
	default:
		return enums.AccountMoney
	}
}

func PaymentTypeEnumToModel(paymentType enum.PaymentType) enums.PaymentTypeEnum {
	switch paymentType {
	case enum.ATM:
		return enums.ATM
	case enum.TICKET:
		return enums.TICKET
	case enum.DEBIT_CARD:
		return enums.DEBIT_CARD
	case enum.CREDIT_CARD:
		return enums.CREDIT_CARD
	case enum.VOUCHER_CARD:
		return enums.VOUCHER_CARD
	case enum.PREPAID_CARD:
		return enums.PREPAID_CARD
	case enum.ACCOUNT_MONEY:
		return enums.ACCOUNT_MONEY
	case enum.BANK_TRANSFER:
		return enums.BANK_TRANSFER
	case enum.DIGITAL_WALLET:
		return enums.DIGITAL_WALLET
	case enum.CRYPTO_TRANSFER:
		return enums.CRYPTO_TRANSFER
	case enum.DIGITAL_CURRENCY:
		return enums.DIGITAL_CURRENCY
	default:
		return enums.ACCOUNT_MONEY
	}
}
