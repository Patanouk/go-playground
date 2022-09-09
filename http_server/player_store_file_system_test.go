package http_server

import (
	"io"
	"os"
	"testing"
)

// file_system_store_test.go
func TestFileSystemStore(t *testing.T) {

	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

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
		database, cleanDatabase := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		got, found := store.getPlayerScore("Chris")
		want := 33

		assertPlayerFound(t, found)
		assertScoreEquals(t, got, want)
	})

	t.Run("store win for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database}
		store.recordWin("Chris")

		got, found := store.getPlayerScore("Chris")
		want := 34

		assertPlayerFound(t, found)
		assertScoreEquals(t, got, want)
	})
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tempFile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tempFile.Write([]byte(initialData))

	removeFile := func() {
		tempFile.Close()
		os.RemoveAll(tempFile.Name())
	}

	return tempFile, removeFile
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
