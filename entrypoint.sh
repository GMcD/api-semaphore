#!/bin/sh -l

echo "Hello $1"
time=$(date)
echo "time=$time" >> $GITHUB_OUTPUT

/usr/local/go/bin/go version

env | sort

psql -v

# /usr/local/go/bin/go test .
