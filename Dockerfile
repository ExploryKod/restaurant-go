FROM golang:1.21 as Builder

# Active le comportement de module indépendant
ENV GO111MODULE=on

# https://dave.cheney.net/2016/01/18/cgo-is-not-go
ENV CGO_ENABLED=0
ENV GOOS=$GOOS
ENV GOARCH=$GOARCH

WORKDIR /goapp
COPY ./goapp .

RUN go mod download \
    && go mod verify \
    && go mod tidy \
    && go build -o gorillachat

FROM scratch as FINAL

COPY ./ ./app
