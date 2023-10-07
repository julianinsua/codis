package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/julianinsua/codis/api"
	database "github.com/julianinsua/codis/database"
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
	server := api.NewServer(store)
	fmt.Println("Just a beautifull day in the server")
	server.Start()

}
