package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type ChatRequest struct {
	UserID string `json:"user_id"`
	Query  string `json:"query"`
}

type FullResponse struct {
	FullResponse string `json:"full_response"`
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	var req ChatRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}

	chunks := []string{
		"Goroutines are lightweight",
		"concurrent execution units in Go.",
		"They allow for massive parallelism.",
		"Each goroutine has a tiny stack.",
		"They are multiplexed onto threads."}

	for i, chunk := range chunks {
		event := fmt.Sprintf("id: %d\nevent: message_part\ndata: {\"text_chunk\": \"%s\"}\n\n", i+1, chunk)
		_, _ = w.Write([]byte(event))
		flusher.Flush()
		time.Sleep(time.Duration(rand.Intn(500)+300) * time.Millisecond)
	}

	endEvent := "id: 999\nevent: stream_end\ndata: {\"status\": \"done\"}\n\n"
	_, _ = w.Write([]byte(endEvent))
	flusher.Flush()
}

func fullHandler(w http.ResponseWriter, r *http.Request) {
	var req ChatRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	response := FullResponse{
		FullResponse: "Goroutines are lightweight, concurrent execution units in Go. They allow for massive parallelism.",
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/v1/chat/stream", streamHandler)
	http.HandleFunc("/v1/chat/full", fullHandler)
	log.Println("[mock_backend] running at http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
