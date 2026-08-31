// Package testutil holds small helpers shared by the tests of every exercise
// in this repo.
package testutil

import (
	"io"
	"os"
	"strings"
	"testing"
)

// CaptureStdout redirects os.Stdout for the duration of fn and returns
// everything written to it. The examples in this repo print their results with
// fmt.Println, so this is how their tests observe them:
//
//	out := testutil.CaptureStdout(t, main)
//
// fn runs on the calling goroutine, but output written by goroutines fn starts
// is only captured if they write before fn returns.
func CaptureStdout(tb testing.TB, fn func()) string {
	tb.Helper()

	r, w, err := os.Pipe()
	if err != nil {
		tb.Fatalf("os.Pipe: %v", err)
	}

	orig := os.Stdout
	os.Stdout = w

	// Drain the pipe concurrently so fn never blocks on a full pipe buffer.
	done := make(chan string, 1)
	go func() {
		var sb strings.Builder
		io.Copy(&sb, r)
		done <- sb.String()
	}()

	defer func() {
		os.Stdout = orig
		r.Close()
	}()

	func() {
		// Close the write end even if fn panics, so the reader finishes.
		defer w.Close()
		fn()
	}()

	return <-done
}

// CaptureLines is CaptureStdout split into trimmed, non-empty lines, which is
// what most exercise tests actually want to assert against.
func CaptureLines(tb testing.TB, fn func()) []string {
	tb.Helper()

	var lines []string
	for _, line := range strings.Split(CaptureStdout(tb, fn), "\n") {
		if line = strings.TrimSpace(line); line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}
