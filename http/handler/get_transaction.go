package handler

import (
	"encoding/json"
	"mileapp-technical-test-backend/models/responses"
	"mileapp-technical-test-backend/repositories/connote"
	"mileapp-technical-test-backend/repositories/customer"
	"mileapp-technical-test-backend/repositories/koli"
	"mileapp-technical-test-backend/repositories/transaction"
	"net/http"

	"github.com/gocraft/dbr"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
)

type GetTransaction struct {
	DBConn *dbr.Connection
}

func (c GetTransaction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathVariable := mux.Vars(r)
	transactionID, _ := uuid.FromString(pathVariable["id"])

	sess := c.DBConn.NewSession(nil)

	currentTransaction, err := transaction.GetTransactionByID(sess, transactionID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	currentConnote, err := connote.GetConnoteByTransactionID(sess, currentTransaction.TransactionID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	currentCustomer, err := customer.GetCustomerByConnoteID(sess, currentConnote.ConnoteID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	currentKoli, err := koli.GetKolisByConnoteID(sess, currentConnote.ConnoteID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	response := responses.TransactionDetailResponse{
		Transaction: *currentTransaction,
		Connote:     *currentConnote,
		Customer:    *currentCustomer,
		Kolis:       currentKoli,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	return
}
