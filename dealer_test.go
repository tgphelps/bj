package main

import (
	"fmt"
	"testing"
)

func TestDealer(t *testing.T) {
	fmt.Println("Testing Dealer")
	s := newShoe(2)
	// s.randomize()
	s.shuffle()
	d := newDealer(s, true, false) // hit soft 17
	for i := 1; i <= 20; i++ {
		d.getHand()
		fmt.Printf("Initial hand: %s\n", d.hand)
		d.playHand()
		fmt.Printf("After playing hand: %s\n", d.hand)
	}
}
