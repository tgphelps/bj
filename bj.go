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

var traceName [2]string = [2]string{"ALWAYS", "INIT"}

var trf *os.File

func main() {
	if !processCmdLine() {
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
