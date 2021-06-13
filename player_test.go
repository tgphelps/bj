package main

import (
	"fmt"
	"testing"
)

func TestPlayer(t *testing.T) {
	fmt.Println("test player...")
	var cfg Config
	var strategy Strategy
	shoe := newShoe(6)
	p := newPlayer(1, shoe, &cfg, strategy, 2, false)
	fmt.Printf("got new player: %v\n", p)
}
