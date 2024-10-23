package matchmaker

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Player struct {
	ID          string
	Level       int
	Country     string
	Competition string
	queueJoined time.Time
	connection  string //Note: For example saving some connection info about the player on how to communicate back
}

func (p *Player) Notify() {
	fmt.Println("Notifying player", p.ID, p.Country, p.Level, "about competition", p.Competition, "on connection", p.connection)
}

func (p *Player) validate() bool {
	return p != nil && p.Level > 0 && p.Level <= 100 && p.Country != ""
}

var COUNTRIES = [...]string{"US", "UK", "FR", "DE", "ES", "IT", "JP", "CN"}

func CreatePlayer(connection string) *Player {
	return &Player{
		ID:         uuid.New().String(),
		Level:      rand.Intn(100) + 1,
		Country:    COUNTRIES[rand.Intn(len(COUNTRIES))],
		connection: connection,
	}
}
