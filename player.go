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
