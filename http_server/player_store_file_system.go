package http_server

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
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
	//TODO implement me
	panic("implement me")
}

func (f *FileSystemPlayerStore) getLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}
