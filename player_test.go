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
