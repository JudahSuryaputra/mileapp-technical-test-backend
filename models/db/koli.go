package db

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/gofrs/uuid"
)

type Koli struct {
	KoliLength           *int            `db:"koli_length" json:"koli_length"`
	AwbURL               string          `db:"awb_url" json:"awb_url"`
	CreatedAt            time.Time       `db:"created_at" json:"created_at"`
	KoliChargeableWeight int             `db:"koli_chargeable_weight" json:"koli_chargeable_weight"`
	KoliWidth            *int            `db:"koli_width" json:"koli_width"`
	KoliSurcharge        *[]string       `db:"koli_surcharge" json:"koli_surcharge"`
	KoliHeight           *int            `db:"koli_height" json:"koli_height"`
	UpdatedAt            time.Time       `db:"updated_at" json:"updated_at"`
	KoliDescription      string          `db:"koli_description" json:"koli_description"`
	KoliFormulaID        *string         `db:"koli_formula_id" json:"koli_formula_id"`
	ConnoteID            uuid.UUID       `db:"connote_id" json:"connote_id"`
	KoliVolume           *int            `db:"koli_volume" json:"koli_volume"`
	KoliWeight           int             `db:"koli_weight" json:"koli_weight"`
	KoliID               uuid.UUID       `db:"koli_id" json:"koli_id"`
	KoliCustomField      KoliCustomField `db:"koli_custom_field" json:"koli_custom_field"`
	KoliCode             string          `db:"koli_code" json:"koli_code"`
}

func (c Koli) TableName() string {
	return "kolis"
}

type KoliCustomField struct {
	AwbSicepat  string `json:"awb_sicepat"`
	HargaBarang string `json:"harga_barang"`
}

func (a KoliCustomField) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *KoliCustomField) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
