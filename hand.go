package main

type Hand struct {
	blackjack bool
	doubled   bool
	busted    bool
	isSplit   bool
	no_hit    bool
	noDouble  bool
	obsolete  bool
	shoe      Shoe
	betAmount int
	value     int8
	bigAces   int
	cards     []int8
}

func (h *Hand) updateValue() {
	var sum int8
	for _, n := range h.cards {
		sum += int8(n)
	}
	h.value = sum
}

func (h *Hand) isSoft() bool {
	return h.bigAces > 0
}

func (h *Hand) double() {
	h.betAmount *= 2
	h.doubled = true
	h.hit()
}

func (h *Hand) harden() {
	firstAce := findCard(ace, h.cards)
	h.cards[firstAce] = softAce
	h.updateValue()
}

func (h *Hand) isPair() bool {
	if len(h.cards) != 2 {
		panic("isPair: not 2-card hand")
	}
	if h.cards[0] == h.cards[1] {
		return true
	}
	if h.cards[0]+h.cards[1] == ace+softAce {
		return true
	}
	return false
}

func (h *Hand) hit() {
	if h.obsolete {
		panic("hand.hit: obsolete")
	}
	c := h.shoe.deal()
	h.cards = append(h.cards, c)
	h.updateValue()
	if c == ace {
		h.bigAces += 1
	}
	if h.value > 21 {
		if h.bigAces > 0 {
			h.harden()
		} else {
			h.busted = true
		}

	}
}

func findCard(card int8, hand []int8) int {
	for i, c := range hand {
		if c == card {
			return i
		}
	}
	panic("find_card no find")
}
