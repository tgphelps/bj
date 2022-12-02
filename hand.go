package main

import (
	// "fmt"
	"log"
	"strconv"
	"strings"
)

type Hand struct {
	blackjack bool
	doubled   bool
	busted    bool
	isSplit   bool
	noHit     bool
	noDouble  bool
	obsolete  bool
	shoe      *Shoe
	betAmount int
	value     int8
	bigAces   int
	cards     []int8
}

const ace = 11
const softAce = 1

// finishHAND is the common code that needs to run for all new hands,
// both inital and split hands.

func (h *Hand) finishHand() {
	h.updateValue()
	if h.value == 21 {
		h.blackjack = true
	}
	if h.bigAces == 2 {
		// This happens if the hand is ace-ace.
		h.harden()
	}
}

// newHand creates a new hand and deals its first 2 cards. This deals both
// player and dealer hands. The caller should set betAmount to 0 for a dealer
// hand.

func newHand(s *Shoe, betAmount int) *Hand {
	var h Hand
	h.shoe = s
	h.betAmount = betAmount
	h.cards = append(h.cards, s.deal())
	h.cards = append(h.cards, s.deal())
	h.finishHand()
	return &h
}

// newSplitHand creates a new hand containing one card of a split pair
// and one new card from the shoe.

func newSplitHand(s *Shoe, betAmount int, firstCard int8) *Hand {
	var h Hand
	h.shoe = s
	h.betAmount = betAmount
	h.cards = append(h.cards, firstCard)
	h.cards = append(h.cards, s.deal())
	h.isSplit = true
	h.finishHand()
	h.blackjack = false // No blackjacks on split pairs
	return &h
}

// String creates a printable string to represent a hand,
// so Hand can be used in fmt.Printf with the "%s" specification.

func (h *Hand) String() string {
	var sb strings.Builder

	sb.WriteString("{")
	sb.WriteString(strconv.Itoa(int(h.value)))
	sb.WriteString(" |")
	for _, s := range h.cards {
		sb.WriteString(" ")
		sb.WriteString(strconv.Itoa(int(s)))
	}
	sb.WriteString(" | ")
	writeFlag(&sb, h.blackjack, "J")
	writeFlag(&sb, h.busted, "B")
	writeFlag(&sb, h.doubled, "D")
	writeFlag(&sb, h.isSplit, "S")
	writeFlag(&sb, h.obsolete, "O")
	sb.WriteString("}")

	return sb.String()
}

// writeFlag is called only from the String function above.

func writeFlag(sb *strings.Builder, f bool, c string) {
	if f {
		sb.WriteString(c)
	}
}

// updateValue updates the 'value' field in the hand with total of all cards
// in the hand.

func (h *Hand) updateValue() {
	var sum int8
	for _, n := range h.cards {
		sum += int8(n)
	}
	h.value = sum
	h.bigAces = countCard(ace, h.cards)
}

// isSoft returns true if the hand is 'soft', i.e., there are one or more
// aces in the hand that, when counted as 1, brings the total under 21.

func (h *Hand) isSoft() bool {
	return h.bigAces > 0
}

// double doubles the player's bet and deals one more card.

func (h *Hand) double() {
	h.betAmount *= 2
	h.doubled = true
	h.hit() // XXX should we do this here?
}

// harden changes the value of the first ace in the hand from 11 to 1.
// It will crash if called with a hand that isn't soft.

func (h *Hand) harden() {
	firstAce := findCard(ace, h.cards)
	h.cards[firstAce] = softAce
	h.updateValue()
}

// isPair returns true if the hand contains a pair.

func (h *Hand) isPair() bool {
	if len(h.cards) != 2 {
		log.Panic("isPair: not 2-card hand")
	}
	if h.cards[0] == h.cards[1] {
		return true
	}
	if h.cards[0] == ace && h.cards[1] == softAce {
		return true
	}
	if h.cards[0] == softAce && h.cards[1] == ace {
		return true
	}
	return false
}

// hit deals one more card to the hand, updates various flags.

func (h *Hand) hit() {
	if h.obsolete {
		log.Panic("hand.hit: obsolete")
	}
	c := h.shoe.deal()
	h.cards = append(h.cards, c)
	h.updateValue()
	if c == ace {
		h.bigAces += 1
	}
	if h.value > 21 {
		if h.isSoft() {
			h.harden()
		} else {
			h.busted = true
		}

	}
}

// utility function to find the first instance of a given card in the hand.
// It panics if no such card is found.

func findCard(card int8, hand []int8) int {
	for i, c := range hand {
		if c == card {
			return i
		}
	}
	log.Panic("find_card no find")
	return 0 // not reached
}

// utility function to count the number of a given card in the hand.

func countCard(card int8, hand []int8) int {
	count := 0
	for _, c := range hand {
		if c == card {
			count += 1
		}
	}
	return count
}
