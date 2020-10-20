FROM golang:latest

WORKDIR /api

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8080

RUN go build -o app

CMD ["./app"]