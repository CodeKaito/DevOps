# Environment GO
FROM golang:1.20

# Creo una folder sotto AP
RUN mkdir /app

# Workingdir la sposto dentro app
WORKDIR /app

# Copio tutta la mia working directory dentro app
COPY . /app

# Installa i pacchetti e mi connetto ad un redis client
RUN go get github.com/go-redis/redis

# Lancio il command sul main.go
CMD ["go","run","main.go"]
