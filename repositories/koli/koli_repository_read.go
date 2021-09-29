package koli

import (
	"mileapp-technical-test-backend/models/db"

	"github.com/gocraft/dbr"
	"github.com/gofrs/uuid"
)

func GetKolisByConnoteID(sess *dbr.Session, connoteID uuid.UUID) ([]db.Koli, error) {
	var kolis []db.Koli

	query := sess.Select("*").
		From(db.Koli{}.TableName()).
		Where("connote_id = ?", connoteID)

	_, err := query.Load(&kolis)
	if err != nil {
		return nil, err
	}

	return kolis, err
}
