# Stage 1: Build the application
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/main

RUN go build -o /movie-watchlist

# Stage 2: Run the application
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /movie-watchlist .

EXPOSE 9800

CMD [ "movie-watchlist" ]