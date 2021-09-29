package customer

import (
	"mileapp-technical-test-backend/models/db"

	"github.com/gocraft/dbr"
)

func InsertCustomer(tx *dbr.Tx, customer db.Customer) (*dbr.Tx, error) {
	columns := []string{
		"connote_id",
		"origin_data",
		"destination_data",
	}

	_, err := tx.InsertInto(db.Customer{}.TableName()).
		Columns(columns...).
		Record(customer).
		Exec()
		
	if err != nil {
		tx.Rollback()
		return tx, err
	}

	return tx, nil
}
