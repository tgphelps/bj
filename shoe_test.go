package main

import (
	"fmt"
	"testing"
)

func TestShoe1(t *testing.T) {
	s := newShoe(2, false)
	if s.decks != 2 {
		t.Error("Bad deck count")
	}
	fmt.Println(s.shoe)
}
