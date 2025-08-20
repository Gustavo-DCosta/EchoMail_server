package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Gustavo-DCosta/server/database"
	"github.com/Gustavo-DCosta/server/router"
	"github.com/subosito/gotenv"
)

func init() {
	// load env vars from the embedded .env file
	gotenv.Load()
}

func main() {
	// init connections
	database.Connect_To_PostgreSQL()
	database.Connect_To_Redis()

	// setup router
	router.Router()

	fmt.Println("Server running on 127.0.0.1:8080")
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
