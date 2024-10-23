package director

import (
	"fmt"
	"matchmaker/matchmaker"
	"time"

	"github.com/google/uuid"
)

type Options struct {
	MaxPlayers    int
	MaxWait       time.Duration
	FetchInterval time.Duration
}

func NewDirector(matchmaker *matchmaker.Matchmaker, options Options) {
	go run(matchmaker, options)
}

func run(matchmaker *matchmaker.Matchmaker, options Options) {
	for {
		matches := matchmaker.GetCompetitions(options.MaxPlayers, options.MaxWait)

		fmt.Println("Matches", len(matches))
		if len(matches) > 0 {
			for _, match := range matches {
				// Request leaderboard UUID and assign that to players
				competition := createCompetition()
				fmt.Println("Competition", competition)
				for _, player := range match.Players {
					player.Competition = competition
					// Note: For simplicity we are notifying the player directly from here.
					player.Notify()
				}
			}
		}
		time.Sleep(time.Duration(options.FetchInterval))
	}
}

func createCompetition() string {
	return uuid.New().String()
}
