package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Game struct {
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

// newGame creates a new Game. We may seed the random number generator.
// We shuffle the Shoe and create a new Dealer. Finally, we create some
// number of new Players for the game.

func newGame(strategy Strategy, numPlayers int, penetration int, repeatable bool, cfg *Config) *Game {
	var g Game

	g.strategy = strategy
	g.cfg = cfg
	g.shufflePoint = cfg.numDecks * 52 * penetration / 100
	g.hitS17 = cfg.hitS17
	g.shoe = newShoe(cfg.numDecks)
	g.numPlayers = numPlayers
	if !repeatable {
		// fmt.Println("randomize")
		g.shoe.randomize()
	}
	// fmt.Println("initial shuffle")
	g.shoe.shuffle()
	g.dealer = newDealer(g.shoe, g.hitS17)
	for i := 0; i < numPlayers; i++ {
		// fmt.Printf("create new player...\n")
		p := newPlayer(i+1, g.shoe, g.cfg, g.strategy, betAmount)
		g.players = append(g.players, p)
		// fmt.Printf("XXX player: %v\n", p)
	}
	return &g
}

// playRound:
// Shuffle if we've hit the cut card.
// Deal hands to each player and the dealer.
// If the dealer has a BJ, settle all hands now.
// Otherwise, play each player hand, and then the dealer hand.
// Collect data on win/loss/push.

func (g *Game) playRound() {
	if g.shoe.remaining < g.shufflePoint {
		log.Printf("remaining: %d => shuffle", g.shoe.remaining)
		g.shoe.shuffle()
	}
	for _, p := range g.players {
		p.getHand()
		g.st.handsPlayed += 1
		log.Printf("P%d hand: %s\n", p.seat, p.hands[0])
		if p.hands[0].isBlackjack {
			log.Printf("P%d blackjack\n", p.seat)
		}
	}
	g.dealer.getHand()
	log.Printf("dealer hand: %s\n", g.dealer.hand)
	if g.dealer.hand.isBlackjack {
		log.Println("dealer BLACKJACK")
	} else {
		for _, p := range g.players {
			log.Printf("Play seat %d:\n", p.seat)
			p.playHands(g.dealer.upCard())
		}
		g.dealer.playHand()
	}
	g.updateStats()
	log.Println("clear player hands")
	for _, p := range g.players {
		p.endRound()
	}
	g.st.roundsPlayed += 1
}

// updateStats updates our Statistics struct with the betting results
// of each player hand.

func (g *Game) updateStats() {
	dlrVal := g.dealer.hand.value
	dlrBj := g.dealer.hand.isBlackjack
	dlrBust := g.dealer.hand.isBusted

	log.Println("----------RESULTS")
	log.Printf("   Dealer has %d\n", dlrVal)
	for _, p := range g.players {
		for n, h := range p.hands {
			// n = hand number, h = hand struct
			log.Printf("   P%d hand %d: %d\n", p.seat, n+1, h.value)
			if h.isObsolete {
				log.Println("   OBSOLETE")
				continue
			}
			// Hand is not obsolete
			g.st.totalBet += h.betAmount
			if h.isBlackjack && !dlrBj {
				g.st.blackjacksWon += 1
				win := (3 * h.betAmount / 2)
				g.st.totalWon += win
				log.Printf("   WIN %d: BJ\n", win)

				continue
			}
			// This hand did not win with a blackjack
			if dlrBj {
				if h.isBlackjack {
					g.st.totalPush += h.betAmount
					log.Println("   PUSH: blackjacks")
				} else {
					g.st.totalLost += h.betAmount
					fmt.Printf("LOSE %d: dealer BJ\n", h.betAmount)
				}
			} else {
				// Nobody had a blackjack
				if h.isBusted {
					g.st.totalLost += h.betAmount
					log.Printf("   LOSE %d: bust\n", h.betAmount)
				} else if dlrBust {
					g.st.totalWon += h.betAmount
					log.Printf("   WIN %d: dealer bust\n", h.betAmount)
				} else if dlrVal > h.value {
					g.st.totalLost += h.betAmount
					log.Printf("   LOSE %d\n", h.betAmount)
				} else if h.value > dlrVal {
					g.st.totalWon += h.betAmount
					log.Printf("   WIN %d\n", h.betAmount)
				} else {
					g.st.totalPush += h.betAmount
					log.Println("   PUSH")
				}
			}
		}
	}
	fmt.Println("----------END")
}

// writeStats is called after all rounds have been played. It appends
// a set of lines to the statistics file, that show the results of all
// the rounds we just played.

func (g *Game) writeStats(fileName string, strategyName string) {
	var gain float32

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("FATAL: ", err)
	}

	defer f.Close()
	// XXX The following assumes a blackjack pays 3:2
	if g.st.totalWon+g.st.totalLost+g.st.totalPush-g.st.blackjacksWon != g.st.totalBet {
		fmt.Fprintln(f, "ERROR: stats don't match")
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
		gain = float32(100*(g.st.totalWon-g.st.totalLost)) / float32(g.st.totalBet)
		fmt.Fprintf(f, "pct win = %6.4f\n", gain)
	}
}
