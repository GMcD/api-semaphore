#!/bin/sh -l

echo "Hello $1"
time=$(date)
echo "time=$time" >> $GITHUB_OUTPUT

/usr/local/go/bin/go version

env | sort

psql --version

APP_DB_HOST=$(hostname -I)

echo ${APP_DB_HOST}

/usr/local/go/bin/go test .
