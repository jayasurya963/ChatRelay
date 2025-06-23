# Dockerfile
FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o chatrelay-bot ./cmd/bot

EXPOSE 3000

CMD ["./chatrelay-bot"]
