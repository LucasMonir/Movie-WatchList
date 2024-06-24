FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/main

RUN go build -o /movie-watchlist

EXPOSE 9800

CMD [ "/movie-watchlist" ]
