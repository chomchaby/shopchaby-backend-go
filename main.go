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

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// host := os.Getenv("HOST")
	// port := os.Getenv("DBPORT")
	// user := os.Getenv("USER")
	// password := os.Getenv("PASSWORD")
	// dbname := os.Getenv("DBNAME")

	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err = db.Ping(); err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	// fmt.Println("successfully connected to database")

}
