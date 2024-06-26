package main

import (
	"database/sql"
	"fmt"
	"log"

	db "github.com/Manasseh2352/PipeOps/backend/internal/db/sqlc"
	api "github.com/Manasseh2352/PipeOps/backend/internal/server"
	"github.com/Manasseh2352/PipeOps/backend/internal/utils"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("start")
	
	config, err :=  utils.LoadConfig(".")
	fmt.Printf("printing it %s", config.ServerAddress)
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	dbConn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(dbConn)
	server, err := api.NewServer(store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}
	fmt.Println(config.ServerAddress)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot Start Server: ", err)
	}

}
