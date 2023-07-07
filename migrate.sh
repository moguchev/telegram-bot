#!/bin/sh

export MIGRATION_DIR="./migrations"
export DB_DSN="host=localhost port=5432 user=user password=password dbname=wb_helper sslmode=disable"

if [ "$1" = "--dryrun" ]; then
    goose -v -dir ${MIGRATION_DIR} postgres "${DB_DSN}" status
else
    goose -v -dir ${MIGRATION_DIR} postgres "${DB_DSN}" up
fi
