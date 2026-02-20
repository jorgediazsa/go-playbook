package infrastructure

// Context: Build Metadata via `ldflags`
// You are deploying a microservice. You need the `/version` endpoint to output
// the exact Git Commit Hash and the Build Time so Datadog can track deployments.
//
// Why this matters: You could read a `.env` file, but that's slow and fragile.
// Go allows the CI/CD pipeline (GitHub Actions) to inject strings directly into
// the binary at compile time.
//
// Requirements:
// 1. Declare two package-level variables: `GitCommit` and `BuildTime`.
// 2. Initialize them with default values (e.g., "unknown").
// 3. Refactor `GetBuildInfo` to return these strings formatted gracefully.
//
// To actually test injecting values locally, you would run:
// `go test -ldflags="-X 'go-playbook/advanced/29-go-infrastructure.GitCommit=abc1234' -X 'go-playbook/advanced/29-go-infrastructure.BuildTime=2023-01-01'" ./...`

// TODO: Declare GitCommit and BuildTime vars here
// var GitCommit string = "unknown"
// ...

// BUG: It currently returns hardcoded strings!
// TODO: Return the global variables that you declared so they can be overridden by `ldflags`.
func GetBuildInfo() (commit string, time string) {
	return "hardcoded", "hardcoded"
}
