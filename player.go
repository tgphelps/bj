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
		// Play a hand.
		switch {
		case p.shouldSurrender(h):
			log.Panic("not implemented")
		case p.shouldSplit(h):
			log.Panic(("not implemented"))
		case p.shouldDouble(h):
			log.Panic("not implemented")
		default:
			p.playNormal(h, upcard)
		}
	}
}

// Throws away all hands just played, and gets ready for another round.

func (p *Player) endRound() {
	p.hands = nil
}

func (p *Player) shouldSurrender(h *Hand) bool {
	return false
}

func (p *Player) shouldSplit(h *Hand) bool {
	return false
}

func (p *Player) shouldDouble(h *Hand) bool {
	return false
}

// playNormal plays a hand using only the hit/stand strategy.

func (p *Player) playNormal(h *Hand, upcard int8) {
	var s StrPoint
	var busted bool
	for {
		if h.isSoft() {
			s = StrPoint{keyHitSoft, h.value, upcard}
			busted = !p.playStrategy(s, h)
			if busted {
				break
			}
		} else {
			s = StrPoint{keyHitHard, h.value, upcard}
			busted = !p.playStrategy(s, h)
			if busted {
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
		log.Printf("play: HIT. Hand: %s\n", h)
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
