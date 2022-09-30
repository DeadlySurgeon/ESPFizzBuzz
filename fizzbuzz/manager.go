package fizzbuzz

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// GameManager manages the various games that are played.
type GameManager struct {
	sync.Mutex
	Games map[string]*Game
}

// Game is the main entrypoint for playing this game.
func (gm *GameManager) Game(rw http.ResponseWriter, r *http.Request) {

	var id, cmd, entry string

	if id = param(rw, r, "id"); id == "" {
		return
	}

	if cmd = param(rw, r, "cmd"); id == "" {
		return
	}

	if entry = param(rw, r, "entry"); id == "" {
		return
	}

	gm.Lock()
	defer gm.Unlock()

	switch cmd {
	case "new":
		delete(gm.Games, id)
		gm.Games[id] = &Game{LastEntry: time.Now()}

		resp := gm.Games[id].Generate()

		rw.Write([]byte(resp))
		fmt.Printf("%v started a new game\n", id)
		fmt.Printf("%v\t%v\n%v\n", id, 1, resp)

	case "submit":
		if entry == "incorrect" {
			// Handle if we're called out. Have a random chance of throwing out
			// a bad move to ensure that we keep the client on it's toes.
		}

		game := gm.Games[id]
		if game == nil {
			// Game isn't set up right.
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		if entry == "startover" {
			game.count = 0
			game.LastEntry = time.Now()

			fmt.Printf("%v decided to start over\n", id)

			resp := game.Generate()
			
			fmt.Printf("%v\t%v\n%v\n", id, 1, resp)
			fmt.Print(rw, resp)

			return
		}

		if resp := game.Verify(entry); resp != "" {
			fmt.Printf(
				"%v gave a wrong response. Expected %v got %v\n",
				id,
				resp,
				entry,
			)

			rw.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(rw, "Expected %v got %v\n", resp, entry)
			return
		}

		if game.count+1 == 500 {
			// END Game, start from zero.
			// Send "startover"
			// Set game count to 0
			// Expect client to start at 1.
		}

		next := game.Generate()
		// Give next one.
		fmt.Fprint(rw, next)

		fmt.Printf("%v\t%v\t%v\n", id, game.count-1, entry)
		fmt.Printf("%v\t%v\t%v\n", id, game.count, next)

	default:
		delete(gm.Games, id)

		fmt.Printf("%v gave bad cmd\n", id)
		fmt.Fprintf(rw, "Bad CMD: %v", cmd)
		rw.WriteHeader(http.StatusBadRequest)
	}

}

func param(rw http.ResponseWriter, r *http.Request, key string) string {
	value := r.URL.Query().Get(key)
	if value == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return ""
	}
	return value
}
