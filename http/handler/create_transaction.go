package handler

import (
	"encoding/json"
	"mileapp-technical-test-backend/models/db"
	"mileapp-technical-test-backend/models/requests"
	"mileapp-technical-test-backend/repositories/connote"
	"mileapp-technical-test-backend/repositories/customer"
	"mileapp-technical-test-backend/repositories/koli"
	"mileapp-technical-test-backend/repositories/transaction"
	"net/http"
	"strconv"

	"github.com/gocraft/dbr"
	"github.com/gofrs/uuid"
)

type CreateTransaction struct {
	DBConn *dbr.Connection
}

func (c CreateTransaction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request requests.CreateTransactionRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	// check admin auth

	sess := c.DBConn.NewSession(nil)

	tx, err := sess.Begin()
	tx.RollbackUnlessCommitted()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	transactionID, tx, err := createTransaction(tx, request)
	if err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	connoteID, tx, err := createConnote(tx, request, transactionID)
	if err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	tx, err = insertCustomer(tx, request, connoteID)
	if err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	tx, err = insertKolis(tx, request, connoteID)
	if err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	tx.Commit()

	return
}

func createTransaction(tx *dbr.Tx, r requests.CreateTransactionRequest) (uuid.UUID, *dbr.Tx, error) {
	// transactionID, _ := uuid.NewV4()
	createTransactionRequest := db.Transaction{
		CustomerName:               r.OriginData.CustomerName,
		CustomerCode:               "this is customer code",
		TransactionAmount:          r.ConnoteData.ConnoteAmount,
		TransactionDiscount:        0,
		TransactionAdditionalField: "",
		TransactionPaymentType:     "29",
		TransactionState:           "PAID",
		TransactionCode:            "CGKFT20200715121",
		TransactionOrder:           121,
		LocationID:                 "5cecb20b6c49615b174c3e74", // Location ID should be from the system
		OrganizationID:             6,
		TransactionPaymentTypeName: r.TransactionData.TransactionPaymentTypeName,
		TransactionCashAmount:      0,
		TransactionCashChange:      0,
		CustomerAttribute: db.CustomerAttribute{
			NamaSales:      "Radit Fitrawikarsa",
			Top:            "14 Hari",
			JenisPelanggan: "B2B",
		},
	}

	transactionID, tx, err := transaction.CreateTransaction(tx, createTransactionRequest)
	if err != nil {
		return transactionID, tx, err
	}

	return transactionID, tx, nil
}

func createConnote(tx *dbr.Tx, r requests.CreateTransactionRequest, transactionID uuid.UUID) (uuid.UUID, *dbr.Tx, error) {
	createConnoteRequest := db.Connote{
		ConnoteService:      r.ConnoteData.ConnoteService,
		ConnoteServicePrice: r.ConnoteData.ConnoteServicePrice,
		ConnoteAmount:       r.ConnoteData.ConnoteAmount,
		ConnoteCode:         "should be a code for connote",
		ConnoteBookingCode:  "connote booking code",
		ConnoteOrder:        326931,
		ConnoteState:        "PAID",
		ConnoteStateID:      2,
		ZoneCodeFrom:        r.OriginData.ZoneCode,
		ZoneCodeTo:          r.DestinationData.ZoneCode,
		TransactionID:       transactionID,
		ActualWeight:        r.ConnoteData.ActualWeight,
		ChargeableWeight:    r.ConnoteData.ChargeableWeight,
		OrganizationID:      6,
		LocationID:          "location ID from system",
		ConnoteTotalPackage: len(r.KoliDatas),
		ConnoteSLADay:       "4",
		LocationName:        r.ConnoteData.LocationName,
		LocationType:        r.ConnoteData.LocationType,
		SourceTariffDb:      r.ConnoteData.SourceTariffDb,
		IDSourceTariff:      r.ConnoteData.IDSourceTariff,
		CustomField:         r.CustomField,
		CurrentLocation:     r.CurrentLocation,
	}

	connoteID, tx, err := connote.CreateConnote(tx, createConnoteRequest)
	if err != nil {
		return connoteID, tx, err
	}

	return connoteID, tx, nil
}

func insertCustomer(tx *dbr.Tx, r requests.CreateTransactionRequest, connoteID uuid.UUID) (*dbr.Tx, error) {
	insertCustomer := db.Customer{
		ConnoteID:       connoteID,
		OriginData:      r.OriginData,
		DestinationData: r.DestinationData,
	}

	tx, err := customer.InsertCustomer(tx, insertCustomer)
	if err != nil {
		return tx, err
	}

	return tx, nil
}

func insertKolis(tx *dbr.Tx, r requests.CreateTransactionRequest, connoteID uuid.UUID) (*dbr.Tx, error) {
	var kolis []db.Koli
	koliID, _ := uuid.NewV4()
	for _, data := range r.KoliDatas {
		var count int = 1
		itemNum := strconv.Itoa(count)
		koli := db.Koli{
			KoliLength:           data.KoliLength,
			AwbURL:               "http.tracking.mileapp/awb.url." + itemNum,
			KoliChargeableWeight: data.KoliChargeableWeight,
			KoliWidth:            data.KoliWidth,
			KoliHeight:           data.KoliHeight,
			KoliDescription:      data.KoliDescription,
			KoliVolume:           data.KoliVolume,
			KoliWeight:           data.KoliWeight,
			ConnoteID:            connoteID,
			KoliID:               koliID,
			KoliCode:             "KOLICODE." + itemNum,
		}
		kolis = append(kolis, koli)
	}

	tx, err := koli.InsertKolis(tx, kolis)
	if err != nil {
		return tx, err
	}

	return tx, nil
}
