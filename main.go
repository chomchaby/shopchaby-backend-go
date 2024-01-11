package main

import (
	"database/sql"
	"log"

	"github.com/chomchaby/shopchaby-backend-go/api"
	db "github.com/chomchaby/shopchaby-backend-go/db/sqlc"
	"github.com/chomchaby/shopchaby-backend-go/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	storeTx := db.NewStoreTx(conn)
	server, err := api.NewServer(config, storeTx)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
