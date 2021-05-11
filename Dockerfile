FROM golang:1.16

WORKDIR /src/app
COPY . .

RUN go build .

ENV GIN_MODE="release"

CMD ["./proxygeoip"]