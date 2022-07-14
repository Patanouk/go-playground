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
}
