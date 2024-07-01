FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/bin/app cmd/main/main.go

FROM alpine:3.18

RUN apk --no-cache add libc6-compat

WORKDIR /app

COPY --from=builder /app/bin/app /app/app

EXPOSE 9800

CMD ["/app/app"]
