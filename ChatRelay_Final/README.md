# ChatRelay - High-Performance Slack Bot in Go

ChatRelay is a production-grade Slack bot written in Go that listens to messages, forwards them to a chat backend (streaming or complete response), and streams replies back to users in Slack. It features concurrency, observability with OpenTelemetry, OAuth installation flow, and clean architecture for scalability.

---

## 🚀 Features

* Socket Mode Slack integration using `slack-go/slack`
* Streaming and complete response handling
* High-performance concurrency with goroutines and channels
* Resilient error handling and retries
* OpenTelemetry-based observability (tracing + structured logging)
* Mock backend for development/testing
* OAuth 2.0 install flow
* Unit + integration tests
* Dockerized environment with `docker-compose`

---

## 📁 Project Structure

```bash

```

---

## ⚙️ Setup & Run Instructions

### 1. Clone & Configure

```bash
git clone https://github.com/yourusername/chatrelay.git
cd chatrelay
cp .env.example .env
```

Edit `.env` with your real values:

```
SLACK_BOT_TOKEN=xoxb-...
SLACK_APP_TOKEN=xapp-...
SLACK_CLIENT_ID=...
SLACK_CLIENT_SECRET=...
SLACK_SIGNING_SECRET=...
CHAT_BACKEND_URL=http://mockbackend:8081/v1/chat/stream
OTEL_EXPORTER_OTLP_ENDPOINT=http://otel-collector:4317
```

### 2. Start Bot + Backend + Jaeger

```bash
docker-compose up --build
```

### 3. Create Slack App

Go to [https://api.slack.com/apps](https://api.slack.com/apps) → Create App:

* Enable **Socket Mode** and **Event Subscriptions**
* Scopes:

  * `chat:write`
  * `channels:history`
  * `im:history`
  * `app_mentions:read`
  * `commands`
* Add `https://<your-ngrok-url>/oauth/callback` as Redirect URL

Install the app to your workspace and note tokens.

### 4. Test Interaction

* Mention the bot in a channel (e.g., `@chatrelay what is Go?`)
* It will stream back chunks from the mock backend

---

## 📡 Observability

* Visit `http://localhost:16686` to view traces in Jaeger
* All logs include `trace_id`, `span_id`
* Traces include:

  * Slack event receive
  * Backend request/response
  * Slack message stream back

---

## 🧪 Testing

Run all tests:

```bash
go test ./...
```

Run individual test file:

```bash
go test -v ./test/bot_test.go
```

---

## 📦 Docker Compose Services

* `chatrelay-bot` — the Slack bot
* `mockbackend` — simulated backend
* `otel-collector` — OpenTelemetry collector
* `jaeger` — UI for traces/logs

---

## 🚀 Slack App Directory Publishing Checklist

* [x] Fully working Slack app with OAuth 2.0 flow
* [x] Secure token handling via env vars
* [x] App manifest YAML included
* [x] Logging and telemetry enabled
* [x] README + LICENSE (MIT)
* [x] Docker support
* [ ] Add-to-Slack button on marketing site

---



---

## 🧠 Design Decisions

* Socket Mode for easier dev (no public endpoint required)
* SSE + JSON mock backend to simulate chat response
* Channel-based stream relay for scalability
* OpenTelemetry to trace message lifecycle from Slack to backend to user

---

## 📈 Future Enhancements

* Slack interactive components (e.g., buttons)
* Prometheus metrics via OpenTelemetry
* Full OAuth + stateful user session handling
* Redis-backed stream buffering
* Rate-limit handling and exponential backoff

```

