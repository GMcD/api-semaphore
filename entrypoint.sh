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

    PGPASS=${DB_PASS}
    psql -Atx "host=${INSTANCE_UNIX_SOCKET} dbname=${DB_NAME} user=${DB_USER}" -c 'select current_date' 

    APP_DB_HOST=${INSTANCE_UNIX_SOCKET} /usr/local/go/bin/go test .

    exit $?
fi

# if [ "$2" = "run" ]; then
    echo "Running Mode..."
    /usr/local/go/bin/go version

    APP_DB_HOST=${INSTANCE_UNIX_SOCKET} /usr/local/go/bin/go run .
# fi
