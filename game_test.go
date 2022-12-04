package main

import (
	"io"
	"log"
	"testing"
)

func TestGame(t *testing.T) {
	var cfg Config
	//var strategy Strategy

	log.SetOutput(io.Discard)
	err := readConfigFile("data/house.cfg", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	_ = newGame(make(Strategy), 2, 25, true, &cfg)
	// fmt.Printf("game: %v\n", g)
}
