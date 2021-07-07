#!/bin/sh

rm STATS.txt
./bj  -n 2000000 -s 5 data/house.cfg data/never-hit.txt
../bjstats/bjstats STATS.txt
