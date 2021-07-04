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
	if g.dealer.hand.blackjack {
		if g.verbose {
			fmt.Println("dealer BLACKJACK")
		}
	} else {
		for _, p := range g.players {
			fmt.Printf("Play seat %d:\n", p.seat)
			p.playHands(g.dealer.upCard())
		}
	}
	g.updateStats()
	if g.verbose {
		fmt.Println("clear player hands")
	}
	for _, p := range g.players {
		p.endRound()
	}
}

func (g *Game) updateStats() {
	dlrVal := g.dealer.hand.value
	dlrBj := g.dealer.hand.blackjack
	dlrBust := g.dealer.hand.busted
	if g.verbose {
		fmt.Println("----------RESULTS")
		fmt.Printf("   Dealer has %d\n", dlrVal)
	}
	for _, p := range g.players {
		for n, h := range p.hands {
			// n = hand number, h = hand struct
			if g.verbose {
				fmt.Printf("   P%d hand %d: %d\n", p.seat, n+1, h.value)

			}
			if h.obsolete {
				if g.verbose {
					fmt.Println("   OBSOLETE")
				}
				continue
			}
			// Hand is not obsolete
			g.st.totalBet += h.betAmount
			if h.blackjack && !dlrBj {
				g.st.blackjacksWon += 1
				win := (3 * h.betAmount / 2)
				g.st.totalWon += win
				if g.verbose {
					fmt.Printf("   WIN %d: BJ\n", win)
				}

				continue
			}
			// This hand did not win with a blackjack
			if dlrBj {
				if h.blackjack {
					g.st.totalPush += h.betAmount
					if g.verbose {
						fmt.Println("   PUSH: blackjacks")
					}
				} else {
					g.st.totalLost += h.betAmount
					if g.verbose {
						fmt.Printf("LOSE %d: dealer BJ: %d\n", h.betAmount)
					}
				}
			} else {
				// Nobody had a blackjack
				if h.busted {
					g.st.totalLost += h.betAmount
					if g.verbose {
						fmt.Printf("   LOSE %d: bust\n", h.betAmount)
					}
				} else if dlrBust {
					g.st.totalWon += h.betAmount
					if g.verbose {
						fmt.Printf("   WIN %d: dealer bust\n", h.betAmount)
					}
				} else if dlrVal > h.value {
					g.st.totalLost += h.betAmount
					if g.verbose {
						fmt.Printf("   LOSE %d\n", h.betAmount)
					}
				} else if h.value > dlrVal {
					g.st.totalWon += h.betAmount
					if g.verbose {
						fmt.Printf("   WIN %d\n", h.betAmount)
					}
				} else {
					g.st.totalPush += h.betAmount
					if g.verbose {
						fmt.Println("   PUSH")
					}
				}
			}
		}
	}
	if g.verbose {
		fmt.Println("----------END")
	}
}

func (g *Game) writeStats(fileName string, strategyName string) {
	var gain float32
	// XXX Stats file must already exist
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
