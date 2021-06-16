package main

// "fmt"

type Game struct {
	verbose      bool
	numPlayers   int
	players      []*Player
	strategy     Strategy
	numDecks     int
	shufflePoint int
	st           Statistics
	shoe         *Shoe
	hitS17       bool
}
