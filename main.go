package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/julianinsua/codis/files"
	"github.com/julianinsua/codis/http"
	"github.com/julianinsua/codis/internal/database"
	"github.com/julianinsua/codis/parser"
	"github.com/julianinsua/codis/token"
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
		log.Fatal("sql error", err)
	}

	store := database.NewStore(db)
	parser := parser.NewMdParser()
	tokenMaker, err := token.NewPASETOMaker(config.SymetricKey)
	if err != nil {
		log.Fatal("token maker error", err)
	}
	fileManager := files.NewLocalFileManager(config.UploadFilePath)
	server := http.NewServer(store, parser, config, tokenMaker, fileManager)
	fmt.Println("Just another beautifull day in the server")
	server.Start()
}
