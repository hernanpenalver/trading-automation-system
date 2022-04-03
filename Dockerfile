# syntax=docker/dockerfile:1
# Build stage
FROM golang:1.18.0-alpine3.15 AS builder

ENV GO111MODULE=on
WORKDIR /app
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o main ./api/cmd/main.go

#Run stage
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .

# Expose port 9000 to the outside world
EXPOSE 9000
CMD /app/main