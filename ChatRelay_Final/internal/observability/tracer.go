// OpenTelemetry setup for tracing and logging
version: '3.9'

services:
  chatrelay-bot:
    build: .
    container_name: chatrelay-bot
    env_file:
      - .env
    depends_on:
      - mock-backend
    ports:
      - "3000:3000"

  oauth:
    build:
      context: ./oauth
    container_name: chatrelay-oauth
    env_file:
      - .env
    ports:
      - "4000:4000"

  mock-backend:
    build:
      context: ./mock-backend
    container_name: mock-backend
    ports:
      - "5000:5000"

  otel-collector:
    image: otel/opentelemetry-collector:latest
    container_name: otel-collector
    volumes:
      - ./otel-collector-config.yaml:/etc/otelcol/config.yaml
    command: ["--config", "/etc/otelcol/config.yaml"]
    ports:
      - "4317:4317"
      - "55681:55681"

  jaeger:
    image: jaegertracing/all-in-one:latest
    container_name: jaeger
    ports:
      - "6831:6831/udp"
      - "16686:16686"
      - "14250:14250"
