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

func readStrategyFile(strategyFile string, strategy Strategy) error {
	//var s StrPoint
	f, err := os.Open(strategyFile)
	if err != nil {
		return fmt.Errorf("FATAL: %s", err)
	}
	defer f.Close()
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
			var a []string
			// handle the strategy line here
			if !strings.HasPrefix(s, "#") {
				// This is NOT a comment
				a = strings.Fields(s)
				if len(a) == 0 {
					continue
				}
				// The line is not blank. It should be a real strategy line.
				switch a[0] {
				case "hit":
					switch a[1] {
					case "hard":
						do_strategy_hit(strategy, keyHitHard, a)
					case "soft":
						do_strategy_hit(strategy, keyHitSoft, a)
					default:
						log.Panic("invalid strategy line")
					}
				default:
					log.Panic("illegal strategy line")
				}
			}
		}
	}
	return nil
}

func NewStrPoint(key int8, val int8, upcard int8) StrPoint {
	var s StrPoint
	s[0] = key
	s[1] = val
	s[2] = upcard
	return s
}

func do_strategy_hit(s Strategy, key int8, a []string) {
	vals := strings.Split(a[2], ",")
	upcards := strings.Split(a[4], ",")
	for _, val := range vals {
		for _, upcard := range upcards {
			//fmt.Printf("val: %s upcard: %s\n", val, upcard)
			intVal, _ := strconv.Atoi(val)
			intUp, _ := strconv.Atoi(upcard)
			strp := NewStrPoint(key, int8(intVal), int8(intUp))
			//fmt.Println(strp)
			s[strp] = true
		}
	}
}
