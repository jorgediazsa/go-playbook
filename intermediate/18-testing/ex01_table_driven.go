package testingadv

import (
	"errors"
	"strings"
)

// Context: Table-Driven Tests
// You have an incredibly basic string parser.
//
// Why this matters: Instead of writing TestParseEmpty, TestParseMissingPrefix, etc.
// Idiomatic Go groups all scenarios into a struct array inside a single test.
//
// Requirements:
// 1. The code below is fine. Open `ex01_table_driven_test.go`.
// 2. You will implement table-driven tests for this function.

var ErrInvalidFormat = errors.New("invalid format")

func ParseMetric(metric string) (string, int, error) {
	if metric == "" {
		return "", 0, ErrInvalidFormat
	}

	parts := strings.Split(metric, ":")
	if len(parts) != 2 {
		return "", 0, ErrInvalidFormat
	}

	if parts[1] == "" {
		return parts[0], 0, nil
	}

	return parts[0], len(parts[1]), nil
}
