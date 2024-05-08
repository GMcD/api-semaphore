#!/bin/sh -l

# Load Env from .env
. ./.env

echo "Hello $1"
time=$(date)
echo "time=$time" >> $GITHUB_OUTPUT

env | sort

echo Mode: $2

# Check Db Connectivity
PGPASS=${APP_DB_PASSWORD}
psql -Atx "hostaddr=${APP_DB_HOST} port=${APP_DB_PORT} dbname=${APP_DB_NAME} user=${APP_DB_USERNAME} sslmode=disable" -c 'select current_database()' 

if [ "$2" = "test" ]; then
    echo "Testing Mode..."
    env | sort

    psql --version

    /usr/local/go/bin/go test .

    exit $?
fi

# if [ "$2" = "run" ]; then
    echo "Running Mode..."
    /usr/local/go/bin/go version

    /usr/local/go/bin/go run .
# fi
