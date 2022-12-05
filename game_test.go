package main

import (
	"log"
	"testing"
)

// common initialization

const testStrategyFile = "data/00-never-hit.txt"

func readTestFiles(cfg *Config, strategy Strategy) {
	err := readConfigFile("data/house-test.cfg", cfg)
	if err != nil {
		log.Panic(err)
	}
	err = readStrategyFile(testStrategyFile, strategy)
	if err != nil {
		log.Panic(err)
	}
}

func TestNewGame(t *testing.T) {
	var cfg Config
	var strategy Strategy

	readTestFiles(&cfg, strategy)

	_ = newGame(strategy, 2, true, &cfg)
}

// Cards dealt from one shuffled deck are:
// 10 7 6 3 9 8 8 10 10 6 4 3 5 9 10 3 9 2 4 7
func TestGameOneRound(t *testing.T) {
	var cfg Config
	var strategy Strategy

	//log.SetOutput(io.Discard)
	log.SetFlags(0)
	readTestFiles(&cfg, strategy)
	g := newGame(strategy, 1, true, &cfg)
	// log.Print("start test round")s
	g.playRound()
	g.writeStats(cfg.statsFilename, testStrategyFile)
}

func TestGameObsolete(t *testing.T) {
	var cfg Config
	var strategy Strategy

	//log.SetOutput(io.Discard)
	log.SetFlags(0)
	readTestFiles(&cfg, strategy)
	g := newGame(strategy, 1, true, &cfg)
	// log.Print("start test round")s
	g.playRound()
	g.writeStats(cfg.statsFilename, testStrategyFile)
}

func TestGameDealerBJ(t *testing.T) {
	var cfg Config
	var strategy Strategy

	//log.SetOutput(io.Discard)
	log.SetFlags(0)
	readTestFiles(&cfg, strategy)
	g := newGame(strategy, 1, true, &cfg)
	// log.Print("start test round")s
	g.playRound()
	g.writeStats(cfg.statsFilename, testStrategyFile)
}

func TestGamePlayerBJ(t *testing.T) {
	var cfg Config
	var strategy Strategy

	//log.SetOutput(io.Discard)
	log.SetFlags(0)
	readTestFiles(&cfg, strategy)
	g := newGame(strategy, 1, true, &cfg)
	// log.Print("start test round")s
	g.playRound()
	g.writeStats(cfg.statsFilename, testStrategyFile)
}

func TestGamePlayerBust(t *testing.T) {
	var cfg Config
	var strategy Strategy

	//log.SetOutput(io.Discard)
	log.SetFlags(0)
	readTestFiles(&cfg, strategy)
	g := newGame(strategy, 1, true, &cfg)
	// log.Print("start test round")s
	g.playRound()
	g.writeStats(cfg.statsFilename, testStrategyFile)
}

func TestGameDealerBust(t *testing.T) {
	var cfg Config
	var strategy Strategy

	//log.SetOutput(io.Discard)
	log.SetFlags(0)
	readTestFiles(&cfg, strategy)
	g := newGame(strategy, 1, true, &cfg)
	// log.Print("start test round")s
	g.playRound()
	g.writeStats(cfg.statsFilename, testStrategyFile)
}

func TestGameDealerWin(t *testing.T) {
	var cfg Config
	var strategy Strategy

	//log.SetOutput(io.Discard)
	log.SetFlags(0)
	readTestFiles(&cfg, strategy)
	g := newGame(strategy, 1, true, &cfg)
	// log.Print("start test round")s
	g.playRound()
	g.writeStats(cfg.statsFilename, testStrategyFile)
}

func TestGamePlayerWin(t *testing.T) {
	var cfg Config
	var strategy Strategy

	//log.SetOutput(io.Discard)
	log.SetFlags(0)
	readTestFiles(&cfg, strategy)
	g := newGame(strategy, 1, true, &cfg)
	// log.Print("start test round")s
	g.playRound()
	g.writeStats(cfg.statsFilename, testStrategyFile)
}

func TestGamePush(t *testing.T) {
	var cfg Config
	var strategy Strategy

	//log.SetOutput(io.Discard)
	log.SetFlags(0)
	readTestFiles(&cfg, strategy)
	g := newGame(strategy, 1, true, &cfg)
	// log.Print("start test round")s
	g.playRound()
	g.writeStats(cfg.statsFilename, testStrategyFile)
}
