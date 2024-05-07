#!/bin/sh -l

# Load Env from .env
. ./.env

echo "Hello $1"
time=$(date)
echo "time=$time" >> $GITHUB_OUTPUT

echo Mode: $2

if [ "$2" = "test" ]; then
    echo "Testing Mode..."
    env | sort

    psql --version

    PGPASS=${APP_DB_PASSWORD}
    psql -Atx "hostaddr=${APP_DB_HOST} dbname=${APP_DB_NAME} user=${APP_DB_USERNAME} sslmode=disable" -c 'select current_database()' 

    PGPASS=${APP_DB_PASSWORD}
    psql "hostaddr=35.242.149.106 dbname=${APP_DB_NAME} user=${APP_DB_USERNAME} sslmode=disable" -c 'select current_database()'

    /usr/local/bin/go test .

    exit $?
fi

# if [ "$2" = "run" ]; then
    echo "Running Mode..."
    /usr/local/bin/go version

    /usr/local/bin/go run .
# fi
