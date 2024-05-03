#!/bin/sh -l

echo "Hello $1"
time=$(date)
echo "time=$time" >> $GITHUB_OUTPUT

if [[ $2 == "test" ]]; then
    echo "Testing Mode..."
    env | sort

    psql --version

    PGPASS=postgres
    psql -Atx "host=postgres port=5432 dbname=postgres user=postgres" -c 'select current_date' 

    APP_DB_HOST=postgres /usr/local/go/bin/go test .
fi

if [[ $2 == "run" ]]; then
    echo "Running Mode..."
    /usr/local/go/bin/go version
fi
