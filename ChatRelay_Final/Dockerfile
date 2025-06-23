# Dockerfile for ChatRelay
FROM golang:1.21
WORKDIR /app
COPY . .
RUN go build -o chatrelay ./cmd/bot
CMD ["./chatrelay"]
