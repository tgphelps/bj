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

// Trace points for calls to the trc package.
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

// Config contains all the house rules
var cfg Config

// A StrPoint represents a decision to be made. It consists of 3 numbers:
// 1. key - what we might do (hit, double, split, etc.)
// 2. val - the value (count) of the hand
// 3. upcard - the dealer's upcard.
type StrPoint [3]int8

// strategy is a logical 'set' of StrPoints to with the strategy says
// 'yes, do it'. This map is built by readStrategyFile, and is
// consulted during the play of a hand.
var strategy map[StrPoint]bool

// Trace point names that will appear in the trace file.
var traceName = [...]string{"ALWAYS", "INIT"}

// The open trace file, used by the trc package.
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
