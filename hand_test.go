package main

import (
	"fmt"
	"testing"
)

func TestHand(t *testing.T) {
	fmt.Println("Testing Hand")
	s := newShoe(2)
	s.randomize()
	s.shuffle()
	h := newHand(s, 2)
	fmt.Printf("hand: %s\n", h)
	if h.isPair() {
		t.Error("Hand is NOT a pair")
	}
	fmt.Println("hit...")
	h.hit()
	fmt.Printf("hand: %s\n", h)
	if h.betAmount != 2 {
		t.Error("Bad initial bet amount")
	}
	fmt.Println("doubling...")
	h = newHand(s, 2)
	h.double()
	if h.betAmount != 4 {
		t.Error("Bad doubled bet amount")
	}
	fmt.Printf("hand: %s\n", h)
}
