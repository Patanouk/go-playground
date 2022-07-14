package http_server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	store := stubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := getScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")
		assertStatusCode(t, response.Code, http.StatusOK)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := getScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "10")
		assertStatusCode(t, response.Code, http.StatusOK)
	})

	t.Run("returns 404 on missing player", func(t *testing.T) {
		request := getScoreRequest("Missing")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatusCode(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := stubPlayerStore{
		scores: make(map[string]int),
	}
	server := NewPlayerServer(&store)

	t.Run("server returns increment score when posting win", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatusCode(t, response.Code, http.StatusAccepted)

		request, _ = http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response = httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "1")
	})
}

//server_test.go
func TestLeague(t *testing.T) {
	wantedLeague := []Player{
		{"Cleo", 32},
		{"Chris", 20},
		{"Tiest", 14},
	}
	store := stubPlayerStore{league: wantedLeague}
	server := NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got []Player

		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response '%q' from /league endpoint into slice of players, '%v'", response.Body, err)
		}

		assertStatus(t, response.Code, http.StatusOK)

		if !reflect.DeepEqual(got, wantedLeague) {
			t.Errorf("got %v, want %v", got, wantedLeague)
		}
	})
}

type stubPlayerStore struct {
	scores map[string]int
	league []Player
}

func (p *stubPlayerStore) recordWin(name string) {
	p.scores[name] = p.scores[name] + 1
}

func (p *stubPlayerStore) getPlayerScore(name string) (int, bool) {
	score, ok := p.scores[name]
	return score, ok
}

func (p *stubPlayerStore) getLeague() []Player {
	return p.league
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func assertStatusCode(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d, want %d", got, want)
	}
}

func getScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}
