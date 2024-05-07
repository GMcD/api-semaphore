#!/bin/bash

# Simple Connection Details
export REGION=europe-west2
export INSTANCE=golang-psql
export DB_NAME=golang-db

# Define Project
export PROJECT_ID=gmcd-414115
gcloud config set project ${PROJECT_ID}

# Delete Assets
gcloud sql instances delete ${INSTANCE}
