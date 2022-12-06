#!/bin/sh

# ./bj -n 200000 -s 5 -l LOG.txt data/house.cfg data/00-never-hit.txt
# 10 million hands
./bj -n 2000000 -s 5 data/house.cfg data/00-never-hit.txt
