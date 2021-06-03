package main

import (
	"math/rand"
	// "fmt"
)

const bj_random_seed = 314159

var suit = [13]int8{2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 11}
var deck [52]int8

func init() {
	n := 0
	for i := 0; i < 4; i++ {
		for _, card := range suit {
			deck[n] = card
			n += 1
		}
	}
	// fmt.Println(deck)
	if len(deck) != 52 {
		panic("bad deck creation")
	}
}

type Shoe struct {
	// numDecks    int
	cards      []int8
	shoeSize   int
	next       int
	repeatable bool
	// thisRound []int
	// trackRounds bool
}

func newShoe(decks int) *Shoe {
	var s Shoe
	// s.numDecks = decks
	s.next = 0
	s.repeatable = false
	s.shoeSize = 52 * decks
	for i := 0; i < decks; i++ {
		for _, j := range deck {
			s.cards = append(s.cards, j)
		}
	}
	return &s
}

func (s *Shoe) setRepeatable() {
	s.repeatable = true
	rand.Seed(bj_random_seed)
}
