#!/bin/bash

DIRS="tree skiplist multiset set"
OPS="Lookup Add Remove Clone"

for DIR in $DIRS
do
    cd $DIR
    for OP in $OPS
    do
	go test -run=X -cpu 1 -bench=$OP
    done
    cd ..
done
