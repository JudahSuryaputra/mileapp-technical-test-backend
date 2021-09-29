package handler

import (
	"encoding/json"
	"mileapp-technical-test-backend/repositories/transaction"
	"net/http"

	"github.com/gocraft/dbr"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
)

type DeleteTransaction struct {
	DBConn *dbr.Connection
}

func (c DeleteTransaction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathVariable := mux.Vars(r)
	transactionID, _ := uuid.FromString(pathVariable["id"])

	sess := c.DBConn.NewSession(nil)

	err := transaction.DeleteTransaction(sess, transactionID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	return
}
