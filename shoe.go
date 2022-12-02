package main

import (
	"container/list"
	"log"
	"math/rand"
	"time"
)

const bj_random_seed = 314159

// XXX: These should be in hand.go
// const ace = 11
// const softAce = 1

var suit = [13]int8{2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 11}
var deck [52]int8

// When the Shoe.shuffle method is called, we store a pointer to the shoe
// in 'curShoe', so our 'swap' function can find it. (Crude.)

var curShoe *Shoe

// This initializes the 'deck' variable with its initial contents.

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
		log.Fatal("bad deck creation")
	}
}

// Shoe represents a casino 'shoe' of some number of decks of cards.
// Cards are dealt from here as hands are played.

type Shoe struct {
	// numDecks    int
	cards     []int8
	shoeSize  int
	next      int
	remaining int
	queue     *list.List
}

// newShoe creates a new Shoe and fills it with the number of decks requested.
// Also, we seed the random number generator with a known value.

func newShoe(decks int) *Shoe {
	var s Shoe
	// s.numDecks = decks
	s.next = 0
	s.shoeSize = 52 * decks
	s.remaining = s.shoeSize
	for i := 0; i < decks; i++ {
		for _, j := range deck {
			s.cards = append(s.cards, j)
		}
	}
	rand.Seed(bj_random_seed)
	s.queue = list.New()
	return &s
}

// randomize seeds the random number generator with the current time.
// If you don't call this, you will get the same sequence of cards dealt
// in every execution of the main program.

func (s *Shoe) randomize() {
	rand.Seed(time.Now().UnixNano())
}

// shuffle shuffles the cards in the shoe.
func (s *Shoe) shuffle() {
	curShoe = s
	rand.Shuffle(s.shoeSize, swap)
	s.next = 0
	s.remaining = s.shoeSize
}

// deal returns the next card from the shoe. No check is made for an empty shoe.
// The caller is responsible for not allowing this to happen.

func (s *Shoe) deal() int8 {
	if s.queue.Len() == 0 {
		c := s.cards[s.next]
		s.next += 1
		s.remaining -= 1
		return c
	} else {
		temp := s.queue.Front()
		c := temp.Value.(int8)
		s.queue.Remove(temp)
		return c
	}
}

func (s *Shoe) force1(c int8) {
	s.queue.PushBack(c)
}

func (s *Shoe) force(cards []int8) {
	for _, c := range cards {
		s.queue.PushBack(c)
	}
}

// swap is called by rand.Shuffle() to swap two cards in the shoe.

func swap(i, j int) {
	s := *curShoe
	temp := s.cards[i]
	s.cards[i] = s.cards[j]
	s.cards[j] = temp
}
