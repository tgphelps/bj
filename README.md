
Package bj: Blackjack simulator

This program tries to simulate a real casino blackjack game.
Rule variations are handled via a 'house rules' file, which is read at
startup. The dealer strategy is defined completely by the house rules.
The player strategy is specified via a strategy file which is also read
at startup.
The program simulates the dealing and playing of hands, and keeps track of
the results, money won and lost, and overall player (dis)advantage.

To run the program, do this:
  bj [-n <int>] [-s <int>] [-r] [l <string>] <rules-file> <strategy-file>
where:
  -n <int> = number of rounds to deal
  -s < int> = number of seats at the table to play
  -r means 'deal repeatable hands' so strategy variations can be compared.
  -l <string> = log file name. Default is no log at all.
