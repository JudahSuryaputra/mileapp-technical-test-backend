package repositories

import (
	"fmt"
	"log"

	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
)

const (
	port     = 5432
	host     = "localhost"
	user     = "postgres"
	password = "123456"
	dbname   = "mileapp_dev"
)

func Conn() *dbr.Connection {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	conn, err := dbr.Open("postgres", dsn, nil)
	if err != nil {
		log.Printf("Cannot initialize connection to database: %v", err)
	}

	log.Printf(dsn)

	return conn
}
