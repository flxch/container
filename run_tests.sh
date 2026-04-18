#!/bin/bash

DIRS="option result pair tree skiplist multiset set"

for DIR in $DIRS
do
    cd $DIR
    go test $@
    cd ..
done
