package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// readConfigFile reads the house configuration from a text file.
// It returns true if it was successfully read, else false.

func readConfigFile(cfgFile string) bool {
	f, err := os.Open(cfgFile)
	if err != nil {
		fmt.Printf("FATAL: %s\n", err)
		return false
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		var s string
		s, err := r.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				fmt.Printf("error on config file: %s\n", err)
				return false
			}
		} else {
			if !strings.HasPrefix(s, "#") {
				a := strings.Fields(s)
				if len(a) == 3 {
					fmt.Printf("%v\n", a)
				}
			}
		}
	}
	return true
}

func readStrategyFile(strategyile string) bool {
	f, err := os.Open(strategyFile)
	if err != nil {
		fmt.Printf("FATAL: %s\n", err)
		return false
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		var s string
		s, err := r.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				fmt.Printf("error on strategy file: %s\n", err)
				return false
			}
		} else {
			// handle the strategy line here
			if !strings.HasPrefix(s, "#") {
				a := strings.Fields(s)
				if len(a) > 0 {
					fmt.Printf("%v\n", a)
				}
			}
		}
	}
	return true
}
