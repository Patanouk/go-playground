package http_server

import (
	"strings"
	"testing"
)

//file_system_store_test.go
func TestFileSystemStore(t *testing.T) {

	t.Run("league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)

		store := FileSystemPlayerStore{database}

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		got := store.getLeague()
		assertLeague(t, got, want)

		got = store.getLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database := strings.NewReader(`[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)

		store := FileSystemPlayerStore{database}

		got, found := store.getPlayerScore("Chris")
		want := 33

		assertPlayerFound(t, found)
		assertScoreEquals(t, got, want)
	})
}

func assertPlayerFound(t *testing.T, found bool) {
	t.Helper()
	if !found {
		t.Errorf("Player not found")
	}
}

func assertScoreEquals(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}
