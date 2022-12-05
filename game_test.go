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

// XXX: Obsolete test not working yet.
func TestGameObsolete(t *testing.T) {
	var cfg Config
	var strategy Strategy

	//log.SetOutput(io.Discard)
	log.SetFlags(0)
	readTestFiles(&cfg, strategy)
	g := newGame(strategy, 1, true, &cfg)
	g.playRound()
	g.writeStats(cfg.statsFilename, testStrategyFile)
}

func TestGameOnlyDealerBJ(t *testing.T) {
	var cfg Config
	var strategy Strategy

	//log.SetOutput(io.Discard)
	log.SetFlags(0)
	readTestFiles(&cfg, strategy)
	g := newGame(strategy, 1, true, &cfg)
	g.shoe.force([]int8{10, 10, 10, 11})
	g.playRound()
	g.writeStats(cfg.statsFilename, testStrategyFile)
	if g.st.totalBet != 2 || g.st.totalLost != 2 || g.st.blackjacksWon != 0 {
		t.Error("Dealer only BJ stats failed")
	}
}

func TestGameOnlyPlayerBJ(t *testing.T) {
	var cfg Config
	var strategy Strategy

	//log.SetOutput(io.Discard)
	log.SetFlags(0)
	readTestFiles(&cfg, strategy)
	g := newGame(strategy, 1, true, &cfg)
	g.shoe.force([]int8{10, 11, 10, 10})
	g.playRound()
	g.writeStats(cfg.statsFilename, testStrategyFile)
	if g.st.totalBet != 2 || g.st.totalWon != 3 || g.st.blackjacksWon != 1 {
		t.Error("Player only BJ stats failed")
	}
}

func TestGameBothHaveBJ(t *testing.T) {
	var cfg Config
	var strategy Strategy

	//log.SetOutput(io.Discard)
	log.SetFlags(0)
	readTestFiles(&cfg, strategy)
	g := newGame(strategy, 1, true, &cfg)
	g.shoe.force([]int8{10, 11, 10, 11})
	g.playRound()
	g.writeStats(cfg.statsFilename, testStrategyFile)
	if g.st.totalBet != 2 || g.st.totalWon != 0 || g.st.totalPush != 2 {
		t.Error("Player only BJ stats failed")
	}
}
func TestGamePlayerBust(t *testing.T) {
	var cfg Config
	var strategy Strategy

	//log.SetOutput(io.Discard)
	log.SetFlags(0)
	readTestFiles(&cfg, strategy)
	g := newGame(strategy, 1, true, &cfg)
	g.shoe.force([]int8{10, 6, 10, 10, 6})
	g.playRound()
	g.writeStats(cfg.statsFilename, testStrategyFile)
	if g.st.totalBet != 2 || g.st.totalLost != 2 {
		t.Error("Player bust stats failed")
	}
}

func TestGameDealerBust(t *testing.T) {
	var cfg Config
	var strategy Strategy

	//log.SetOutput(io.Discard)
	log.SetFlags(0)
	readTestFiles(&cfg, strategy)
	g := newGame(strategy, 1, true, &cfg)
	g.shoe.force([]int8{10, 10, 10, 6, 6})
	g.playRound()
	g.writeStats(cfg.statsFilename, testStrategyFile)
	if g.st.totalBet != 2 || g.st.totalWon != 2 {
		t.Error("Dealer bust stats failed")
	}
}

func TestGameDealerWin(t *testing.T) {
	var cfg Config
	var strategy Strategy

	//log.SetOutput(io.Discard)
	log.SetFlags(0)
	readTestFiles(&cfg, strategy)
	g := newGame(strategy, 1, true, &cfg)
	g.shoe.force([]int8{10, 8, 10, 10})
	g.playRound()
	g.writeStats(cfg.statsFilename, testStrategyFile)
	if g.st.totalBet != 2 || g.st.totalLost != 2 {
		t.Error("Dealer win stats failed")
	}
}

func TestGamePlayerWin(t *testing.T) {
	var cfg Config
	var strategy Strategy

	//log.SetOutput(io.Discard)
	log.SetFlags(0)
	readTestFiles(&cfg, strategy)
	g := newGame(strategy, 1, true, &cfg)
	g.shoe.force([]int8{10, 10, 10, 6, 3})
	g.playRound()
	g.writeStats(cfg.statsFilename, testStrategyFile)
	if g.st.totalBet != 2 || g.st.totalWon != 2 {
		t.Error("Player win stats failed")
	}
}

func TestGamePush(t *testing.T) {
	var cfg Config
	var strategy Strategy

	//log.SetOutput(io.Discard)
	log.SetFlags(0)
	readTestFiles(&cfg, strategy)
	g := newGame(strategy, 1, true, &cfg)
	g.shoe.force([]int8{10, 10, 10, 10})
	g.playRound()
	g.writeStats(cfg.statsFilename, testStrategyFile)
	if g.st.totalBet != 2 || g.st.totalPush != 2 {
		t.Error("Push stats failed")
	}
}
