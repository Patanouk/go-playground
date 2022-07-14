package http_server

type PlayerStore interface {
	getPlayerScore(name string) (int, bool)
	recordWin(name string)
	getLeague() []Player
}

type InMemoryPlayerStore struct {
	scores map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		scores: make(map[string]int),
	}
}

func (s *InMemoryPlayerStore) getPlayerScore(name string) (int, bool) {
	score, found := s.scores[name]
	return score, found
}

func (s *InMemoryPlayerStore) recordWin(name string) {
	s.scores[name]++
}

func (s *InMemoryPlayerStore) getLeague() []Player {
	return nil
}
