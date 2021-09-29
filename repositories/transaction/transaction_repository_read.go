package transaction

import (
	"mileapp-technical-test-backend/models/db"

	"github.com/gocraft/dbr"
	"github.com/gofrs/uuid"
)

func GetTransactionByID(sess *dbr.Session, transactionID uuid.UUID) (*db.Transaction, error) {
	var transaction *db.Transaction

	query := sess.Select("*").
		From(db.Transaction{}.TableName()).
		Where("transaction_id = ?", transactionID)

	err := query.LoadOne(&transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func GetTransactions(sess *dbr.Session) ([]db.Transaction, error) {
	var transactions []db.Transaction

	query := sess.Select("*").
		From(db.Transaction{}.TableName())

	_, err := query.Load(&transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
