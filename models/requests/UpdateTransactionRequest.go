package requests

import (
	"mileapp-technical-test-backend/models/db"
)

type UpdateTransactionRequest struct {
	CustomerName               string               `json:"customer_name"`
	CustomerCode               string               `json:"customer_code"`
	TransactionAmount          int                  `json:"transaction_amount"`
	TransactionDiscount        int                  `json:"transaction_discount"`
	TransactionAdditionalField string               `json:"transaction_additional_field"`
	TransactionPaymentType     string               `json:"transaction_payment_type"`
	TransactionState           string               `json:"transaction_state"`
	TransactionCode            string               `json:"transaction_code"`
	TransactionOrder           int                  `json:"transaction_order"`
	LocationID                 string               `json:"location_id"`
	OrganizationID             int                  `json:"organization_id"`
	TransactionPaymentTypeName string               `json:"transaction_payment_type_name"`
	TransactionCashAmount      int                  `json:"transaction_cash_amount"`
	TransactionCashChange      int                  `json:"transaction_cash_change"`
	CustomerAttribute          db.CustomerAttribute `json:"customer_attribute"`
}
