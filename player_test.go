package main

import (
	"testing"
)

func TestPlayer(t *testing.T) {
	var cfg Config
	var strategy Strategy
	s := newShoe(1)
	s.shuffle()
	s.force([]int8{10, 11})
	p := newPlayer(1, s, &cfg, strategy, 2)
	p.getHand()
	if !p.hands[0].isBlackjack {
		t.Error("player should have blackjack")
	}
	p.endRound()
	cfg.dasAllowed = false
	cfg.canHitSplitAces = false
	p.getSplitHand(11)
	h := p.hands[0]
	if !h.doubleNotAllowed {
		t.Error("double should not be allowed")
	}
	if !h.hitNotAllowed {
		t.Error("hit should not be allowed")
	}
}

func TestPlayNeverHit(t *testing.T) {
	var cfg Config
	var strategy Strategy // Empty strategy => never do anything.

	s := newShoe(1)
	s.shuffle()
	p := newPlayer(1, s, &cfg, strategy, 2)
	p.getHand()
	p.playHands(6) // play 17 against 6
	if p.hands[0].value != 17 {
		t.Error(("1 should not have hit"))
	}
	p.endRound()
	s.force([]int8{2, 3})
	p.getHand()
	p.playHands(10) // play 5 against 10
	v := p.hands[0].value
	if v != 5 {
		t.Errorf("value: %d should not have hit", v)
	}
}
