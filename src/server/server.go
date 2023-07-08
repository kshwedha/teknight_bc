package server

import (
	"log"
	"net/http"
	"teknight_bc/config"
)

func Serve() {
	config := config.GetConfig()
	router := Routes()
	log.Fatal(http.ListenAndServe(config.GetString("server.port"), router))
}
