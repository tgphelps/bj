package main

import (
	"testing"
)

func TestDealer(t *testing.T) {
	s := newShoe(1)
	// s.randomize()
	s.shuffle()
	d := newDealer(s, true) // hit soft

	d.getHand()
	if d.hand.value != 17 {
		t.Error("dealer should have 17")
	}
	if d.upCard() != 10 {
		t.Error("dealer upcard should be 10")
	}
	d.playHand()
	if d.hand.value != 17 {
		t.Error("after dealer play, should have 17")
	}

	s.force([]int8{10, 11})
	d.getHand()
	if !d.hand.isBlackjack {
		t.Error("dealer should have blackjack")
	}
}
func TestDealerSoft17(t *testing.T) {
	s := newShoe(1)
	d := newDealer(s, true)
	s.force([]int8{11, 6, 3})
	d.getHand()
	d.playHand()
	if d.hand.value != 20 {
		t.Error("dealer should have hit S17")
	}
	d = newDealer(s, false)
	s.force([]int8{11, 6, 3})
	d.getHand()
	if d.hand.value != 17 {
		t.Error("dealer should not have hit S17")
	}
}
