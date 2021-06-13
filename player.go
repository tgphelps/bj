package main

type Player struct {
	seat        int
	shoe        *Shoe
	cfg         *Config
	strategy    Strategy
	bet_amount  int
	verbose     bool
	splits_done int
	hands       []*Hand
}

func newPlayer(seat int, shoe *Shoe, cfg *Config, strategy Strategy, bet_amount int, verbose bool) *Player {
	var p Player
	p.seat = seat
	p.shoe = shoe
	p.cfg = cfg
	p.strategy = strategy
	p.bet_amount = bet_amount
	p.verbose = verbose
	return &p
}

// Player methods

func (p *Player) logHands() {
	panic("not yet")
}

func (p *Player) getHand() {
	panic("not yet")
}

func (p *Player) getSplitHand() {
	panic("not yet")
}

func (p *Player) playHands(upcard int8) {
	panic("not yet")
}

func (p *Player) endRound() {
	panic("not yet")
}

func (p *Player) maybeSurrender(h *Hand, upcard int8) bool {
	if true {
		panic("not yet")
	}
	return false
}

func (p *Player) maybeDouble(h *Hand, upcard int8) bool {
	if true {
		panic("not yet")
	}
	return false
}

func (p *Player) maybeSplit(h *Hand, upcard int8) bool {
	if true {
		panic("not yet")
	}
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
