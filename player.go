package main

import (
	"fmt"
)

type Player struct {
	seat       int
	shoe       *Shoe
	cfg        *Config
	strategy   Strategy
	betAmount  int
	verbose    bool
	splitsDone int
	hands      []*Hand
}

func newPlayer(seat int, shoe *Shoe, cfg *Config, strategy Strategy, betAmount int, verbose bool) *Player {
	var p Player
	p.seat = seat
	p.shoe = shoe
	p.cfg = cfg
	p.strategy = strategy
	p.betAmount = betAmount
	p.verbose = verbose
	return &p
}

// Player methods

func (p *Player) logHands() {
	panic("not yet")
}

func (p *Player) getHand() {
	p.hands = append(p.hands, newHand(p.shoe, p.betAmount))
}

func (p *Player) getSplitHand(firstCard int8) {
	p.hands = append(p.hands, newSplitHand(p.shoe, p.betAmount, firstCard))
}

func (p *Player) playHands(upcard int8) {
	panic("not yet")
}

func (p *Player) endRound() {
	p.hands = nil
	p.splitsDone = 0
}

func (p *Player) maybeSurrender(h *Hand, upcard int8) bool {
	fmt.Println("XXX fix maybeSurrender")
	return false
}

func (p *Player) maybeDouble(h *Hand, upcard int8) bool {
	fmt.Println("XXX fix maybeDouble")
	return false
}

func (p *Player) maybeSplit(h *Hand, upcard int8) bool {
	fmt.Println("XXX fix maybeSplit")
	return false
}

func (p *Player) playNormal(h *Hand, upcard int8) {
	panic("not yet")
}

func (p *Player) playStrategy(s StrPoint, h *Hand) bool {
	if true {
		panic("not yet")
	}
	return false
}
