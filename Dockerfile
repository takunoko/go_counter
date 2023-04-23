FROM golang:1.20.3-alpine

WORKDIR /app
CMD ["go","run","main.go"]
