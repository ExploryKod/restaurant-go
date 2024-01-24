#!/bin/bash

# start air directly once docker container is set up
if [ "$1" == "air" ]; then
docker exec -it go-api sh -c "go install github.com/cosmtrek/air@latest"

docker exec -it go-api sh -c "air init"

# Modify .air.toml to set poll = true
docker exec -it go-api sed -i 's/poll = false/poll = true/' .air.toml

docker exec -it go-api sh -c "air -c .air.toml"
fi

# start only air init once docker container is set up
if [ "$1" == "air-init" ]; then
docker exec -it go-api sh -c "air init"

# Modify .air.toml to set poll = true
docker exec -it go-api sed -i 's/poll = false/poll = true/' .air.toml

docker exec -it go-api sh -c "air -c .air.toml"
fi

# start only air init
if [ "$2" == "air-init" ]; then
docker exec -it go-api sh -c "air init"

# Modify .air.toml to set poll = true
docker exec -it go-api sed -i 's/poll = false/poll = true/' .air.toml

docker exec -it go-api sh -c "air -c .air.toml"
fi

# Restart air without build directly if a container is already up
if [ "$1" == "restart-air" ]; then
docker exec -it go-api sh -c "air -c .air.toml"
fi

# Restart air without build
if [ "$1" == "restart-air" ]; then
docker exec -it go-api sh -c "air -c .air.toml"
fi
