#!/bin/sh -l

# Load Env from .env
source .env

echo "Hello $1"
time=$(date)
echo "time=$time" >> $GITHUB_OUTPUT

echo Mode: $2

if [ "$2" = "test" ]; then
    echo "Testing Mode..."
    env | sort

    psql --version

    PGPASS=${APP_DB_PASSWORD}
    psql -Atx "host=${APP_DB_HOST} dbname=${APP_DB_NAME} user=${APP_DB_USERNAME}" -c 'select current_date' 

    /usr/local/go/bin/go test .

    exit $?
fi

# if [ "$2" = "run" ]; then
    echo "Running Mode..."
    /usr/local/go/bin/go version

    /usr/local/go/bin/go run .
# fi
