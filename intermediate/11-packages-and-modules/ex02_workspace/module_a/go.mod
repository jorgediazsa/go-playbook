module github.com/enterprise/module_a

go 1.22

// BUG: module_a depends on module_b, but module_b is NOT published to the internet yet.
// It only lives magically in a folder next to us (`../module_b`).
// Without a `go.work` file, `go build` will fail attempting to download module_b.

require github.com/enterprise/module_b v0.0.0
