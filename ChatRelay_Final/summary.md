
# Project Name: ChatRelay

Goal: Create a production-ready Slack bot that forwards user messages to an AI chat backend and streams replies back to Slack in real-time.

Context: High-performance messaging in team collaboration environments, using Go for speed and reliability.

## Approach :
--Language: Go (Golang) – chosen for its speed, concurrency support, and tooling.
--Architecture: Clean separation via cmd/, internal/, and config/.
--Streaming: Supports SSE (Server-Sent Events) for real-time streaming to Slack.
--Fallback: Graceful fallback to full JSON response if backend doesn't support SSE.
--Slack Integration: Uses Slack Events API with Add-to-Slack OAuth flow.
--Observability: Integrated OpenTelemetry for traces and structured logging.
--Testing: Includes unit tests and integration tests.
--Mock Backend: Simulates real backend for both SSE and full response testing.
--Dockerized: Both bot and mock backend are containerized with Docker and Docker Compose.

## I Have Built :
--Slack Bot – Listens for messages and relays them.
--Chat Backend Client – Handles both streaming and full replies.
--Mock Backend – Simulates /v1/chat/stream and /v1/chat/full endpoints.
--OpenTelemetry Setup – Adds tracing, logging, and metrics support.
--OAuth Flow – Includes Slack app manifest and login flow boilerplate.
--Unit Tests – Tests for message handlers and backend logic.
--Integration Test – Verifies end-to-end message flow.
--README.md, LICENSE, .env, .gitignore – Clean project packaging.
--Docker Support – Easily deploy locally or in production.

# Outcome :
--Fully functional and tested Slack bot for streaming chat backend interaction.
--Modular, clean, and extensible code structure.
--Local dev environment runs with docker-compose and mock backend.
--Production practices followed (env config, telemetry, clean commits).
--Ready for deployment, showcasing Go concurrency, Slack integration, and streaming response handling.
