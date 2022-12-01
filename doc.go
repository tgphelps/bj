package main

// Package bj: Blackjack simulator
//
// This program tries to simulate a real casino blackjack game.
// Rule variations are handled via a 'house rules' file, which is read at
// startup. The dealer strategy is defined completely by the house rules.
// The player strategy is specified via a strategy file which is also read
// at startup.
// The program simulates the dealing and playing of hands, and keeps track of
// the results, money won and lost, and overall player (dis)advantage.

// To run the program, do this:
//   bj [-n <int>] [-s <int>] [-r] [l <string>] <rules-file> <strategy-file>
// where:
//   -n <int> = number of rounds to deal
//   -s < int> = number of seats at the table to play
//   -r means 'deal repeatable hands' so strategy variations can be compared.
//   -l <string> = log file name. Default is no log at all.

// House rules
//
// There are EIGHT house rules, and ALL of them must be set in the house rules
// file specified on the command line. Comments (lines starting with '#') and
// blank lines are ignored. There should be 8 rules lines of this format:
//
// param1 = value1
//
// A value must be an integer. If the rule is logically a boolean value, then
// use 1 and 0 to represent true and false. The parameters must be spelled
// exactly as documented below.
//
// The house rules are:
// hitS17 - 1 if the dealer hits soft 17. Typically true.
// dasAllowed - 1 if double after split is allowed. Typically true.
// maxSplitHands - max number of split hands a player may have. Typically 4.
// maxSplitAces - max number of split ace hands a player may have. Typically 2.
//     (To be clear '2' means a pair of aces can be split only ONCE.)
// canHitSplitAces - 1 of a player may hit after having split a pair of aces.
//     Typically false.
// canSurrender - 1 if (late) surrender is allowed.
// penetrationPct - Percentage of a full shoe remaining when a dealer should
//     shuffle. Example: 25 means shuffle when there is 25% of the shoe
//     remaining.

// Player strategy

// There are 6 player strategy decisions to make:
// 1. Hit a hard hand, or not.
// 2. Hit a soft hand, or not.
// 3. Double a hard hand.
// 4. Double a soft hand.
// 5. Split a pair.
// 6. Surrender.

// Here are sample strategy file lines for each decision. In the following,
// <nums1> is a list of player hand values, <nums2> is a list of
// dealer up-cards, and <nums2> is a list of pair cards:
//
// hit hard <nums1> vs <nums2>
// hit soft <nums1> vs <nums2>
// double hard <nums1> vs <nums2>
// double soft <nums1> vs <nums2>
// split <nums3> vs <nums2>
//   example: split 2,3 vs 2,3,4,5,6,7 # split 2's and 3's versus 2 through 7
// surrender <nums1> vs <nums2>

// The strategy file is used to build a strategy map which tells what to do
// in every situation. The map has keys which are tuples of
// (decision, player-hand, dealer-upcard)
// and values which are booleans. The decision field is one of the 6 decision
// types from the strategy file. If a key is not in the map, it is assumed to
// be false.
