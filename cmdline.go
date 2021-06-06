package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// processCmdLine handles the program options and arguments.
// It returns true if all is well, and false if there were
// errors.

func processCmdLine() bool {
	var printVersion bool
	var traceString string
	var traceList []string

	flag.BoolVar(&verbose, "v", false, "verbose output to stdout")
	flag.IntVar(&numRounds, "n", 1, "number of rounds to deal")
	flag.IntVar(&numSeats, "s", 1, "number of table seats in use")
	flag.StringVar(&traceString, "t", "", "list of trace flags to set")
	flag.BoolVar(&repeatable, "r", false, "Same cards dealt every run (for testing")
	flag.BoolVar(&printVersion, "version", false, "print version and exit")
	flag.Parse()
	if printVersion {
		fmt.Printf("BJ version: %s\n", version)
		return false
	}
	if verbose {
		fmt.Printf("verbose = %v\n", verbose)
		fmt.Printf("repeatable = %v\n", repeatable)
		fmt.Printf("trace flags = %s\n", traceString)
		fmt.Printf("num roungs = %d\n", numRounds)
		fmt.Printf("num seats = %d\n", numSeats)
	}
	if len(traceString) > 0 {
		traceList = strings.Split(traceString, ",")
		fmt.Printf("trace strings: %v\n", traceList)
		for _, s := range traceList {
			n, err := strconv.Atoi(s)
			if err != nil {
				fmt.Printf("FATAL: bad trace flag: %v\n", traceList)
				return false
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
