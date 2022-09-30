package fizzbuzz

import (
	"fmt"
	"sync"
	"time"
)

// Game handles the fizzbuzz game.
type Game struct {
	// Ensure game safety, however racetime conditions will lose the game.
	sync.Mutex
	LastEntry time.Time
	count     int
}

// Verify verifies the next input
func (g *Game) Verify(input string) (expected string) {
	g.Lock()
	defer g.Unlock()

	g.count++
	expect := generate(g.count)

	if input != expect {
		return expect
	}

	return ""
}

// Generate will generate a new fizzbuzz.
func (g *Game) Generate() string {
	g.Lock()
	defer g.Unlock()
	
	g.count++

	return generate(g.count)
}

func generate(i int) string {
	s := fmt.Sprintf("%d", i)
	set := false

	if i%3 == 0 {
		s = "fizz"
		set = true
	}
	if i%5 == 0 {
		if !set {
			s = ""
		}
		s += "buzz"
	}

	return s
}
