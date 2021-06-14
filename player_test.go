package main

import (
	"fmt"
	"testing"
)

func TestPlayer(t *testing.T) {
	fmt.Println("test player...")
	var cfg Config
	var strategy Strategy
	shoe := newShoe(6)
	shoe.randomize()
	shoe.shuffle()
	p := newPlayer(1, shoe, &cfg, strategy, 2, false)
	fmt.Printf("got new player: %v\n", p)
	p.endRound()
	p.getHand()
	fmt.Printf("New hand: %s\n", p.hands[0])
	p.getHand()
	fmt.Printf("New hand: %s\n", p.hands[1])
	fmt.Println("get split hand to a 10")
	p.getSplitHand(10)
	fmt.Printf("split hand: %s\n", p.hands[2])
}
