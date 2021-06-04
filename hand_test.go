package main

import (
	"fmt"
	"testing"
)

func TestHand(t *testing.T) {
	fmt.Println("Testing Hand")
	s := newShoe(2)
	s.shuffle()
	h := newHand(s, 2)
	fmt.Printf("hand: %s\n", h)
	if h.isPair() {
		t.Error("Hand is NOT a pair")
	}
	fmt.Println("hit...")
	h.hit()
	fmt.Printf("hand: %s\n", h)
	if h.cards[0] != 10 || h.cards[1] != 2 || h.cards[2] != 8 {
		t.Error("Bad initial cards for hand")
	}
	if h.betAmount != 2 {
		t.Error("Bad initial bet amount")
	}
	if h.blackjack || h.obsolete || h.busted {
		t.Error("Hand attributes not false")
	}
	h.double()
	if h.betAmount != 4 {
		t.Error("Bad doubled bet amount")
	}
	if h.isSoft() {
		t.Error("Hand is NOT soft")
	}
	fmt.Printf("hand: %s\n", h)
}
