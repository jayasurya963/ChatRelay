# Changelog
All notable changes to this project will be documented in this file.
Initial release of ChatRelay bot 🎉 [1.0.0] - 2025-06-18

## Added
Slack Socket Mode integration.
Backend streaming with chunked messages.
Full OAuth2 Slack install flow.
OpenTelemetry tracing for Slack events and backend calls.
Docker Compose support for bot, backend, OTEL, Jaeger.
Jaeger UI integration.
Unit tests for bot, client.
Sample mock backend implementation.
Add-to-Slack manifest YAML template.

## Changed
Switched from REST backend to streaming SSE simulation.
Improved bot relay logic with concurrency handling.

## Fixed
OAuth callback redirect bugs.
Docker networking issues.

## Security
Secrets moved to .env with .gitignore .
No token leaks in logs.
