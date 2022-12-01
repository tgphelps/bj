package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// readConfigFile reads the house configuration from a text file.
// It returns true if it was successfully read, else false.

func readConfigFile(cfgFile string, cfg *Config) error {
	f, err := os.Open(cfgFile)
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
				return fmt.Errorf("error on config file: %s", err)
			}
		} else {
			if !strings.HasPrefix(s, "#") {
				a := strings.Fields(s)
				if len(a) == 0 {
					continue
				}
				if len(a) == 3 && a[1] == "=" {
					setConfigVar(a[0], a[2], cfg)
				} else {
					// trc.TraceIf(trAlways, "bad config: %v", a)
					return fmt.Errorf("FATAL: bad config: %v", a)
				}
			}
		}
	}
	return nil
}

func setConfigVar(tag string, val string, cfg *Config) {
	n, err := strconv.Atoi(val)
	if err != nil {
		log.Panic("BAD config line")
	} else {
		switch tag {
		case "numDecks":
			cfg.numDecks = n
		case "hitS17":
			cfg.hitS17 = toBool(n)
		case "dasAllowed":
			cfg.dasAllowed = toBool(n)
		case "maxSplitHands":
			cfg.maxSplitHands = n
		case "maxSplitAces":
			cfg.maxSplitAces = n
		case "canHitSplitAces":
			cfg.canHitSplitAces = toBool(n)
		case "canSurrender":
			cfg.canSurrender = toBool(n)
		case "penetrationPct":
			cfg.penetrationPct = n
		default:
			panic("BAD onfig line")
		}
	}
}

func toBool(n int) bool {
	if n == 0 {
		return false
	} else {
		return true
	}
}
