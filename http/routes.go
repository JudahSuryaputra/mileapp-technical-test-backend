package http

import (
	"fmt"
	"log"
	"mileapp-technical-test-backend/http/handler"
	"net/http"

	"github.com/gocraft/dbr"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func RunServer(dbConn *dbr.Connection) {
	port := viper.GetString("PORT")
	r := mux.NewRouter()

	// routes here
	createTransaction := handler.CreateTransaction{DBConn: dbConn}
	getTransactions := handler.GetTransactions{DBConn: dbConn}
	getTransaction := handler.GetTransaction{DBConn: dbConn}
	updateTransaction := handler.UpdateTransaction{DBConn: dbConn}
	deleteTransaction := handler.DeleteTransaction{DBConn: dbConn}
	r.Handle("/transaction", createTransaction).Methods(http.MethodPost)
	r.Handle("/transactions", getTransactions).Methods(http.MethodGet)
	r.Handle("/transaction/{id}", getTransaction).Methods(http.MethodGet)
	r.Handle("/transaction/{id}", updateTransaction).Methods(http.MethodPut)
	r.Handle("/transaction/{id}", deleteTransaction).Methods(http.MethodDelete)

	fmt.Printf("\n Server starting on Port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
