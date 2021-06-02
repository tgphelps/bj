package main

import (
	// "math/rand"
	"fmt"
)

var suit = [13]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 11}
var deck [52]int

func init() {
	n := 0
	for i := 0; i < 4; i++ {
		for _, card := range suit {
			deck[n] = card
			n += 1
		}
	}
	fmt.Println(deck)
	if len(deck) != 52 {
		panic("bad deck creation")
	}
}

type Shoe struct {
	decks    int
	shoe     []int
	shoeSize int
	next     int
	// thisRound []int
	// trackRounds bool
}

func newShoe(decks int, repeatable bool) *Shoe {
	var s Shoe
	s.decks = decks
	s.next = 0
	s.shoeSize = 52 * decks
	for i := 0; i < decks; i++ {
		for _, j := range deck {
			s.shoe = append(s.shoe, j)
		}
	}
	return &s
}
