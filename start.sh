#!/bin/sh

set -e

chmod +x /app/start.sh 

echo "run db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"
