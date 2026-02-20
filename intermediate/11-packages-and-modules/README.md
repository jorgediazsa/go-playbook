# 11 - Packages and Modules (Structure and Dependency Hygiene)

Go’s scalability comes from strict dependency rules: packages form a DAG, APIs are explicit, and modules make builds reproducible. Seniors usually stumble not on syntax, but on **boundaries**.

---

## Mental model

- A **package** is the unit of compilation and encapsulation.
- A **module** (go.mod) is the unit of versioning and dependency resolution.
- Imports must form a **Directed Acyclic Graph** (no cycles).

---

## 1) Code organization principles

### Prefer packages by responsibility
- Avoid “utils” dumping grounds.
- Keep packages cohesive: one reason to change.

### Control the public surface
- Export only what must be used externally.
- Keep helpers unexported.
- Consider `internal/` to enforce boundaries.

### `internal/` rules
Any package under `internal/` can only be imported by code within the parent tree.
This is a powerful mechanism to prevent accidental coupling.

---

## 2) Modules and dependency hygiene

### Key commands
- `go mod init` creates a module.
- `go mod tidy` reconciles imports ↔ requirements, removes unused deps.

### Important concepts
- **Semantic import versioning** for major versions (`/v2`).
- `replace` is for local overrides (development, forks). It should not be permanent in released code unless intentional.
- `retract` lets you mark bad versions as retracted.

### Practical hygiene
- Avoid depending on transitive implementation details.
- Avoid `replace` “forever”; prefer real versioned fixes.
- Keep upgrades deliberate; use `go list -m -u all` when auditing.

---

## 3) Workspaces (`go work`)

Workspaces are for multi-module development:
- You can develop multiple modules locally without editing `replace` constantly.
- It’s a developer convenience; it is not meant as a production distribution mechanism.

Common workflow:
- `go work init ./moduleA ./moduleB`
- `go work use ./moduleC`

---

## 4) Import cycles and how to avoid them

If A imports B, B cannot import A. Cycles are usually a design smell.

### Typical fixes
- Extract shared types into a third package (e.g., `domain`, `model`)
- Replace type-level coupling with interfaces
- Push dependencies “down” into leaf packages

### Anti-patterns
- Two “peer” packages importing each other for convenience
- Putting shared state in whichever package you touched last

---

## Common interview traps
- Confusing package vs module responsibilities
- Over-exporting (everything is public)
- Cycles (can’t compile) caused by poor domain boundaries
- Misusing `replace` and not understanding `/v2` module paths

---

## Production checklist
- Packages are cohesive and boundaries are enforced (`internal/` where appropriate)
- No cycles; dependencies are directional
- go.mod is clean (`go mod tidy`)
- Avoid “utility” packages that become dependency magnets

---

## Exercises
These exercises include:
- architectural constraints enforced by tests
- import boundary checks
- a workspace/multi-module scenario
- cycle avoidance with a domain extraction approach
