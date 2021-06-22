package main

type CmdLineParams struct {
	verbose      bool
	repeatable   bool
	traceFlags   []int8
	numRounds    int
	numSeats     int
	configFile   string
	strategyFile string
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
}

// A StrPoint represents a decision to be made. It consists of 3 numbers:
// 1. key - what we might do (hit, double, split, etc.)
// 2. val - the value (count) of the hand
// 3. upcard - the dealer's upcard.
type StrPoint [3]int8

const (
	keyHitHard = iota + 50
	keyHitSoft
	keySplit
	keyDblHard
	keyDblSoft
	keySurrender
)

type Strategy map[StrPoint]bool

type Session struct {
	numPlayers int
	verbose    bool
}

type Statistics struct {
	roundsPlayed  int
	handsPlayed   int
	blackjacksWon int
	totalBet      int
	totalWon      int
	totalLost     int
	totalPush     int
}
