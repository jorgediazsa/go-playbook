package httpjson

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequireAuthMiddleware(t *testing.T) {
	// The core handler we want to protect
	coreHandler := http.HandlerFunc(HealthHandler)

	// Wrap the core handler in our middleware
	protectedHandler := RequireAuth(coreHandler)

	// Test 1: Missing API Key (Should be blocked)
	req1, _ := http.NewRequest("GET", "/health", nil)
	rr1 := httptest.NewRecorder()

	protectedHandler.ServeHTTP(rr1, req1)

	if rr1.Code != http.StatusUnauthorized {
		t.Fatalf("FAILED: Expected HTTP 401 Unauthorized for missing API key, got %d. Did your middleware block the request?", rr1.Code)
	}

	// Test 2: Invalid API Key (Should be blocked)
	req2, _ := http.NewRequest("GET", "/health", nil)
	req2.Header.Set("X-API-Key", "hacker")
	rr2 := httptest.NewRecorder()

	protectedHandler.ServeHTTP(rr2, req2)

	if rr2.Code != http.StatusUnauthorized {
		t.Fatalf("FAILED: Expected HTTP 401 Unauthorized for invalid API key, got %d.", rr2.Code)
	}

	// Test 3: Valid API Key (Should be allowed through)
	req3, _ := http.NewRequest("GET", "/health", nil)
	req3.Header.Set("X-API-Key", "secret")
	rr3 := httptest.NewRecorder()

	protectedHandler.ServeHTTP(rr3, req3)

	if rr3.Code != http.StatusOK {
		t.Fatalf("FAILED: Expected HTTP 200 OK for valid API key, got %d.", rr3.Code)
	}

	if rr3.Body.String() != "HEALTHY" {
		t.Fatalf("FAILED: Expected body 'HEALTHY', got %q", rr3.Body.String())
	}
}
