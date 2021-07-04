package main

import "fmt"

// "fmt"

// Dealer represents a casino dealer, who plays his hand according to the
// rules of the casino.

type Dealer struct {
	shoe    *Shoe
	hit_s17 bool
	hand    *Hand
	value   int8
	verbose bool
}

// newShoe creates a new Shoe and fills it with the number of decks requested.
// Also, we seed the random number generator with a known value.

func newDealer(shoe *Shoe, hit_s17 bool, verbose bool) *Dealer {
	var d Dealer
	d.shoe = shoe
	d.hit_s17 = hit_s17
	d.verbose = verbose
	return &d
}

// getHand gets a new hand for the dealer.`

func (d *Dealer) getHand() {
	d.hand = newHand(d.shoe, 0)
}

// upCard returns the value of the dealer's up-card.

func (d *Dealer) upCard() int8 {
	card := d.hand.cards[0]
	if card == softAce {
		card = ace
	}
	return card
}

func (d *Dealer) playHand() {
	for (d.hand.value < 17) || (d.hand.value == 17 && d.hand.isSoft() && d.hit_s17) {
		d.hand.hit()
		if d.verbose {
			fmt.Printf("dealer hit. Hand: %s\n", d.hand)
		}
	}
}
