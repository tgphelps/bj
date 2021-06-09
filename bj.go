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
//     -t <flags>           Set trace flags.
//     -n <rounds>          Number of rounds to play.
//     -s <seats>           Number of players to play.
//     -r                   Use repeatable card sequence.

import (
	// "fmt"
	"fmt"
	"log"
	"os"

	"tgphelps.com/trc"
)

// Global constants.

const version = "0.0.1"
const traceFileName = "TRACE.txt"

const (
	trAlways = iota
	trInit
)

// Global variables that never change after being set.

var verbose bool        // default: false
var repeatable bool     // default: false
var traceFlags []int8   // default: empty
var numRounds int       // default: 1
var numSeats int        // default: 1
var configFile string   // mandatory
var strategyFile string // mandatory

type Config struct {
	numDecks        int
	hitS17          bool
	dasAllowed      bool
	maxSplitHands   int
	maxSplitAces    int
	canHitSplitAces bool
	canSurrender    bool
}

var cfg Config

var traceName = [...]string{"ALWAYS", "INIT"}

var trf *os.File

func main() {
	err := processCmdLine()
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(traceFlags) > 0 {
		openTraceFile()
		defer closeTraceFile()
		trc.TraceOpen(trf)
		for n := range traceFlags {
			trc.TraceOn(n, traceName[n])
		}
		trc.Trace(trAlways, "trace open")
	}
	if trc.Tracing(trInit) {
		traceInitialParams()
	}
	err = readConfigFile(configFile)
	if err != nil {
		return
	}
	err = readStrategyFile(strategyFile)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func openTraceFile() {
	f, err := os.Create(traceFileName)
	if err != nil {
		log.Fatal("FATAL: ", err)
	}
	trf = f
}

func closeTraceFile() {
	trc.Trace(trAlways, "trace close")
	trf.Close()
}

func traceInitialParams() {
	trc.Trace(trInit, "verbose: %v", verbose)
	trc.Trace(trInit, "repeatable: %v", repeatable)
	trc.Trace(trInit, "traceFlags: %v", traceFlags)
	trc.Trace(trInit, "numRounds: %d", numRounds)
	trc.Trace(trInit, "numSeats: %d", numSeats)
	trc.Trace(trInit, "configFile: %s", configFile)
	trc.Trace(trInit, "strategyFile: %s", strategyFile)
}
