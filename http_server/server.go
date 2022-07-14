package http_server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerServer struct {
	store PlayerStore
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	return &PlayerServer{store}
}

func (p PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, p.store.getPlayerScore(player))
}

type PlayerStore interface {
	getPlayerScore(name string) int
}

type InMemoryPlayerStore struct{}

func (s InMemoryPlayerStore) getPlayerScore(name string) int {
	return 123
}
