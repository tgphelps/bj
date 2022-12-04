package main

// "fmt"

// Dealer represents a casino dealer, who plays his hand according to the
// rules of the casino.

type Dealer struct {
	shoe    *Shoe
	hit_s17 bool
	hand    *Hand
}

// newDealer creates and initializes a new Dealer.

func newDealer(shoe *Shoe, hit_s17 bool) *Dealer {
	var d Dealer
	d.shoe = shoe
	d.hit_s17 = hit_s17
	return &d
}

// getHand gets a new hand for the dealer.`

func (d *Dealer) getHand() {
	d.hand = newHand(d.shoe, 0)
	if d.hand.value == 21 {
		d.hand.isBlackjack = true
	}
}

// upCard returns the value of the dealer's up-card.

func (d *Dealer) upCard() int8 {
	card := d.hand.cards[0]
	// If the hand contains 2 aces, one of them has been converted to a 1.
	// Make sure to return 11 and not 1 to the caller.
	if card == softAce {
		card = ace
	}
	return card
}

// playHand plays the dealer's hand. The only variale is whether we hit or
// stand on soft 17.

func (d *Dealer) playHand() {
	for (d.hand.value < 17) || (d.hand.value == 17 && d.hand.isSoft() && d.hit_s17) {
		d.hand.hit()
		// log.Printf("dealer hit. Hand: %s\n", d.hand)
	}
}
