package main

import (
	"log"
	"net/http"

	"jwt_server/config"
	"jwt_server/server"
)

func main() {
	cfg := config.GetConfig()

	log.Printf("Starting server on port %s\n", cfg.HTTPPort)

	log.Fatal(http.ListenAndServe(":"+cfg.HTTPPort, server.GetRouter()))
}
