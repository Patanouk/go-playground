package http_server

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) getPlayerScore(name string) (int, bool) {
	for _, player := range f.getLeague() {
		if player.Name == name {
			return player.Wins, true
		}
	}

	return 0, false
}

func (f *FileSystemPlayerStore) recordWin(name string) {
	league := f.getLeague()

	for i, player := range league {
		if player.Name == name {
			league[i].Wins++
		}
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}

func (f *FileSystemPlayerStore) getLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}
