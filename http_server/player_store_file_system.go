package http_server

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func (f *FileSystemPlayerStore) getPlayerScore(name string) (int, bool) {
	//TODO implement me
	panic("implement me")
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
