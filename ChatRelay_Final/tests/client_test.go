// Unit test for backend client
package backend_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"chatrelay/backend"
)

func TestStreamChat(t *testing.T) {
	// Simulate a streaming backend with chunked response
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			t.Fatal("expected http.Flusher")
		}
		w.Header().Set("Content-Type", "application/json")
		for _, chunk := range []string{"Hi", " there"} {
			w.Write([]byte(chunk))
			flusher.Flush()
			time.Sleep(10 * time.Millisecond)
		}
	}))
	defer ts.Close()

	client := backend.NewClient(ts.URL)
	ctx := context.Background()
	ch, err := client.StreamChat(ctx, "U123", "hello")
	assert.NoError(t, err)

	var result string
	for c := range ch {
		result += c
	}

	assert.Equal(t, "Hi there", result)
}
