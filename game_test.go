package main

import (
	"fmt"
	"log"
	"testing"
)

func TestGame(t *testing.T) {
	var cfg Config
	//var strategy Strategy
	err := readConfigFile("data/house.cfg", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	g := newGame(make(Strategy), 2, 25, true, &cfg, true)
	fmt.Printf("game: %v\n", g)
}
