package db

import "github.com/gofrs/uuid"

type Transaction struct {
	TransactionID              uuid.UUID         `db:"transaction_id" json:"transaction_id"`
	CustomerName               string            `db:"customer_name" json:"customer_name"`
	CustomerCode               string            `db:"customer_code" json:"customer_code"`
	TransactionAmount          int               `db:"transaction_amount" json:"transaction_amount"`
	TransactionDiscount        int               `db:"transaction_discount" json:"transaction_discount"`
	TransactionAdditionalField string            `db:"transaction_additional_field" json:"transaction_additional_field"`
	TransactionPaymentType     string            `db:"transaction_payment_type" json:"transaction_payment_type"`
	TransactionState           string            `db:"transaction_state" json:"transaction_state"`
	TransactionCode            string            `db:"transaction_code" json:"transaction_code"`
	TransactionOrder           int               `db:"transaction_order" json:"transaction_order"`
	LocationID                 string            `db:"location_id" json:"location_id"`
	OrganizationID             int               `db:"organization_id" json:"organization_id"`
	CreatedAt                  string            `db:"created_at" json:"created_at"`
	UpdatedAt                  string            `db:"updated_at" json:"updated_at"`
	TransactionPaymentTypeName string            `db:"transaction_payment_type_name" json:"transaction_payment_type_name"`
	TransactionCashAmount      int               `db:"transaction_cash_amount" json:"transaction_cash_amount"`
	TransactionCashChange      int               `db:"transaction_cash_change" json:"transaction_cash_change"`
	CustomerAttribute          CustomerAttribute `db:"customer_attribute" json:"customer_attribute"`
}

func (c Transaction) TableName() string {
	return "transactions"
}

type CustomerAttribute struct {
	NamaSales      string `db:"nama_sales" json:"nama_sales"`
	Top            string `db:"top" json:"top"`
	JenisPelanggan string `db:"jenis_pelanggan" json:"jenis_pelanggan"`
}
