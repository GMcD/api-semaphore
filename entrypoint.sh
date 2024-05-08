#!/bin/sh -l

# Load Env from .env
. ./.env

echo "Hello ${INPUT_}"
time=$(date)
echo "time=$time" >> $GITHUB_OUTPUT

env | sort

echo Mode: ${INPUT_MODE}

# Check Psql Version and Db Connectivity
/usr/local/go/bin/go version
psql --version
PGPASS=${INPUT_APP_DB_PASSWORD}
psql -Atx "hostaddr=${INPUT_APP_DB_HOST} port=${INPUT_APP_DB_PORT} dbname=${INPUT_APP_DB_NAME} user=${INPUT_APP_DB_USERNAME} sslmode=disable" -c 'select current_database()' 

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
