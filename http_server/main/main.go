package main

import (
	"go-playground/http_server"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(http_server.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
