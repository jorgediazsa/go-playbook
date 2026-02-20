package generate

// Context: Code Generation (go:generate)
// You want to embed the build timestamp into the compiled binary so that
// running `my-app --version` outputs exactly when it was compiled.
//
// Why this matters: Instead of using complex Makefile scripts, Go provides
// a built-in macro system. `go generate ./...` scans your source files for
// specifically formatted comments and executes the commands.
//
// Requirements:
// 1. Add a `//go:generate` directive exactly where designated.
// 2. The directive must execute the `tool/generator.go` file using `go run`.
// 3. The generator creates a file called `generated_version.go` defining a
//    constant `BuildTimestamp`.
// 4. Ensure `TestGeneration` passes!

// BUG: Missing the generation directive!
// TODO: Write the go:generate comment below. (e.g., //go:generate go run tool/generator.go)

var VersionInfo = "Dev"

func GetVersion() string {
	// After generation, BuildTimestamp will exist.
	// We use it here to combine VersionInfo and BuildTimestamp.
	// You will get compilation errors until you actually run `go generate .`!
	return VersionInfo + " - " + BuildTimestamp
}
