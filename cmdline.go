package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// processCmdLine handles the program options and arguments.
// It returns true if all is well, and false if there were
// errors.

func processCmdLine(params *CmdLineParams) error {
	var printVersion bool
	var traceString string
	var traceList []string

	flag.BoolVar(&params.verbose, "v", false, "verbose output to stdout")
	flag.IntVar(&params.numRounds, "n", 1, "number of rounds to deal")
	flag.IntVar(&params.numSeats, "s", 1, "number of table seats in use")
	flag.StringVar(&traceString, "t", "", "list of trace flags to set")
	flag.BoolVar(&params.repeatable, "r", false, "Same cards dealt every run (for testing")
	flag.BoolVar(&printVersion, "version", false, "print version and exit")
	flag.Parse()
	if printVersion {
		fmt.Printf("BJ version: %s\n", version)
		os.Exit(0)
	}
	if params.verbose {
		fmt.Printf("verbose = %v\n", params.verbose)
		fmt.Printf("repeatable = %v\n", params.repeatable)
		fmt.Printf("trace flags = %s\n", traceString)
		fmt.Printf("num roungs = %d\n", params.numRounds)
		fmt.Printf("num seats = %d\n", params.numSeats)
	}
	if len(traceString) > 0 {
		traceList = strings.Split(traceString, ",")
		// fmt.Printf("trace strings: %v\n", traceList)
		for _, s := range traceList {
			n, err := strconv.Atoi(s)
			if err != nil {
				return fmt.Errorf("FATAL: bad trace flag: %v", traceList)

			}
			params.traceFlags = append(params.traceFlags, int8(n))
		}
		// fmt.Printf("trace flags: %v\n", traceFlags)
	}
	if flag.NArg() != 2 {
		usage()
		os.Exit(0)
	} else {
		params.configFile = flag.Arg(0)
		params.strategyFile = flag.Arg(1)
	}
	return nil
}

func usage() {
	fmt.Println("usage:")
	fmt.Println("  ./bj <options> <config-file> <strategy-file>")
	fmt.Println("Use -h to see all options.")
}
