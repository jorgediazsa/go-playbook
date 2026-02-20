package generate

import (
	"os"
	"os/exec"
	"testing"
)

func TestGeneration(t *testing.T) {
	// 1. Ensure `generated_version.go` doesn't exist to simulate a fresh clone.
	_ = os.Remove("generated_version.go")

	// 2. Run `go generate` which should invoke the tool if the comment is correct.
	cmd := exec.Command("go", "generate", ".")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go generate failed (perhaps the `//go:generate` comment is missing?): %v\nOutput: %s", err, string(out))
	}

	// 3. Did it create the file?
	if _, err := os.Stat("generated_version.go"); os.IsNotExist(err) {
		t.Fatalf("FAILED: go generate succeeded, but `generated_version.go` was not created. Did you type the command correctly? (e.g., go run tool/generator.go)")
	}

	// 4. Does the function compile and work now?
	// We verify compilation implicitly because if `BuildTimestamp` isn't declared,
	// `go test` wouldn't even have compiled this package!
	// (However, since we deleted the file above, we have to verify it via a sub-cmd
	// because `go test` requires all symbols to be present at PARSE time).

	cmdTest := exec.Command("go", "build", ".")
	outTest, errTest := cmdTest.CombinedOutput()
	if errTest != nil {
		t.Fatalf("FAILED: go generate ran, but the package still doesn't compile: %v\n%s", errTest, string(outTest))
	}

	// Clean up after the test
	_ = os.Remove("generated_version.go")
}
