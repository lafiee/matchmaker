package main

import (
	"fmt"
	"matchmaker/director"
	"matchmaker/matchmaker"
	"net/http"
	"time"
)

func main() {
	mm := matchmaker.NewMatchmaker()

	// Director works as a stateful service requesting competitions. Could be separate service (Join volumes might be too big to check this after every join?)
	directorOptions := director.Options{MaxPlayers: 10, MaxWait: time.Second * 30, FetchInterval: time.Second * 5}
	director.NewDirector(mm, directorOptions)

	http.HandleFunc("/join", handleJoin(mm))
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func handleJoin(mm *matchmaker.Matchmaker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Note: The post request is empty, we are not using the body, but generate a random player for simulation
		player := matchmaker.CreatePlayer(r.Host)

		// Note: We could write to database here, but for simplicity we add to slice protected by mutex.
		err := mm.JoinMatchmaking(player)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("Player joined", player.ID, player.Level, player.Country)
	}
}
