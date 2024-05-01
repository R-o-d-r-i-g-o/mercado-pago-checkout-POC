package dto

import (
	"code-space-backend-api/app/Payment/gateway/mercado_pago/enum"
	"time"
)

type (
	PaymentSearchOutputDTO struct {
		ID                        int                `json:"id"`
		DateCreated               time.Time          `json:"date_created"`
		DateApproved              time.Time          `json:"date_approved"`
		DateLastUpdated           time.Time          `json:"date_last_updated"`
		MoneyReleaseDate          time.Time          `json:"money_release_date"`
		PaymentMethodID           enum.PaymentMethod `json:"payment_method_id"`
		PaymentTypeID             enum.PaymentType   `json:"payment_type_id"`
		Status                    enum.PaymentStatus `json:"status"`
		StatusDetail              string             `json:"status_detail"`
		CurrencyID                enum.Currency      `json:"currency_id"`
		Description               string             `json:"description"`
		CollectorID               int                `json:"collector_id"`
		Payer                     Payer              `json:"payer"`
		TransactionAmount         float64            `json:"transaction_amount"`
		TransactionAmountRefunded float64            `json:"transaction_amount_refunded"`
		CouponAmount              int                `json:"coupon_amount"`
		TransactionDetails        TransactionDetails `json:"transaction_details"`
		Installments              int                `json:"installments"`
	}

	Identification struct {
		Type   string `json:"type"`
		Number int64  `json:"number"`
	}

	Payer struct {
		ID             int            `json:"id"`
		Email          string         `json:"email"`
		Identification Identification `json:"identification"`
		Type           string         `json:"type"`
	}

	TransactionDetails struct {
		NetReceivedAmount int `json:"net_received_amount"`
		TotalPaidAmount   int `json:"total_paid_amount"`
		OverpaidAmount    int `json:"overpaid_amount"`
		InstallmentAmount int `json:"installment_amount"`
	}
)
