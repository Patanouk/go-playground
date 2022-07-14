package main

import (
	"go-playground/http_server"
	"log"
	"net/http"
)

func main() {
	store := &http_server.InMemoryPlayerStore{}
	server := http_server.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
