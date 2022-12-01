package main

import (
	"flag"
	"fmt"
	"os"
)

// processCmdLine handles the program options and arguments.
// It returns true if all is well, and false if there were
// errors.

func processCmdLine(params *CmdLineParams) {
	var printVersion bool

	flag.IntVar(&params.numRounds, "n", 1, "number of rounds to deal")
	flag.IntVar(&params.numSeats, "s", 1, "number of table seats in use")
	flag.BoolVar(&params.repeatable, "r", false, "Same cards dealt every run (for testing)")
	flag.StringVar(&params.logFile, "l", "", "log file name")
	flag.BoolVar(&printVersion, "version", false, "print version and exit")
	flag.Parse()
	if printVersion {
		fmt.Printf("BJ version: %s\n", version)
		os.Exit(0)
	}

	if flag.NArg() != 2 {
		usage()
		os.Exit(2)
	} else {
		params.configFile = flag.Arg(0)
		params.strategyFile = flag.Arg(1)
	}
}

func usage() {
	fmt.Println("usage:")
	fmt.Println("  ./bj <options> <config-file> <strategy-file>")
	fmt.Println("Use -h to see all options.")
}
