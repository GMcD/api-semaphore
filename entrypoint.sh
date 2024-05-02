#!/bin/sh -l

echo "Hello $1"
time=$(date)
echo "time=$time" >> $GITHUB_OUTPUT

/usr/local/go/bin/go version

env | sort

psql --version

DB_HOST=$(hostname -I)

echo ${DB_HOST}

PGPASS=postgres
psql -Atx "host=postgres port=5432 dbname=postgres user=postgres" -c 'select current_date' 

APP_DB_HOST=postgres /usr/local/go/bin/go test .
