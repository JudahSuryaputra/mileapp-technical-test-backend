package handler

import (
	"encoding/json"
	"mileapp-technical-test-backend/models/requests"
	"mileapp-technical-test-backend/repositories/transaction"
	"net/http"

	"github.com/gocraft/dbr"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
)

type UpdateTransaction struct {
	DBConn *dbr.Connection
}

func (c UpdateTransaction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathVariable := mux.Vars(r)
	transactionID, _ := uuid.FromString(pathVariable["id"])

	var request requests.UpdateTransactionRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	sess := c.DBConn.NewSession(nil)

	err = transaction.UpdateTransaction(sess, request, transactionID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	return
}
