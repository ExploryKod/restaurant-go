FROM golang:1.21

# Active le comportement de module indépendant
ENV GO111MODULE=on

# Je vais faire une build en 2 étapes
# https://dave.cheney.net/2016/01/18/cgo-is-not-go
ENV CGO_ENABLED=0
ENV GOOS=$GOOS
ENV GOARCH=$GOARCH

WORKDIR /app
ENV MODE "dev-docker"
COPY . /app

#RUN go install github.com/cosmtrek/air@latest
#RUN apt-get update
#RUN apt-get install nano -y

RUN go mod download \
    && go mod verify \
    && go mod tidy \
    && go build -o restaurant-go-dev

