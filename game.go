package main

// "fmt"

type Game struct {
	verbose      bool
	numPlayers   int
	players      []*Player
	strategy     Strategy
	cfg          *Config
	shufflePoint int
	st           Statistics
	shoe         *Shoe
	hitS17       bool
}

func newGame(strategy Strategy, numPlayers int, penetration int, repeatable bool, cfg *Config, verbose bool) *Game {
	var g Game

	g.verbose = verbose
	g.strategy = strategy
	g.cfg = cfg
	g.shufflePoint = cfg.numDecks * 52 * penetration / 100
	g.hitS17 = cfg.hitS17
	g.shoe = newShoe(cfg.numDecks)
	if !repeatable {
		g.shoe.randomize()
	}
	g.shoe.shuffle()
	g.players = nil
	// XXX more to do
	return &g
}
