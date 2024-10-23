package matchmaker

import (
	"testing"
)

func TestJoinMatchmaking(t *testing.T) {
	mm := NewMatchmaker()
	player := &Player{ID: "1", Level: 10, Country: "US"}

	err := mm.JoinMatchmaking(player)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(mm.players) != 1 {
		t.Errorf("Expected 1 player, got %d", len(mm.players))
	}
}

func TestGetExpiredPlayers(t *testing.T) {
	t.Run("No expired players", func(t *testing.T) {
		mm := NewMatchmaker()
		player := &Player{ID: "1", Level: 10, Country: "US"}
		mm.addPlayer(player)

		expiredPlayers := mm.getExpiredPlayers(10)
		if len(expiredPlayers) != 0 {
			t.Errorf("Expected 0 expired players, got %d", len(expiredPlayers))
		}
		if len(mm.players) != 1 {
			t.Errorf("Expected 1 player, got %d", len(mm.players))
		}
	})

	t.Run("Expired players", func(t *testing.T) {
		mm := NewMatchmaker()
		player1 := &Player{ID: "1", Level: 10, Country: "US"}
		mm.addPlayer(player1)

		expiredPlayers := mm.getExpiredPlayers(-1)
		if len(expiredPlayers) != 1 {
			t.Errorf("Expected 1 expired player, got %d", len(expiredPlayers))
		}
		if len(mm.players) != 0 {
			t.Errorf("Expected 0 players, got %d", len(mm.players))
		}
	})

	t.Run("Multiple expired players", func(t *testing.T) {
		mm := NewMatchmaker()
		player1 := &Player{ID: "1", Level: 10, Country: "US"}
		player2 := &Player{ID: "2", Level: 10, Country: "US"}
		mm.addPlayer(player1)
		mm.addPlayer(player2)

		expiredPlayers := mm.getExpiredPlayers(-1)
		if len(expiredPlayers) != 2 {
			t.Errorf("Expected 2 expired players, got %d", len(expiredPlayers))
		}
		if len(mm.players) != 0 {
			t.Errorf("Expected 0 players, got %d", len(mm.players))
		}
	})
}

func TestGetPlayersFromQueue(t *testing.T) {
	t.Run("Not enough players", func(t *testing.T) {
		mm := NewMatchmaker()
		player1 := &Player{ID: "1", Level: 10, Country: "US"}
		mm.addPlayer(player1)

		players, ok := mm.getPlayersFromQueue(2)
		if ok {
			t.Errorf("Expected not enough players, got %v", players)
		}
		if len(players) != 0 {
			t.Errorf("Expected 0 players, got %d", len(players))
		}
		if len(mm.players) != 1 {
			t.Errorf("Expected 1 player, got %d", len(mm.players))
		}
	})

	t.Run("Enough players", func(t *testing.T) {
		mm := NewMatchmaker()
		player1 := &Player{ID: "1", Level: 10, Country: "US"}
		player2 := &Player{ID: "2", Level: 10, Country: "US"}
		mm.addPlayer(player1)
		mm.addPlayer(player2)

		players, ok := mm.getPlayersFromQueue(2)
		if !ok {
			t.Errorf("Expected enough players, got %v", players)
		}
		if len(players) != 2 {
			t.Errorf("Expected 2 players, got %d", len(players))
		}
		if len(mm.players) != 0 {
			t.Errorf("Expected 0 players, got %d", len(mm.players))
		}
	})
}
