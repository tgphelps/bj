// Bj is a blackjack casino simulator,
// very much under construction

package main

// bj.py: Blackjack simulator, for studying the game.
//
// Usage:
//    bj [-d <flags>] [-v] [-t] [-n <rounds>] [-s <seats>] [--test] \
//       CONFIG STRATEGY
//
// Options:
//     -h  --help           Show this screen, and exit.
//     --version            Show version, and exit.
//     -v                   Be verbose.
//     -n <rounds>          Number of rounds to play.
//     -s <seats>           Number of players to play.
//     -r                   Use repeatable card sequence.

import (
	"io"
	"log"
	"os"
)

// Global constants.

const version = "0.0.1"

// const logFilename = "LOG.txt"
// const statsFileName = "STATS.txt"

// This should be even, so blackjack payoff is an integer.
const betAmount = 2

func init() {
	if betAmount%2 != 0 {
		log.Panic("betAmount must be even")
	}
}

func main() {
	var params CmdLineParams
	var cfg Config
	var strategy map[StrPoint]bool

	processCmdLine(&params)

	// If there are any command line problems, processCmdLine does not return.
	if params.logFile == "" {
		log.SetOutput(io.Discard)
	} else {
		f, err := os.Create(params.logFile)
		if err != nil {
			log.Panicf("FATAL: %s", err)
		}
		defer f.Close()
		log.SetOutput(f)
	}
	log.SetFlags(0)
	log.Println("init: log opened")
	log.Printf("init: repeatable = %v\n", params.repeatable)
	log.Printf("init: num rounds = %d\n", params.numRounds)
	log.Printf("init: num seats = %d\n", params.numSeats)

	err := readConfigFile(params.configFile, &cfg)
	if err != nil {
		log.Panic(err)
	}
	strategy = make(Strategy)
	// fmt.Printf("penetration: %f\n", cfg.penetrationPct)
	err = readStrategyFile(params.strategyFile, strategy)
	if err != nil {
		log.Panic(err)
	}

	//game := newGame(strategy, params.numSeats, cfg.penetrationPct, params.repeatable, &cfg)
	//for i := 0; i < params.numRounds; i++ {
	//log.Printf("ROUND %d\n", i+1)
	//game.playRound()
	//}
	//game.writeStats(statsFileName, params.strategyFile)
	log.Println("term: log closing")
	// XXX Check for needing to close the log file.
}
