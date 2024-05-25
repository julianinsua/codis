package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/julianinsua/codis/http"
	"github.com/julianinsua/codis/internal/database"
	"github.com/julianinsua/codis/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("failed to log config file: ", err)
	}

	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}

	store := database.NewStore(db)
	parser := http.NewMdParser()
	server := http.NewServer(store, parser, config)
	fmt.Println("Just another beautifull day in the server")
	server.Start()
}
