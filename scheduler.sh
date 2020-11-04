#! /bin/sh

curl -X GET "${APP_URL}api/schedules/exec" -H "Authorization: Bearer ${ACCESS_TOKEN}"