package db

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/gofrs/uuid"
)

type Customer struct {
	ConnoteID       uuid.UUID    `db:"connote_id" json:"connote_id"`
	OriginData      CustomerData `db:"origin_data" json:"origin_data"`
	DestinationData CustomerData `db:"destination_data" json:"destination_data"`
}

func (c Customer) TableName() string {
	return "customers"
}

type CustomerData struct {
	CustomerName          string  `db:"customer_name" json:"customer_name"`
	CustomerAddress       string  `db:"customer_address" json:"customer_address"`
	CustomerEmail         *string `db:"customer_email" json:"customer_email,omitempty"`
	CustomerPhone         string  `db:"customer_phone" json:"customer_phone"`
	CustomerAddressDetail *string `db:"customer_address_detail" json:"customer_address_detail,omitempty"`
	CustomerZipCode       string  `db:"customer_zip_code" json:"customer_zip_code"`
	ZoneCode              string  `db:"zone_code" json:"zone_code"`
	OrganizationID        int     `db:"organization_id" json:"organization_id"`
	LocationID            string  `db:"location_id" json:"location_id"`
}

func (a CustomerData) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *CustomerData) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
