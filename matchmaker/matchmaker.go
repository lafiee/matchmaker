package matchmaker

import (
	"errors"
	"sort"
	"sync"
	"time"
)

type Matchmaker struct {
	mu      sync.Mutex // Note: For simplicity we are not using redis
	players []*Player
}

type Competition struct {
	Players []*Player
}

func NewMatchmaker() *Matchmaker {
	return &Matchmaker{
		mu:      sync.Mutex{},
		players: []*Player{},
	}
}

func (m *Matchmaker) JoinMatchmaking(player *Player) error {
	// Note: integrate into the server authentication to validate the request, and/or to fetch the actual player data from database
	if !player.validate() {
		return errors.New("invalid player: validation failed")
	}

	m.addPlayer(player)

	return nil
}

func (m *Matchmaker) GetCompetitions(maxPlayers int, maxWait time.Duration) []*Competition {
	if len(m.players) == 0 {
		return nil
	}

	var matches []*Competition

	// Apply matchmaking functions
	for players, ok := m.getPlayersFromQueue(maxPlayers); ok; players, ok = m.getPlayersFromQueue(maxPlayers) {
		matches = append(matches, &Competition{Players: players})
	}

	// Handle players who have been in the queue for more than the maxWait
	for expiredPlayers := m.getExpiredPlayers(maxWait); len(expiredPlayers) > 0; expiredPlayers = m.getExpiredPlayers(maxWait) {
		matches = append(matches, &Competition{Players: expiredPlayers})
	}

	return matches
}

func (m *Matchmaker) addPlayer(player *Player) {
	m.mu.Lock()
	defer m.mu.Unlock()
	player.queueJoined = time.Now()
	m.players = append(m.players, player)
}

func (m *Matchmaker) getPlayersFromQueue(maxPlayers int) ([]*Player, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.players) < maxPlayers {
		return nil, false
	}

	// To apply a matchmaking algorithm to match players together based on level and country
	// Note: For simplicity sort by country, then by level in descending order
	sort.Slice(m.players, func(i, j int) bool {
		// First compare by country
		if m.players[i].Country != m.players[j].Country {
			return m.players[i].Country < m.players[j].Country
		}
		// If countries are the same, compare by level (descending)
		return m.players[i].Level > m.players[j].Level
	})
	// Note: By no means is this a fair matchmaking algorithm. It could have buckets or dynamic level deltas.
	// And possibly some country factor.

	players := m.players[:maxPlayers]
	m.players = m.players[maxPlayers:]

	return players, true
}

func (m *Matchmaker) getExpiredPlayers(maxWait time.Duration) []*Player {
	m.mu.Lock()
	defer m.mu.Unlock()

	var expiredPlayers []*Player
	var remainingPlayers []*Player

	for _, player := range m.players {
		if time.Since(player.queueJoined) > maxWait {
			expiredPlayers = append(expiredPlayers, player)
		} else {
			remainingPlayers = append(remainingPlayers, player)
		}
	}

	// Update the matchmaker's player list to only include remaining players
	m.players = remainingPlayers
	return expiredPlayers
}
