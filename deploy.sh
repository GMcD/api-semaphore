#!/bin/bash

# Simple Connection Details
export REGION=europe-west2
export SERVICE_ACCOUNT=sa-gmcd-gcr@gmcd-414115.iam.gserviceaccount.com
export INSTANCE=golang-psql
export DB_NAME=golang-db
export DB_USER=golang-user
export PASSWORD=Noddy.Pass

# Define Project
export PROJECT_ID=gmcd-414115
gcloud config set project ${PROJECT_ID}

# Enable Cloud Run for golang-api
gcloud services enable run.googleapis.com

# Enable Cloud SQL and friends
gcloud services enable compute.googleapis.com sqladmin.googleapis.com \
   containerregistry.googleapis.com cloudbuild.googleapis.com servicenetworking.googleapis.com

# Create Insecure Server Instance for transient data
gcloud sql instances create ${INSTANCE} \
   --database-version=POSTGRES_15 \
   --cpu=1 \
   --memory=4GB \
   --region=europe-west2 \
   --root-password=${PASSWORD}

# Add a Db and User
gcloud sql databases create ${DB_NAME} --instance=${INSTANCE}
gcloud sql users create ${DB_USER} \
   --instance=${INSTANCE} \
   --password=${PASSWORD}
   
# Build API
gcloud builds submit --tag gcr.io/${PROJECT_ID}/golang-api .

# Add Cloud SQL
gcloud projects add-iam-policy-binding ${PROJECT_ID} \
  --member="serviceAccount:${SERVICE_ACCOUNT}" \
  --role="roles/cloudsql.client"

# Deploy API Container
gcloud run deploy golang-api \
   --memory 2G \
   --platform managed \
   --image gcr.io/${PROJECT_ID}/golang-api \
   --add-cloudsql-instances ${PROJECT_ID}:${REGION}:${DB_NAME} \
   --set-env-vars TARGET=Gary \
   --set-env-vars INSTANCE_UNIX_SOCKET=/cloudsql/414115:${REGION}:${DB_NAME} \
   --set-env-vars DB_NAME=${DB_NAME} \
   --set-env-vars DB_USER=${DB_USER} \
   --set-env-vars DB_PASS=${PASSWORD} \
   --allow-unauthenticated

# Connect as Root User
gcloud sql connect ${INSTANCE} --user=postgres --quiet

# Connect as the App User with GCloud
echo ${PASSWORD} | gcloud sql connect ${INSTANCE} -d=${DB_NAME} -u=${DB_USER} --quiet

# Connect as the App User with Psql
PGPASS=${APP_DB_PASSWORD}
psql "hostaddr=35.242.149.106 dbname=${APP_DB_NAME} user=${APP_DB_USERNAME} sslmode=disable" -c 'select current_database()'


# # Deploy DB Container
# gcloud run deploy golang-sql \
#    --memory 2G \
#    --platform managed \
#    --image gcr.io/${PROJECT_ID}/golang-sql \
#    --add-cloudsql-instances ${PROJECT_ID}:${REGION}:${DB_NAME} \
#    --set-env-vars INSTANCE_UNIX_SOCKET=/cloudsql/414115:${REGION}:${DB_NAME} \
#    --set-env-vars DB_NAME=${DB_NAME} \
#    --set-env-vars DB_USER=${DB_USER} \
#    --set-env-vars DB_PASS=${PASSWORD} \
#    --allow-unauthenticated
