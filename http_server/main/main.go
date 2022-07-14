package main

import (
	"go-playground/http_server"
	"log"
	"net/http"
)

func main() {
	server := http_server.NewPlayerServer(http_server.NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}
