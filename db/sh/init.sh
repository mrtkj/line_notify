#!/bin/bash
# set -eu

for sqlfile in $IMPORT_FROM_DIR/*.sql; do
  psql -f $sqlfile -d $POSTGRES_DB -U $POSTGRES_USER
done
