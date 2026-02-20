# 29 - Go Infrastructure and Deployment

Writing the code is only half the battle. Delivering a tiny, fast, statically-linked binary to production is why Go dominates the cloud infrastructure space.

---

## 1. Cross-Compilation

Go can compile binaries for almost any OS and Architecture without needing a complex toolchain natively installed.

```bash
# Compile for Linux (AMD64) from your Mac
GOOS=linux GOARCH=amd64 go build -o myapp-linux main.go

# Compile for Apple Silicon (ARM64)
GOOS=darwin GOARCH=arm64 go build -o myapp-mac main.go
```

---

## 2. Injecting Build Metadata (`-ldflags`)

If a bug happens in production, you need to know *exactly* which Git commit is currently running. You should not hardcode this in source code!

We use linker flags (`ldflags`) to inject strings at compile time:

```go
package main
var Version string = "development"
```

```bash
go build -ldflags="-X 'main.Version=v1.2.3'" main.go
```

When the app boots, `Version` will contain `"v1.2.3"` instead of `"development"`.

---

## 3. CGO and Static Linkage

By default, Go tries to build **Statically Linked** binaries. This means every library (even the Go runtime) is packed into a single standalone file. You can drop it on an Alpine Linux server and it just runs.

However, if you import certain packages (like `net` or `sqlite3` driver) or enable `CGO_ENABLED=1`, Go might dynamically link to the host OS's C libraries (like `libc`). This breaks portability!
**Best Practice:** Compile with `CGO_ENABLED=0` to guarantee a 100% pure, static binary.

---

## 4. Distroless Docker Images

Because Go binaries are static, they do not need Ubuntu, Debian, or even Alpine to run. They don't even need `bash`!

A standard production Dockerfile uses multi-stage builds:
1. Stage 1 (Builder): Use `golang:1.21` to compile the binary.
2. Stage 2 (Runner): Use `gcr.io/distroless/static`. Copy the binary over.

The resulting Docker image is often `< 20MB` and contains NO shell, making it infinitely faster to pull and incredibly secure against hackers.

---

## Exercises

- `ex01_ldflags.go`
- `ex02_healthcheck.go`
