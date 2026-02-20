package functions

import (
	"errors"
	"strings"
	"testing"
)

func TestExecuteTx(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		err := ExecuteTx(func() error { return nil })
		if err != nil {
			t.Fatalf("Expected nil, got %v", err)
		}
	})

	t.Run("Error Wrapping", func(t *testing.T) {
		cause := errors.New("db timeout")
		err := ExecuteTx(func() error { return cause })

		if err == nil {
			t.Fatal("Expected an error, got nil")
		}

		expected := "tx failed: db timeout"
		if err.Error() != expected {
			t.Errorf("Expected %q, got %q", expected, err.Error())
		}

		if !errors.Is(err, cause) {
			t.Errorf("The wrapped error must use %%w so that errors.Is works")
		}
	})

	t.Run("Panic Recovery", func(t *testing.T) {
		err := ExecuteTx(func() error {
			panic("segfault")
		})

		if err == nil {
			t.Fatal("Expected recovered panic to be returned as an error, got nil")
		}

		if !strings.Contains(err.Error(), "tx failed: panic: segfault") {
			t.Errorf("Unexpected panic error format: %v", err)
		}
	})
}
