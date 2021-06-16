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

// These values can appear in the 'key' position of a StrPoint.
const (
	spHitHard = iota + 50
	spHitSoft
	spSplit
	spDblHard
	spDblSoft
	spSurrender
)

// strategy is a logical 'set' of StrPoints to with the strategy says
// 'yes, do it'. This map is built by readStrategyFile, and is
// consulted during the play of a hand.
// var strategy map[StrPoint]bool

// Trace point names that will appear in the trace file.
// This is really a constant.
var traceName = [...]string{"ALWAYS", "INIT"}

// The open trace file, used by the trc package.
//var trf *os.File

func main() {
	var params CmdLineParams
	var cfg Config
	var strategy map[StrPoint]bool
	var trf *os.File

	err := processCmdLine(&params)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(params.traceFlags) > 0 {
		trf = openTraceFile()
		defer closeTraceFile(trf)
		trc.TraceOpen(trf)
		for n := range params.traceFlags {
			trc.TraceOn(n, traceName[n])
		}
		trc.Trace(trAlways, "trace open")
	}
	if trc.Tracing(trInit) {
		traceInitialParams(&params)
	}
	err = readConfigFile(params.configFile, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	strategy = make(Strategy)
	// fmt.Printf("penetration: %f\n", cfg.penetrationPct)
	err = readStrategyFile(params.strategyFile, strategy)
	if err != nil {
		fmt.Println(err)
		return
	}
	// XXX testing
	fmt.Println("strategy:")
	for k, v := range strategy {
		fmt.Printf("strat: %v %v\n", k, v)
	}
	// Initialization is complete. Now, play blackjack.
	// XXX think about penetrationPct
	game := newGame(strategy, params.numSeats, int(cfg.penetrationPct), params.repeatable, &cfg, params.verbose)
	fmt.Println(game)
}

func openTraceFile() *os.File {
	f, err := os.Create(traceFileName)
	if err != nil {
		log.Fatal("FATAL: ", err)
	}
	return f
}

func closeTraceFile(trf *os.File) {
	trc.Trace(trAlways, "trace close")
	trf.Close()
}

func traceInitialParams(params *CmdLineParams) {
	trc.Trace(trInit, "verbose: %v", params.verbose)
	trc.Trace(trInit, "repeatable: %v", params.repeatable)
	trc.Trace(trInit, "traceFlags: %v", params.traceFlags)
	trc.Trace(trInit, "numRounds: %d", params.numRounds)
	trc.Trace(trInit, "numSeats: %d", params.numSeats)
	trc.Trace(trInit, "configFile: %s", params.configFile)
	trc.Trace(trInit, "strategyFile: %s", params.strategyFile)
}
