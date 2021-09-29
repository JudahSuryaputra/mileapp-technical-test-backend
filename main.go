package main

import (
	"mileapp-technical-test-backend/http"
	"mileapp-technical-test-backend/repositories"
)

func main() {
	conn := repositories.Conn()

	http.RunServer(conn)
}
