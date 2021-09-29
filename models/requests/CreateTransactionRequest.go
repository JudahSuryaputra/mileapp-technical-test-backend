package requests

import "mileapp-technical-test-backend/models/db"

type CreateTransactionRequest struct {
	OriginData      db.CustomerData    `json:"origin_data"`
	DestinationData db.CustomerData    `json:"destination_data"`
	KoliDatas       []KoliData         `json:"koli_datas"`
	ConnoteData     ConnoteData        `json:"connote_data"`
	TransactionData TransactionData    `json:"transaction_data"`
	CustomField     db.CustomField     `json:"custom_field"`
	CurrentLocation db.CurrentLocation `json:"current_location"`
}

type KoliData struct {
	KoliLength           *int
	KoliChargeableWeight int
	KoliWidth            *int
	KoliHeight           *int
	KoliDescription      string
	KoliVolume           *int
	KoliWeight           int
}

type ConnoteData struct {
	ConnoteService      string
	ConnoteServicePrice int
	ConnoteAmount       int
	ActualWeight        int
	ChargeableWeight    int
	LocationName        string
	LocationType        string
	SourceTariffDb      string
	IDSourceTariff      string
}

type TransactionData struct {
	TransactionPaymentTypeName string
	LocationID                 string
}
