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
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.recordWin(player)
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
	recordWin(name string)
}

type InMemoryPlayerStore struct {
	scores map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		scores: make(map[string]int),
	}
}

func (s InMemoryPlayerStore) getPlayerScore(name string) (int, bool) {
	score, found := s.scores[name]
	return score, found
}

func (s InMemoryPlayerStore) recordWin(name string) {
	s.scores[name]++
}
