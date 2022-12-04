package main

import (
	"log"
)

// Type Player represents a player who occupies a seat, and plays one or
// more hands against the dealer.

type Player struct {
	seat      int      // Table seat this player occupies
	shoe      *Shoe    // Shoe from which we get cards
	cfg       *Config  // House rules in effect
	strategy  Strategy // Strategy we should follow
	betAmount int      // Amount to bet on each hand
	// splitCount int
	hands []*Hand // Hands we are playing now
}

// newPlayer creates a new Player

func newPlayer(seat int, shoe *Shoe, cfg *Config, strategy Strategy, betAmount int) *Player {
	var p Player
	// fmt.Println("in newPlayer")
	p.seat = seat
	p.shoe = shoe
	p.cfg = cfg
	p.strategy = strategy
	p.betAmount = betAmount
	return &p
}

// getHand gets one new 2-card hand

func (p *Player) getHand() {
	h := newHand(p.shoe, p.betAmount)
	if h.value == 21 {
		h.isBlackjack = true
	}
	p.hands = append(p.hands, h)
}

// getSplitHand get one new card to go with one of a split pair

func (p *Player) getSplitHand(firstCard int8) {
	h := newSplitHand(p.shoe, p.betAmount, firstCard)
	if firstCard == 11 && !p.cfg.canHitSplitAces {
		h.hitNotAllowed = true
	}
	if !p.cfg.dasAllowed {
		h.doubleNotAllowed = true
	}
	p.hands = append(p.hands, h)
}

// playHands plays all the hands that a player has. During the play, more
// hands can be created if pairs are split. We must check for the 'exotic'
// options surrender, split, and double, in that order, before doing the basic
// 'hit' strategy.

func (p *Player) playHands(upcard int8) {
	for _, h := range p.hands {
		if !p.maybeSurrender(h, upcard) {
			if !p.maybeSplit(h, upcard) {
				if !p.maybeDouble(h, upcard) {
					p.playNormal(h, upcard)
					log.Printf("P%d final hand: %s\n", p.seat, h)
				}
			}
		}
	}
}

// Throws away all hands just played, and gets ready for another round.

func (p *Player) endRound() {
	p.hands = nil
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
	if h.isSplit && !p.cfg.dasAllowed {
		log.Println("DAS not allowed")
		return false
	}
	if !h.isSoft() {
		// Double hard hand?
		return false
	} else {
		// Double soft hand?
		return false
	}
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
// It returns true if we hit AND we don't bust, else it returns false.
// If it returns false, the caller will not need to do any more with this hand.

func (p *Player) playStrategy(s StrPoint, h *Hand) bool {
	var ret bool
	if p.strategy[s] {
		h.hit()
		log.Printf("HIT. Hand: %s\n", h)
		if h.isBusted {
			ret = false
		} else {
			ret = true
		}
	} else {
		ret = false
	}
	return ret
}
