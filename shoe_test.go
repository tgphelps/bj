package main

import (
	"testing"
)

func TestShoeBasic(t *testing.T) {
	s := newShoe(2)
	if s.shoeSize != 104 {
		t.Error("Bad initial `card count")
	}
	s.shuffle()

	if s.remaining != 2*52 {
		t.Error("Bad remaining() 1")
	}

	_ = s.deal()
	_ = s.deal()
	_ = s.deal()
	_ = s.deal()

	if s.remaining != 2*52-4 {
		t.Error("Bad remaining() 2")
	}

	// Deal all the cards in the shoe.
	for i := 0; i < 100; i++ {
		_ = s.deal()
	}

	if s.remaining != 0 {
		t.Error("Bad remaining() 3")
	}
}

func TestShoeDeal(t *testing.T) {
	s := newShoe(1)
	c1 := s.deal()
	c2 := s.deal()
	if c1 != 2 || c2 != 3 {
		t.Error("Should have dealt 2 and 3")
	}
	for i := 0; i < 49; i++ {
		_ = s.deal()
	}
	c := s.deal()
	if c != 11 {
		t.Error("last card should be 11")
	}
	s.shuffle()
	if s.next != 0 {
		t.Errorf("after shuffle 'next' was %d", s.next)
	}
	if s.remaining != 52 {
		t.Errorf("after shuffle 'remaining' was %d'", s.remaining)
	}
	if s.shoeSize != 52 {
		t.Errorf("after shutfle shoeSize was %d", s.shoeSize)
	}
}

func TestShoeShuffle(t *testing.T) {
	s := newShoe(1)
	s.shuffle()
	c1 := s.deal()
	c2 := s.deal()
	c3 := s.deal()
	// fmt.Printf("after shuffle: %d %d %d\n", c1, c2, c3)
	if c1 != 10 || c2 != 7 || c3 != 6 {
		t.Errorf("after shuffle, no 10,7,6, but %d,%d,%d", c1, c2, c3)
	}
}
