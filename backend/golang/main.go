package main

import (
	"test-api/config"
	"test-api/server"
)

func main() {
	db := config.ConnectDatabase()
	defer db.Close()
	server.StartServer(db)
}
