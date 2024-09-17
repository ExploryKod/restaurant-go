# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.22 AS Builder  
WORKDIR /app

COPY . .
RUN go mod download \
    && go mod verify

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o /build/restaurantgo ./main/main.go

# Deploy the application binary into a lean image
FROM scratch AS build-release-stage

WORKDIR /appgo

ENV BDD_PORT=bg34o0geswbybq906ljp-mysql.services.clever-cloud.com:3306
ENV BDD_NAME=bg34o0geswbybq906ljp
ENV BDD_USER=ueill1e8a2djeyha
ENV BDD_PASSWORD=qbpkI0vidAOA5qha2Q6X


COPY --from=Builder /build/restaurantgo ./restaurantgo

ENTRYPOINT ["./restaurantgo"]
