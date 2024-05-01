#!/bin/sh -l

echo "Hello $1"
time=$(date)
echo "time=$time" >> $GITHUB_OUTPUT

/opt/hostedtoolcache/go/1.20.14/x64/bin/go version
