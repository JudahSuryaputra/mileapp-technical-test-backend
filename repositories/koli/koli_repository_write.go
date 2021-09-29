package koli

import (
	"mileapp-technical-test-backend/models/db"

	"github.com/gocraft/dbr"
)

func InsertKolis(tx *dbr.Tx, kolis []db.Koli) (*dbr.Tx, error) {
	columns := []string{
		"koli_length",
		"awb_url",
		"koli_chargeable_weight",
		"koli_width",
		"koli_height",
		"koli_description",
		"koli_volume",
		"koli_weight",
		"connote_id",
		"koli_id",
		"koli_code",
	}

	query := tx.InsertInto(db.Koli{}.TableName()).Columns(columns...)

	for _, row := range kolis {
		query.Record(row)
	}

	_, err := query.Exec()
	if err != nil {
		tx.Rollback()
		return tx, err
	}

	return tx, nil
}
