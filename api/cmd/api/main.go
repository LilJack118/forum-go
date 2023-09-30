package main

import (
	"forum/api/config"
	"forum/api/internal/server"
	"forum/api/pkg/db"
	"log"
)

func main() {
	client, err := db.InitMongoClient()
	if err != nil {
		log.Fatal(err)
	}

	app := server.NewServer(client)

	port, _ := config.Config("SERVER_PORT", "string")
	if err := app.Run(port.(string)); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
