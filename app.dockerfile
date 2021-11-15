FROM golang:alpine

WORKDIR /game-currency

ADD . .

RUN go mod download

RUN go build -o game-currency ./app/main.go

ENTRYPOINT ["./game-currency"]
