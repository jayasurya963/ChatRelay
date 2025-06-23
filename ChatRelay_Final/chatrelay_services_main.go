# ChatRelay Slack Bot 🧵

ChatRelay is a high-performance Slack bot built in Go that relays user messages to a chat backend (mock or real) and streams responses back in real-time. Designed for concurrency, observability, and production-readiness.

---

## 🔧 Features
- Slack App with Socket Mode support
- Real-time message relay to chat backend
- Streaming responses (SSE or JSON chunked)
- Modular microservice architecture
- OpenTelemetry tracing (Jaeger-ready)
- Secure OAuth2 Add-to-Slack flow
- Dockerized with `docker-compose`

---

## 🧱 Architecture Overview

```
Slack App ─► ChatRelay Bot ─► Chat Backend (Mock/Real)
   ▲             │
   │             └── OpenTelemetry (tracing)
   ▼
OAuth Server (Slack Install Flow)
```

Services:
- `chatrelay-bot`: Main Slack bot
- `oauth-server`: Handles Slack's OAuth install
- `mock-backend`: Simulates chat backend for local testing
- `otel-collector`: Observability pipeline
- `jaeger`: Trace viewer

---

## 🚀 Getting Started

### 1. Clone & Configure
```bash
git clone https://github.com/your-username/chatrelay.git
cd chatrelay
cp .env.example .env  # edit with your credentials
```

### 2. Set Up Slack App
- Go to https://api.slack.com/apps → Create App (from manifest)
- Use the included `slack-manifest.yaml`
- Add Bot Token Scopes: `chat:write`, `app_mentions:read`, etc.
- Enable Socket Mode & Install App

### 3. Run Locally with Docker Compose
```bash
docker-compose up --build
```

Visit:
- `http://localhost:4000/oauth/start` – Add to Slack
- `http://localhost:16686` – Jaeger UI for traces

---

## 📁 Directory Structure

```
cmd/
├── bot/          # Main entry for Slack bot
├── oauth/        # OAuth2 redirect + install
└── mock/         # Fake backend chat responder

internal/
├── bot/          # Slack listener + relayer
├── backend/      # Backend client (stream/fetch)
├── config/       # Loads env vars
└── observability/# OpenTelemetry setup

.env               # Environment config
otel-collector-config.yaml
slack-manifest.yaml
```

---

## ✅ Environment Variables (.env)
```
SLACK_BOT_TOKEN=...
SLACK_CLIENT_ID=...
SLACK_CLIENT_SECRET=...
SLACK_SIGNING_SECRET=...
SLACK_APP_TOKEN=...
OAUTH_REDIRECT_URI=http://localhost:4000/oauth/callback
CHAT_BACKEND_URL=http://mock-backend:5000/v1/chat/stream
OTEL_EXPORTER_OTLP_ENDPOINT=http://otel-collector:4317
```

---

## 🧪 Tests

### Unit Tests
```bash
go test ./internal/bot
```

### Integration Tests (manual for now)
- Send a message to the Slack app
- Watch it relay and stream responses

---

## 📦 Deployment
You can deploy each component independently via Docker or Kubernetes. OpenTelemetry can be configured with your production observability stack (e.g., Prometheus, Grafana).

---

## 📜 License
MIT

---

## 👨‍💻 Author
Built by [Your Name] as part of a take-home engineering challenge. Contributions welcome!

---

## 📝 TODOs
- [ ] Add full integration test suite
- [ ] Add metrics collection
- [ ] Extend OAuth handler for automatic token exchange
- [ ] Add Redis or PubSub support for fanout scalability
