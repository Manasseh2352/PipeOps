package main

import (
	"database/sql"
	"log"

	db "github.com/Manasseh2352/PipeOps/backend/internal/db/sqlc"
	api "github.com/Manasseh2352/PipeOps/backend/internal/server"
	_ "github.com/lib/pq"
)

func main() {
	// config, err :=  utils.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("cannot load config: ", err)
	// }
	// gin.SetMode("deb")
	config := struct {
		DBDriver      string
		DBSource      string
		ServerAddress string
	}{
		DBDriver:      "postgres",
		DBSource:      "postgres://root:secret@localhost:5432/skudoosh?sslmode=disable",
		ServerAddress: ":8080",
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

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot Start Server: ", err)
	}

}
