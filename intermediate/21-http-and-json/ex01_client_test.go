package httpjson

import (
	"net/http"
	"testing"
	"time"
)

func TestNewWebhookClient(t *testing.T) {
	client := NewWebhookClient()

	if client == http.DefaultClient || client == nil {
		t.Fatalf("FAILED: You returned the DefaultClient or nil!")
	}

	if client.Timeout != 5*time.Second {
		t.Fatalf("FAILED: Expected Timeout to be 5s, got %v", client.Timeout)
	}

	transport, ok := client.Transport.(*http.Transport)
	if !ok || transport == nil {
		t.Fatalf("FAILED: Expected a configured *http.Transport, got %T", client.Transport)
	}

	if transport.MaxIdleConnsPerHost != 5 {
		t.Fatalf("FAILED: Expected MaxIdleConnsPerHost to be 5, got %d", transport.MaxIdleConnsPerHost)
	}
}
