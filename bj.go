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

// Global constants.

const version = "0.0.1"
const traceFileName = "TRACE.txt"

const (
	trAlways = iota
	trInit
)

// Global variables that never changed after being set.

var verbose bool        // default: false
var repeatable bool     // default: false
var traceFlags []int8   // default: empty
var numRounds int       // default: 1
var numSeats int        // default: 1
var configFile string   // mandatory
var strategyFile string // mandatory

func main() {
	if !processCmdLine() {
		return
	}
}
