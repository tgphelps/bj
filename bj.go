// Bj is a blackjack casino simulator

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
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Global constants.

const version = "0.0.1"

// Global variables. Never changed after being set.

var verbose bool
var repeatable bool
var traceFlags []int8
var numRounds int
var numSeats int
var configFile string
var strategyFile string

func main() {
	if !handleArguments() {
		return
	}
}

func handleArguments() bool {
	var printVersion bool
	var traceString string
	var traceList []string

	flag.BoolVar(&verbose, "v", false, "verbose output to stdout")
	flag.IntVar(&numRounds, "n", 1, "number of rounds to deal")
	flag.IntVar(&numSeats, "s", 1, "number of table seats in use")
	flag.StringVar(&traceString, "t", "", "trace flags to set")
	flag.BoolVar(&repeatable, "r", false, "Same cards dealt every run (for testing")
	flag.BoolVar(&printVersion, "version", false, "print version and exit")
	flag.Parse()
	if printVersion {
		fmt.Printf("BJ version: %s\n", version)
		return false
	}
	fmt.Printf("verbose = %v\n", verbose)
	fmt.Printf("repeatable = %v\n", repeatable)
	fmt.Printf("trace flags = %s\n", traceString)
	fmt.Printf("num roungs = %d\n", numRounds)
	fmt.Printf("num seats = %d\n", numSeats)
	if len(traceString) > 0 {
		traceList = strings.Split(traceString, ",")
		fmt.Printf("trace strings: %v\n", traceList)
		for _, s := range traceList {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("FATAL: bad trace flag: %v", traceList)
			}
			traceFlags = append(traceFlags, int8(n))
		}
		fmt.Printf("trace flags: %v\n", traceFlags)
	}
	if flag.NArg() != 2 {
		usage()
		return false
	} else {
		configFile = flag.Arg(0)
		strategyFile = flag.Arg(1)
	}
	return true
}

func usage() {
	fmt.Println("usage:")
	fmt.Println("  ./bj <options> <config-file> <strategy-file>")
	fmt.Println("Enter './bj -h' to see all options.")
}
