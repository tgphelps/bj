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

	d := newDealer(s, true) // hit soft 17
	d.getHand()
	fmt.Printf("Initial hand: %s\n", d.hand)
	if d.hand.value != 12 {
		t.Error("bad initial hand")
	}
	d.playHand()
	if d.hand.value != 20 {
		t.Error("bad final hand")
	}
	fmt.Printf("After playing hand: %s\n", d.hand)
}
