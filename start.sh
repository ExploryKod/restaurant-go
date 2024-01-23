#!/bin/bash
if [ "$1" == "start" ]; then
  # If the argument is "start", bring the containers up (use it for first time)
  docker-compose up -d --build
fi

if [ "$1" == "start-run" ]; then
  # If the argument is "start", bring the containers up (use it for first time)
  docker-compose up -d --build
  docker exec -it restaurant-go sh -c "go run main/main.go"
fi

if [ "$1" == "gorun" ]; then
  docker exec -it restaurant-go sh -c "go run main/main.go"
fi


##!/bin/bash
#set -e
#
#if [ "$1" = 'postgres' ]; then
#    chown -R postgres "$PGDATA"
#
#    if [ -z "$(ls -A "$PGDATA")" ]; then
#        gosu postgres initdb
#    fi
#
#    exec gosu postgres "$@"
#fi
#
#exec "$@"
