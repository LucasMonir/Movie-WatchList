FROM golang:1.22 AS builder

WORKDIR /logger

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /logger/bin/logger cmd/main/main.go

FROM alpine:3.18

RUN apk --no-cache add libc6-compat

WORKDIR /logger

COPY --from=builder /logger/bin/logger /logger/logger

CMD ["/logger/logger"]
