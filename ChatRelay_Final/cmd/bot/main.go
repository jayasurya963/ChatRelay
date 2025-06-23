// main.go
package main

import (
	"chatrelay/internal/backend"
	"chatrelay/internal/bot"
	"chatrelay/internal/config"
	"chatrelay/internal/otel"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		log.Println("Shutting down ChatRelay bot...")
		cancel()
	}()

	// Load config
	cfg := config.Load()

	// Initialize OpenTelemetry
	ool, err := otel.InitTracerProvider(cfg)
	if err != nil {
		log.Fatalf("failed to initialize OpenTelemetry: %v", err)
	}
	defer ol.Shutdown(ctx)

	// Init chat backend client
	chatClient := backend.NewClient(cfg.BackendURL)

	// Start Slack bot
	err = bot.Start(ctx, cfg, chatClient)
	if err != nil {
		log.Fatalf("failed to start bot: %v", err)
	}
}
