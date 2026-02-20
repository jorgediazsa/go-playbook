package main

// Context: Go Workspaces (go work)
// You are simultaneously developing an application (`module_a`) and an early-stage
// internal library (`module_b`) that isn't published to the internet yet.
//
// Why this matters: Historically, you had to use `replace` directives in `go.mod`
// to point to local folders. Developers would often accidentally commit these `replace`
// lines, breaking the build pipeline for everyone else.
// `go.work` solves this by creating a purely local workspace file (which you `.gitignore`)
// that stitches multiple `go.mod` directories together.
//
// Requirements:
// 1. If you run `cd module_a && go build .`, it will fail because it can't download `module_b`.
// 2. Setup a workspace file (`go.work`) in the `ex02_workspace/` directory that
//    includes both `module_a` and `module_b`.
// 3. Once configured correctly, `module_a` will compile locally without any `replace` hacks!

import (
	"fmt"

	"github.com/enterprise/module_b/calc"
)

func main() {
	// If you setup `go.work` correctly, this compiles and prints "10 + 20 = 30".
	result := calc.Add(10, 20)
	fmt.Printf("10 + 20 = %d\n", result)
}
