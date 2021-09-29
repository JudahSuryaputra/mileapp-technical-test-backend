package customer

import (
	"mileapp-technical-test-backend/models/db"

	"github.com/gocraft/dbr"
	"github.com/gofrs/uuid"
)

func GetCustomerByConnoteID(sess *dbr.Session, connoteID uuid.UUID) (*db.Customer, error) {
	var customer *db.Customer

	query := sess.Select("*").
		From(db.Customer{}.TableName()).
		Where("connote_id = ?", connoteID)

	err := query.LoadOne(&customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
