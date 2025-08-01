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
	gotenv.Load()
}

func main() {
	database.Connect_To_PostgreSQL()
	database.Connect_To_Redis()
	router.Router()

	fmt.Println("Server running...")
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
