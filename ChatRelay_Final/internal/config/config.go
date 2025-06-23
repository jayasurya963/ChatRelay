package config

import (
	"log"
	"os"
)

type Config struct {
	SlackBotToken   string
	SlackAppToken   string
	BackendURL      string
	TelemetryURL    string
}

func Load() *Config {
	cfg := &Config{
		SlackBotToken:   os.Getenv("SLACK_BOT_TOKEN"),
		SlackAppToken:   os.Getenv("SLACK_APP_TOKEN"),
		BackendURL:      os.Getenv("CHAT_BACKEND_URL"),
		TelemetryURL:    os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"),
	}
	if cfg.SlackBotToken == "" || cfg.SlackAppToken == "" || cfg.BackendURL == "" {
		log.Fatal("Missing required environment variables")
	}
	return cfg
}

