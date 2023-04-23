FROM golang:1.20.3-alpine as ops

RUN apk update && apk add git
RUN go install github.com/cosmtrek/air@latest

WORKDIR /app
CMD ["air","-c",".air.toml"]
