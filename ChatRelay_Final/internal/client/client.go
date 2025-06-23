// Interacts with chat backend (streaming or full response)
package backend

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Client struct {
	url string
}

type ChatRequest struct {
	UserID string `json:"user_id"`
	Query  string `json:"query"`
}

type StreamHandler func(part string, done bool)

func NewClient(url string) *Client {
	return &Client{url: url}
}

func (c *Client) StreamResponse(ctx context.Context, req ChatRequest, handler StreamHandler) error {
	data, _ := json.Marshal(req)
	r, err := http.NewRequestWithContext(ctx, "POST", c.url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	for decoder.More() {
		var chunk map[string]string
		if err := decoder.Decode(&chunk); err != nil && err != io.EOF {
			log.Printf("error decoding chunk: %v", err)
			continue
		}
		if text, ok := chunk["text_chunk"]; ok {
			handler(text, false)
		}
		if chunk["status"] == "done" {
			handler("", true)
			break
		}
	}
	return nil
}
