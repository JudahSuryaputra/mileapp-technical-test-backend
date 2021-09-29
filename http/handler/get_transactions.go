package handler

import (
	"encoding/json"
	"mileapp-technical-test-backend/repositories/transaction"
	"net/http"

	"github.com/gocraft/dbr"
)

type GetTransactions struct {
	DBConn *dbr.Connection
}

func (c GetTransactions) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sess := c.DBConn.NewSession(nil)

	transactions, err := transaction.GetTransactions(sess)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
	return
}
