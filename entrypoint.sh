#!/bin/sh -l

# Load Env from .env -> These come from GitHub Action Workflow now...
# . ./.env

echo "Hello ${INPUT_GREETINGS}"
time=$(date)
echo "time=$time" >> $GITHUB_OUTPUT

# Golang expects standard names, not GA mangled ones
export APP_DB_HOST=${INPUT_APP_DB_HOST}
export APP_DB_PORT=${INPUT_APP_DB_PORT}
export APP_DB_NAME=${INPUT_APP_DB_NAME}
export APP_DB_USERNAME=${INPUT_APP_DB_USERNAME}
export APP_DB_PASSWORD=${INPUT_APP_DB_PASSWORD}

env | sort

echo Mode: ${INPUT_MODE}

# Check Psql Version and Db Connectivity
/usr/local/go/bin/go version
psql --version
PGPASSWORD=${INPUT_APP_DB_PASSWORD}
psql -Atx "host=${INPUT_APP_DB_HOST} port=${INPUT_APP_DB_PORT} dbname=${INPUT_APP_DB_NAME} user=${INPUT_APP_DB_USERNAME} sslmode=disable" -c 'select current_database()' 

if [ "${INPUT_MODE}" = "test" ]; then
    echo "Testing Mode..."

    cd module && /usr/local/go/bin/go test .

    exit $?
fi

# if [ "${INPUT_MODE}" = "run" ]; then
    echo "Running Mode..."

    /usr/local/go/bin/go run .
# fi
