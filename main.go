package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rifqiahmad/crispy-engine/api"
	db "github.com/rifqiahmad/crispy-engine/db/sqlc"
	"github.com/rifqiahmad/crispy-engine/util"
	"log"
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

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
