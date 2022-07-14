package http_server

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, getPlayerScore(player))
}

func getPlayerScore(playerName string) int {
	switch playerName {
	case "Floyd":
		return 10
	case "Pepper":
		return 20
	}

	return 0
}
