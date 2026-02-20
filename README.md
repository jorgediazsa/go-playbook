# Go Playbook - Senior Engineer Track

Welcome to the **Go Playbook**, a comprehensive, practical, and production-oriented guide to mastering Go (Golang). This repository is designed for experienced software engineers who want to move beyond basic syntax and learn how to build robust, concurrent, and high-performance systems in Go.

## 🎯 Project Overview

This playbook is hands-on. Instead of long academic documents, it is broken down into tiered topics. Each topic folder contains a `README.md` detailing the production concepts, common pitfalls, and architectural decisions, accompanied by a set of real-world exercises (with failing tests) for you to solve.

The curriculum is divided into three tiers:

---

### 🟢 [Basic Tier](./basic)
Focuses on the fundamentals of the language, but strictly through the lens of Go's unique mechanics, memory behavior, and idioms. 

- [01. Core Syntax](./basic/01-core-syntax)
- [02. Control Flow](./basic/02-control-flow)
- [03. Functions](./basic/03-functions)
- [04. Standard Collections](./basic/04-standard-collections)
- [05. Pointers](./basic/05-pointers)
- [06. Errors (Foundational)](./basic/06-errors-foundational)
- [07. Defer, Panic, Recover](./basic/07-defer-panic-recover)
- [08. Strings, Bytes, and Runes](./basic/08-strings-bytes-and-runes)
- [09. Time](./basic/09-time)

---

### 🟡 [Intermediate Tier](./intermediate)
Focuses on writing idiomatic, safe, and easily testable components. It heavily targets real-world production concurrency primitives and data handling.

- [10. Methods and Composition](./intermediate/10-methods-and-composition)
- [11. Packages and Modules](./intermediate/11-packages-and-modules)
- [12. Tooling & Static Analysis](./intermediate/12-tooling-and-static-analysis)
- [13. Goroutines](./intermediate/13-goroutines)
- [14. Channels](./intermediate/14-channels)
- [15. Synchronization Primitives](./intermediate/15-synchronization-primitives)
- [16. Context](./intermediate/16-context)
- [17. Idiomatic Error Handling (Advanced)](./intermediate/17-idiomatic-error-handling-advanced)
- [18. Testing](./intermediate/18-testing)
- [19. Filesystem & OS](./intermediate/19-filesystem-and-os)
- [20. Database & IO Patterns](./intermediate/20-database-and-io-patterns)
- [21. HTTP and JSON](./intermediate/21-http-and-json)

---

### 🔴 [Advanced Tier](./advanced)
Focuses on the deepest parts of Go: memory management, compiler optimizations, architectural patterns at scale, and observability.

- [22. Advanced Concurrency Patterns](./advanced/22-advanced-concurrency-patterns)
- [23. Memory Model](./advanced/23-memory-model)
- [24. Garbage Collector](./advanced/24-garbage-collector)
- [25. Generics](./advanced/25-generics)
- [26. Reflection](./advanced/26-reflection)
- [27. Profiling & Observability](./advanced/27-profiling-and-observability)
- [28. Go in Production](./advanced/28-go-in-production)
- [29. Go Infrastructure](./advanced/29-go-infrastructure)

---

## 🚀 How to Use This Playbook

1. **Pick a Topic:** Navigate to a topic folder (e.g., `basic/04-standard-collections`).
2. **Read the `README.md`:** Learn the production context, pitfalls, and idiomatic patterns.
3. **Run the Tests:** Run `go test -v ./...` inside the folder. The tests will fail intentionally.
4. **Solve the Exercises:** Open the `.go` files, follow the `TODO` markers, and refactor the buggy or incomplete code.
5. **Verify:** Keep running the tests (and `go test -race ./...` for concurrency) until you've successfully implemented the resilient solution.

## 📖 Roadmap

For a detailed breakdown of every single sub-topic covered within these sections, refer to the [roadmap.md](./roadmap.md).
