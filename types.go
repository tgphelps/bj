package main

import (
	"strconv"
	"strings"
)

type CmdLineParams struct {
	repeatable   bool
	numRounds    int
	numSeats     int
	configFile   string
	strategyFile string
	logFile      string
}

type Config struct {
	numDecks        int
	hitS17          bool
	dasAllowed      bool
	maxSplitHands   int
	maxSplitAces    int
	canHitSplitAces bool
	canSurrender    bool
	penetrationPct  int
	statsFilename   string
}

// XXX: This constant MUST be kept in sync with the Config struct, and
// XXX: with the code in read_config_file.go.
const numConfigParams = 9

// A StrPoint (strategy point) represents a decision to be made. It consists of 3 numbers:
// 1. key - what we might do (hit, double, split, etc.)
// 2. val - the value (count) of the hand
// 3. upcard - the dealer's upcard.

type StrPoint [3]int8

// These go in the 'key' field of a StrPoint
const (
	keyHitHard = iota + 50
	keyHitSoft
	keySplit
	keyDblHard
	keyDblSoft
	keySurrender
)

var keymap = map[int8]string{
	keyHitHard:   "HH",
	keyHitSoft:   "HS",
	keySplit:     "SP",
	keyDblHard:   "DH",
	keyDblSoft:   "DS",
	keySurrender: "SU",
}

func (s *StrPoint) String() string {
	var sb strings.Builder
	sb.WriteString("str{")
	sb.WriteString(keymap[s[0]])
	sb.WriteString(", ")
	sb.WriteString(strconv.Itoa(int(s[1])))
	sb.WriteString(", ")
	sb.WriteString(strconv.Itoa(int(s[2])))
	sb.WriteString("}")
	return sb.String()
}

// One Strategy drives all playing decisions.  It is built by reading the
// strategy file, and defines the circumstances under which we surrender,
// split, double, and hit.

type Strategy map[StrPoint]bool

// This contains all information collected during a Game.

type Statistics struct {
	roundsPlayed  int
	handsPlayed   int
	blackjacksWon int
	totalBet      int
	totalWon      int
	totalLost     int
	totalPush     int
}
