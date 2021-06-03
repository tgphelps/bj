package main

import (
	"fmt"
	"testing"
)

func TestShoe1(t *testing.T) {
	var c int8
	s := newShoe(2)
	if s.shoeSize != 104 {
		t.Error("Bad deck count")
	}
	s.shuffle()
	fmt.Println(s.cards)
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
	for i := 0; i < 100; i++ {
		c = s.deal()
	}
	if s.remaining() != 0 {
		t.Error("Bad remaining() 3")
	}

	// s.randomize()
	// s.shuffle()
	// fmt.Println(s.cards)
}
