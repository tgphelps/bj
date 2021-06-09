package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func readStrategyFile(strategyile string) error {
	f, err := os.Open(strategyFile)
	if err != nil {
		return fmt.Errorf("FATAL: %s", err)
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
				return fmt.Errorf("error on strategy file: %s", err)
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
	return nil
}
