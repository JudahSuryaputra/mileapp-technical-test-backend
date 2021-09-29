package connote

import (
	"mileapp-technical-test-backend/models/db"

	"github.com/gocraft/dbr"
	"github.com/gofrs/uuid"
)

func GetConnoteByTransactionID(sess *dbr.Session, transactionID uuid.UUID) (*db.Connote, error) {
	var connote *db.Connote

	query := sess.Select("*").
		From(db.Connote{}.TableName()).
		Where("transaction_id = ?", transactionID)

	err := query.LoadOne(&connote)
	if err != nil {
		return nil, err
	}

	return nil, err
}
