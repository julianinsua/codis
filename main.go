package main

import (
	"database/sql"
	"fmt"
	"log"

	database "github.com/julianinsua/codis/database"
	"github.com/julianinsua/codis/http"
	_ "github.com/lib/pq"
)

const (
	DB_SOURCE = "postgres"
	DB_URL    = "postgresql://postgres:password@localhost:5432/codis?sslmode=disable"
)

func main() {
	db, err := sql.Open(DB_SOURCE, DB_URL)
	if err != nil {
		log.Fatal(err)
	}

	store := database.NewStore(db)
	server := http.NewServer(store)
	fmt.Println("Just a beautifull day in the server")
	server.Start()

}
