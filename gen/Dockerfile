FROM golang

WORKDIR /gen

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .

RUN go build -o /usr/bin/phoenix local/main.go

CMD phoenix