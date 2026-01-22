package main

import (
	"database/sql"
	"log"

	"github.com/ethantiger/simplebank/util"

	"github.com/ethantiger/simplebank/api"
	db "github.com/ethantiger/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Default().Printf("[ERR] %v", err)
		log.Fatal("cannot connect to database")
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalf("Cannot create server: %v", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalf("Cannot start server %v", err)
	}
}
