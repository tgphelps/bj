package main

import (
	"testing"
)

// Cards dealt from one shuffled deck are:
// 10 7 6 3 9 8 8 10 10 6 4 3 5 9 10 3 9 2 4 7

func TestHand(t *testing.T) {
	s := newShoe(1)
	s.shuffle()
	h := newHand(s, 2)
	// should be 10 7
	if h.betAmount != 2 {
		t.Errorf("bad bet amount: %d", h.betAmount)
	}
	if h.value != 17 {
		t.Errorf("bad value: %d", h.value)
	}
	if h.cards[0] != 10 || h.cards[1] != 7 {
		t.Errorf("bad cards: %d %d", h.cards[0], h.cards[1])
	}
	s.force([]int8{10, 11})
	h = newHand(s, 2)
	if h.value != 21 {
		t.Errorf("value should be 21: %d", h.value)
	}
	s.force([]int8{11, 11})
	h = newHand(s, 2)
	if h.value != 12 {
		t.Errorf("value should be 12: %d", h.value)
	}
	if !h.isSoft() {
		t.Error("hand should be soft")
	}
	if !h.isPair() {
		t.Error("hand should be a pair")
	}
}

func TestHandHitting(t *testing.T) {
	s := newShoe(1)
	s.shuffle()
	h := newHand(s, 2)
	// should be 10 7
	h.double()
	if !h.isDoubled {
		t.Error("hand should be doubled")
	}
	if !h.isBusted {
		t.Error("hand should be busted")
	}
	if h.betAmount != 4 {
		t.Error("bet amount should be 4")
	}
	h = newHand(s, 2)
	// Should be 9, 3
	h.hit()
	if h.value != 20 {
		t.Error("value should be 20")
	}
	h.hit()
	if h.value <= 21 {
		t.Error("value should be > 21")
	}
	if !h.isBusted {
		t.Error("hand should be busted")
	}
}

func TestSplitHand(t *testing.T) {
	s := newShoe(1)
	s.shuffle()
	h := newSplitHand(s, 2, 10)
	if h.value != 20 {
		t.Error("value should be 20")
	}
	if !h.isSplit {
		t.Error("hand should be split")
	}
	if !h.isPair() {
		t.Error("hand should be a pair")
	}
}

//func TestCards(t *testing.T) {
//	s := newShoe(1)
//	s.shuffle()
//	for i := 0; i < 20; i++ {
//		fmt.Println(s.deal())
//	}
//}
