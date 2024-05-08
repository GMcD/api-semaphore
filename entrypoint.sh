#!/bin/sh -l

# Load Env from .env
. ./.env

echo "Hello ${INPUT_WHO-TO-GREET}"
time=$(date)
echo "time=$time" >> $GITHUB_OUTPUT

env | sort

echo Mode: ${INPUT_MODE}

# Check Psql Version and Db Connectivity
/usr/local/go/bin/go version
psql --version
PGPASS=${APP_DB_PASSWORD}
psql -Atx "hostaddr=${APP_DB_HOST} port=${APP_DB_PORT} dbname=${APP_DB_NAME} user=${APP_DB_USERNAME} sslmode=disable" -c 'select current_database()' 

if [ "${INPUT_MODE}" = "test" ]; then
    echo "Testing Mode..."
    env | sort

    /usr/local/go/bin/go test .

    exit $?
fi

# if [ "${INPUT_MODE}" = "run" ]; then
    echo "Running Mode..."

    /usr/local/go/bin/go run .
# fi
