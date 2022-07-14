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

	switch r.Method {
	case http.MethodPost:
		p.processWin(w)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p PlayerServer) processWin(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
}

func (p PlayerServer) showScore(w http.ResponseWriter, player string) {
	score, found := p.store.getPlayerScore(player)
	if found {
		fmt.Fprint(w, score)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

type PlayerStore interface {
	getPlayerScore(name string) (int, bool)
}

type InMemoryPlayerStore struct{}

func (s InMemoryPlayerStore) getPlayerScore(name string) (int, bool) {
	return 123, true
}
