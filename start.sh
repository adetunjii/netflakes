#!/bin/sh

# exit immediately if there is an error
set -e
echo "running db migration...."
/app/migrate -path /app/migration -database <dsn> -verbose up

echo "migration ran successfully"
rm -rf /app/migrate
rm -rf /app/migration
