package infrastructure

import (
	"net/http"
)

// Context: Health and Readiness Endpoints
// You are deploying an application to Kubernetes. Kubernetes uses two endpoints
// to monitor your app:
// - Liveness Probe (`/healthz`): "Is the application running?" (If no, K8s restarts the Pod).
// - Readiness Probe (`/readyz`): "Is the application ready to accept traffic?" (If no, K8s removes it from the Load Balancer).
//
// Why this matters: If your database goes down, your app is still *alive*, but it
// is no longer *ready* to serve traffic. If `/healthz` checks the DB and fails, K8s
// will restart your Pod infinitely, creating a restart-loop outage!
//
// Requirements:
//  1. You must implement the two HTTP handlers below correctly.
//  2. `HealthzHandler` should simply return an HTTP 200 OK with `{"status":"UP"}` instantly.
//     Do not check the DB.
//  3. `ReadyzHandler` must call `checkDBConnection()` (provided). If the DB fails,
//     return an HTTP 503 Service Unavailable with `{"status":"DOWN", "error":"db failed"}`.
//     If the DB succeeds, return an HTTP 200 OK with `{"status":"READY"}`.
//
// Assume this function exists to check the DB:
var checkDBConnection = func() error {
	return nil // mock success
}

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	// BUG: Does nothing.
	// TODO: Return 200 OK with `{"status":"UP"}`
}

func ReadyzHandler(w http.ResponseWriter, r *http.Request) {
	// BUG: Does nothing.
	// TODO: Call checkDBConnection()
	// TODO: If error -> 503 {"status":"DOWN"}
	// TODO: If success -> 200 {"status":"READY"}
}
