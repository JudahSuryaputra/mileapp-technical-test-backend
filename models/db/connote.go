package db

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/gofrs/uuid"
)

type Connote struct {
	ConnoteID              uuid.UUID       `db:"connote_id" json:"connote_id"`         // from db -- uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	ConnoteNumber          int             `db:"connote_number" json:"connote_number"` // serial
	ConnoteService         string          `db:"connote_service" json:"connote_service"`
	ConnoteServicePrice    int             `db:"connote_service_price" json:"connote_service_price"`
	ConnoteAmount          int             `db:"connote_amount" json:"connote_amount"`
	ConnoteCode            string          `db:"connote_code" json:"connote_code"`
	ConnoteBookingCode     string          `db:"connote_booking_code" json:"connote_booking_code"`
	ConnoteOrder           int             `db:"connote_order" json:"connote_order"`
	ConnoteState           string          `db:"connote_state" json:"connote_state"`
	ConnoteStateID         int             `db:"connote_state_id" json:"connote_state_id"`
	ZoneCodeFrom           string          `db:"zone_code_from" json:"zone_code_from"`
	ZoneCodeTo             string          `db:"zone_code_to" json:"zone_code_to"`
	SurchargeAmount        *int            `db:"surcharge_amount" json:"surcharge_amount,omitempty"`
	TransactionID          uuid.UUID       `db:"transaction_id" json:"transaction_id"`
	ActualWeight           int             `db:"actual_weight" json:"actual_weight"`
	VolumeWeight           *int            `db:"volume_weight" json:"volume_weight"`
	ChargeableWeight       int             `db:"chargeable_weight" json:"chargeable_weight"`
	CreatedAt              time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt              time.Time       `db:"updated_at" json:"updated_at"`
	OrganizationID         int             `db:"organization_id" json:"organization_id"`
	LocationID             string          `db:"location_id" json:"location_id"`
	ConnoteTotalPackage    int             `db:"connote_total_package" json:"connote_total_package"`
	ConnoteSurchargeAmount *string          `db:"connote_surcharge_amount" json:"connote_surcharge_amount"`
	ConnoteSLADay          string          `db:"connote_sla_day" json:"connote_sla_day"`
	LocationName           string          `db:"location_name" json:"location_name"`
	LocationType           string          `db:"location_type" json:"location_type"`
	SourceTariffDb         string          `db:"source_tariff_db" json:"source_tariff_db"`
	IDSourceTariff         string          `db:"id_source_tariff" json:"id_source_tariff"`
	Pod                    *string         `db:"pod" json:"pod,omitempty"`
	History                *[]string       `db:"history" json:"history,omitempty"`
	CustomField            CustomField     `db:"custom_field" json:"custom_field"`
	CurrentLocation        CurrentLocation `db:"current_location" json:"current_location"`
}

func (c Connote) TableName() string {
	return "connotes"
}

type CustomField struct {
	CatatanTambahan string `db:"catatan_tambahan" json:"catatan_tambahan"`
}

type CurrentLocation struct {
	Name string `db:"name" json:"name"`
	Code string `db:"code" json:"code"`
	Type string `db:"type" json:"type"`
}

func (a CustomField) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *CustomField) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

func (a CurrentLocation) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *CurrentLocation) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
