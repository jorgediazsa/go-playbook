package infrastructure

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestK8sProbes(t *testing.T) {
	// Test Healthz Always Passes
	reqHealth, _ := http.NewRequest("GET", "/healthz", nil)
	rrHealth := httptest.NewRecorder()
	HealthzHandler(rrHealth, reqHealth)

	if rrHealth.Code != http.StatusOK {
		t.Fatalf("FAILED: /healthz must return 200 OK. Got %d", rrHealth.Code)
	}

	if !strings.Contains(rrHealth.Body.String(), `"status":"UP"`) {
		t.Fatalf("FAILED: /healthz output missing expected status JSON")
	}

	// Mock DB check success
	checkDBConnection = func() error {
		return nil
	}

	reqReady, _ := http.NewRequest("GET", "/readyz", nil)
	rrReady := httptest.NewRecorder()
	ReadyzHandler(rrReady, reqReady)

	if rrReady.Code != http.StatusOK {
		t.Fatalf("FAILED: /readyz must return 200 OK when DB is healthy. Got %d", rrReady.Code)
	}

	// Mock DB check failure
	checkDBConnection = func() error {
		return errors.New("timeout")
	}

	reqReadyFail, _ := http.NewRequest("GET", "/readyz", nil)
	rrReadyFail := httptest.NewRecorder()
	ReadyzHandler(rrReadyFail, reqReadyFail)

	// In Kubernetes, if the Readiness probe returns < 200 or >= 400, K8s considers it not ready.
	if rrReadyFail.Code != http.StatusServiceUnavailable {
		t.Fatalf("FAILED: /readyz must return 503 Service Unavailable when DB is down. Got %d", rrReadyFail.Code)
	}
}
