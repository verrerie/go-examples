package main

import (
	"testing"

	"go-examples/internal/testutil"
)

func TestMainFunc(t *testing.T) {
	lines := testutil.CaptureLines(t, main)

	want := []string{
		"timer1 fired",
		"timer2 fired",
		"timer2 already stopped",
	}

	if len(lines) != len(want) {
		t.Fatalf("got %q, want %q", lines, want)
	}
	for i := range want {
		if lines[i] != want[i] {
			t.Errorf("line %d = %q, want %q", i, lines[i], want[i])
		}
	}
}
