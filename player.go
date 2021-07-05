package main

import (
	"fmt"
)

// Type Player represents a player who occupies a seat, and plays one or
// more hands against the dealer.

type Player struct {
	seat       int
	shoe       *Shoe
	cfg        *Config
	strategy   Strategy
	betAmount  int
	verbose    bool
	splitCount int
	hands      []*Hand
}

// newPlayer creates a new Player
func newPlayer(seat int, shoe *Shoe, cfg *Config, strategy Strategy, betAmount int, verbose bool) *Player {
	var p Player
	// fmt.Println("in newPlayer")
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

// getHand gets one new 2-card hand
func (p *Player) getHand() {
	p.hands = append(p.hands, newHand(p.shoe, p.betAmount))
}

// getSplitHand get one new card to go with one of a split pair
func (p *Player) getSplitHand(firstCard int8) {
	p.hands = append(p.hands, newSplitHand(p.shoe, p.betAmount, firstCard))
}

// playHands plays all the hands that a player has. During the play, more
// hands can be created if pairs are split.
func (p *Player) playHands(upcard int8) {
	for _, h := range p.hands {
		if !p.maybeSurrender(h, upcard) {
			if !p.maybeSplit(h, upcard) {
				if !p.maybeDouble(h, upcard) {
					p.playNormal(h, upcard)
					if p.verbose {
						fmt.Printf("P%d final hand: %s\n", p.seat, h)
					}
				}
			}
		}
	}
}

// Throws away all hands just played, and gets ready for another round.
func (p *Player) endRound() {
	p.hands = nil
	p.splitCount = 0
}

// maybeSurrender will surrender, if the strategy says to do so. If it
// does surrender, it returns true, else false.
func (p *Player) maybeSurrender(h *Hand, upcard int8) bool {
	// XXX do this some day
	return false
}

// maybeDouble will double a hand, if the strategy says to do so. If it
// does double, it returns true, else false.
func (p *Player) maybeDouble(h *Hand, upcard int8) bool {
	// fmt.Println("XXX fix maybeDouble")
	return false
}

// maybeSplit will split a pair, if the strategy says to do so. If it
// does split, it returns true, else false.
func (p *Player) maybeSplit(h *Hand, upcard int8) bool {
	// fmt.Println("XXX fix maybeSplit")
	return false
}

// playNormal plays a hand using only the hit/stand strategy.
func (p *Player) playNormal(h *Hand, upcard int8) {
	var s StrPoint
	for {
		if h.isSoft() {
			s = StrPoint{keyHitSoft, h.value, upcard}
			if !p.playStrategy(s, h) {
				break
			}
		} else {
			s = StrPoint{keyHitHard, h.value, upcard}
			if !p.playStrategy(s, h) {
				break
			}
		}
	}
}

// playStrategy will hit the hand, if the strategy says to do so.
// It returns true if we do hit AND we don't bust. Else it returns false.
// If we return false, the caller will not need to do any more with this hand.
func (p *Player) playStrategy(s StrPoint, h *Hand) bool {
	var ret bool
	if p.strategy[s] {
		h.hit()
		if p.verbose {
			fmt.Printf("HIT. Hand: %s\n", h)
		}
		if h.busted {
			ret = false
		} else {
			ret = true
		}
	} else {
		ret = false
	}
	return ret
}
