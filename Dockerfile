# Use Go v1.22
FROM golang:1.22

# Add Postgres Client
RUN apt-get update && apt-get install -y postgresql-client

# Set Up ENV
ENV APP_DB_HOST $INPUT_APP_DB_HOST
ENV APP_DB_PORT $INPUT_APP_DB_PORT
ENV APP_DB_NAME $INPUT_APP_DB_NAME
ENV APP_DB_USERNAME $INPUT_APP_DB_USERNAME
ENV APP_DB_PASSWORD $INPUT_APP_DB_PASSWORD

# Set destination for COPY
WORKDIR /app

# Copy the code over
COPY module/ ./
# Download Go modules and sub modules
RUN go mod download

# Copy the .env
COPY .env ./

# Build
# RUN CGO_ENABLED=0 GOOS=linux go build -o /action-api-ping

# Optional, as overwritten
EXPOSE 8100

# Copies your entrypoint.sh to root
COPY entrypoint.sh /entrypoint.sh

# Docker Container (`entrypoint.sh`), expects 'test' or 'run' mode
ENTRYPOINT ["/entrypoint.sh"]
