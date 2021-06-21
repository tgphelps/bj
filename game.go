package main

import "fmt"

type Game struct {
	verbose      bool
	dealer       *Dealer
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
	g.dealer = newDealer(g.shoe, g.hitS17)
	for i := 0; i < numPlayers; i++ {
		// fmt.Printf("create new player...\n")
		p := newPlayer(i+1, g.shoe, g.cfg, g.strategy, betAmount, g.verbose)
		g.players = append(g.players, p)
		fmt.Printf("player: %v\n", p)
	}
	return &g
}

func (g *Game) playRound() {
	fmt.Println("Playing round...")
}

func (g *Game) updateStats() {
	fmt.Println("Updating stats...")
}

func (g *Game) writeStats() {

}
