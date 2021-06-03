package main

import (
	"fmt"
	"testing"
)

func TestShoe1(t *testing.T) {
	var c int8
	s := newShoe(2)
	if s.shoeSize != 104 {
		t.Error("Bad initial `card count")
	}
	s.shuffle()
	fmt.Println(s.cards) // just to look

	if s.remaining() != 2*52 {
		t.Error("Bad remaining() 1")
	}

	c = s.deal()
	fmt.Printf("deal: %d\n", c)
	c = s.deal()
	fmt.Printf("deal: %d\n", c)
	c = s.deal()
	fmt.Printf("deal: %d\n", c)
	c = s.deal()
	fmt.Printf("deal: %d\n", c)

	if s.remaining() != 2*52-4 {
		t.Error("Bad remaining() 2")
	}

	// Deal all the cards in the shoe.
	for i := 0; i < 100; i++ {
		c = s.deal()
	}

	if s.remaining() != 0 {
		t.Error("Bad remaining() 3")
	}

	// Just to look.
	s.randomize()
	s.shuffle()
	fmt.Println(s.cards)
}
