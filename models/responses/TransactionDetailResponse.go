package responses

import "mileapp-technical-test-backend/models/db"

type TransactionDetailResponse struct {
	Transaction db.Transaction `json:"transaction"`
	Connote     db.Connote     `json:"connote"`
	Customer    db.Customer    `json:"customer"`
	Kolis       []db.Koli      `json:"koli_data"`
	CustomField CustomField    `json:"custom_field"`
}

type CustomField struct {
	CatatanTambahan *string `json:"catatan_tambahan,omitempty"`
}

type CurrentLocation struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Type string `json:"type"`
}
