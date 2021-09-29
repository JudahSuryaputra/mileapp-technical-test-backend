package transaction

import (
	"context"
	"mileapp-technical-test-backend/models/db"
	"mileapp-technical-test-backend/models/requests"

	"github.com/gocraft/dbr"
	"github.com/gofrs/uuid"
)

func CreateTransaction(tx *dbr.Tx, insertTransaction db.Transaction) (uuid.UUID, *dbr.Tx, error) {
	columns := []string{
		"customer_name",
		"customer_code",
		"transaction_amount",
		"transaction_discount",
		"transaction_additional_field",
		"transaction_payment_type",
		"transaction_state",
		"transaction_code",
		"transaction_order",
		"location_id",
		"organization_id",
		"created_at",
		"updated_at",
		"transaction_payment_type_name",
		"transaction_cash_amount",
		"transaction_cash_change",
		"customer_attribute",
	}

	var transactionID uuid.UUID

	err := tx.InsertInto(db.Transaction{}.TableName()).
		Columns(columns...).
		Record(insertTransaction).
		Returning("transaction_id").
		LoadContext(context.Background(), &transactionID)
	if err != nil {
		tx.Rollback()
		return uuid.Nil, tx, err
	}

	return transactionID, tx, nil
}

func DeleteTransaction(sess *dbr.Session, transactionID uuid.UUID) error {
	_, err := sess.DeleteFrom(db.Transaction{}.TableName()).
		Where("transaction_id = ?", transactionID).Exec()

	if err != nil {
		return err
	}

	return nil
}

func UpdateTransaction(sess *dbr.Session, r requests.UpdateTransactionRequest, transactionID uuid.UUID) error {
	var data map[string]interface{}

	data["customer_name"]= r.CustomerName
	data["customer_code"]= r.CustomerCode
	data["transaction_amount"]= r.TransactionAmount
	data["transaction_discount"]= r.TransactionDiscount
	data["transaction_additional_field"]= r.TransactionAdditionalField
	data["transaction_state"]= r.TransactionState
	data["transaction_code"]= r.TransactionCode
	data["transaction_order"]= r.TransactionOrder
	data["location_id"]= r.LocationID
	data["organization_id"]= r.OrganizationID
	data["transaction_payment_type_name"]= r.TransactionPaymentTypeName
	data["transaction_cash_amount"]= r.TransactionCashAmount
	data["transaction_cash_change"]= r.TransactionCashChange
	data["customer_attribute"]= r.CustomerAttribute

	_, err := sess.Update(db.Transaction{}.TableName()).
		SetMap(data).
		Where("transaction_id = ?", transactionID).
		Exec()

	if err != nil {
		return err
	}

	return nil
}
