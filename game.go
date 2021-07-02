package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

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
	g.numPlayers = numPlayers
	if !repeatable {
		if verbose {
			fmt.Println("randomize")
		}
		g.shoe.randomize()
	}
	if verbose {
		fmt.Println("initial shuffle")
	}
	g.shoe.shuffle()
	g.dealer = newDealer(g.shoe, g.hitS17)
	for i := 0; i < numPlayers; i++ {
		// fmt.Printf("create new player...\n")
		p := newPlayer(i+1, g.shoe, g.cfg, g.strategy, betAmount, g.verbose)
		g.players = append(g.players, p)
		// fmt.Printf("XXX player: %v\n", p)
	}
	return &g
}

// Deal hands to each player and the dealer.
// If the dealer has a BJ, settle all hands now.
// Otherwise, play each player hand, and then the dealer hand.
// Collect data on win/loss/push.

func (g *Game) playRound() {
	fmt.Println("XXX Playing round...")
	if g.shoe.remaining < g.shufflePoint {
		if g.verbose {
			fmt.Println("shuffle")
		}
		g.shoe.shuffle()
	}
	for _, p := range g.players {
		p.getHand()
		if g.verbose {
			fmt.Printf("P%d hand: %s\n", p.seat, p.hands[0])
		}
		if p.hands[0].blackjack && g.verbose {
			fmt.Printf("P%d blackjack\n", p.seat)
		}
	}
	g.dealer.getHand()
	if g.verbose {
		fmt.Printf("dealer hand: %s\n", g.dealer.hand)
	}
	// if no dealer BJ, play player hands. Then play delaer hand.
	g.updateStats()
	if g.verbose {
		fmt.Println("clear player hands")
	}
	for _, p := range g.players {
		p.endRound()
	}
}

func (g *Game) updateStats() {
	fmt.Println("XXX Updating stats...")
}

func (g *Game) writeStats(fileName string, strategyName string) {
	var gain float32
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("FATAL: ", err)
	}
	defer f.Close()
	// XXX The following assumes a blackjack pays 3:2
	if g.st.totalWon+g.st.totalLost+g.st.totalPush-g.st.blackjacksWon != g.st.totalBet {
		fmt.Fprintln(f, "ERROR")
	}
	now := time.Now()
	fmt.Fprintln(f, "time", now.Format("2006-01-02 15:04:05"))
	fmt.Fprintf(f, "strategy %s\n", strategyName)
	fmt.Fprintf(f, "roundsPlayed %d\n", g.st.roundsPlayed)
	fmt.Fprintf(f, "handsPlayed %d\n", g.st.handsPlayed)
	fmt.Fprintf(f, "totalBet %d\n", g.st.totalBet)
	fmt.Fprintf(f, "totalWon %d\n", g.st.totalWon)
	fmt.Fprintf(f, "totalLost %d\n", g.st.totalLost)
	fmt.Fprintf(f, "totalPush %d\n", g.st.totalPush)
	fmt.Fprintf(f, "blackjacksWon %d\n", g.st.blackjacksWon)
	if g.st.totalBet > 0 {
		gain = float32(100 * (g.st.totalWon - g.st.totalLost) / g.st.totalBet)
		fmt.Fprintf(f, "pct win = %f5.4\n", gain)
	}
}
