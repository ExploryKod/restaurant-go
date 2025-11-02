# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.22 AS Builder  
WORKDIR /app

COPY . .
RUN go mod download \
    && go mod verify

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o /build/restaurantgo ./main/main.go

# Deploy the application binary into a lean image
# Using alpine instead of scratch for better compatibility with Render.com
FROM alpine:latest AS build-release-stage

# Install CA certificates for HTTPS connections (required for external database connections)
RUN apk --no-cache add ca-certificates

WORKDIR /appgo

COPY --from=Builder /build/restaurantgo ./restaurantgo
# Copier le dossier src pour servir les fichiers statiques (CSS, JS, images)
COPY --from=Builder /app/src ./src

# Expose the port the app runs on
EXPOSE 9999

ENTRYPOINT ["./restaurantgo"]
