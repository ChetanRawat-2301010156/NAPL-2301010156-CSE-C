
package main

import "testing"

func TestSafeDivide(t *testing.T) {
	got, err := safeDivide(10, 2)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 5 {
		t.Errorf("got %v, want 5", got)
	}

	_, err = safeDivide(10, 0)
	if err == nil {
		t.Error("expected division-by-zero error")
	}
}

func TestSafeSqrt(t *testing.T) {
	got, err := safeSqrt(25)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 5 {
		t.Errorf("got %v, want 5", got)
	}

	_, err = safeSqrt(-4)
	if err == nil {
		t.Error("expected error for negative input")
	}
}

func TestParsePositiveInt(t *testing.T) {
	got, err := parsePositiveInt("42")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 42 {
		t.Errorf("got %d, want 42", got)
	}

	// Empty input must fail.
	if _, err := parsePositiveInt(""); err == nil {
		t.Error("expected error for empty input")
	}

	// Non-numeric input must fail.
	if _, err := parsePositiveInt("abc"); err == nil {
		t.Error("expected error for invalid integer")
	}

	// Zero and negative numbers must fail.
	if _, err := parsePositiveInt("0"); err == nil {
		t.Error("expected error for zero")
	}
	if _, err := parsePositiveInt("-5"); err == nil {
		t.Error("expected error for negative input")
	}
}
