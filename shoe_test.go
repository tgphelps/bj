package main

import (
	"fmt"
	"testing"
)

func TestShoe1(t *testing.T) {
	s := newShoe(2)
	if s.shoeSize != 104 {
		t.Error("Bad deck count")
	}
	fmt.Println(s.cards)
}
