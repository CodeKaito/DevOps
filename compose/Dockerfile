FROM golang:latest

WORKDIR /app

COPY ./src/go.mod .

RUN go mod download

RUN go mod verify

COPY ./src/main.go .

RUN go get -d ./...

RUN go build -o /bin/sh

CMD ["/bin/sh"]
