package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"tgphelps.com/trc"
)

func readStrategyFile(strategyFile string, strategy Strategy) error {
	var s StrPoint
	f, err := os.Open(strategyFile)
	if err != nil {
		return fmt.Errorf("FATAL: %s", err)
	}
	defer f.Close()
	trc.TraceIf(trInit, "reading strategy file")
	// strategy = make(map[StrPoint]bool)
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
					// XXX
					panic("Strategy file not implemented")
				}
			}
		}
	}
	// XXX testing
	s[0] = 1
	s[1] = 2
	s[2] = 3
	strategy[s] = true
	trc.TraceIf(trInit, "end strategy file")
	// XXX

	return nil
}

func NewStrPoint(key int8, val int8, upcard int8) StrPoint {
	var s StrPoint
	s[0] = key
	s[1] = val
	s[2] = upcard
	return s
}

// func inStrategyMap(key int8, val int8, upcard int8) bool {
// s := NewStrPoint(key, val, upcard)
// return strategy[s]
// }
