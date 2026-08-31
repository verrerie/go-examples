package testutil

import (
	"fmt"
	"os"
	"testing"
)

func TestCaptureStdout(t *testing.T) {
	orig := os.Stdout

	out := CaptureStdout(t, func() { fmt.Print("hello") })

	if out != "hello" {
		t.Errorf("CaptureStdout = %q, want %q", out, "hello")
	}
	if os.Stdout != orig {
		t.Error("os.Stdout was not restored")
	}
}

func TestCaptureLines(t *testing.T) {
	lines := CaptureLines(t, func() {
		fmt.Println("one")
		fmt.Println("")
		fmt.Println("two")
	})

	want := []string{"one", "two"}
	if len(lines) != len(want) {
		t.Fatalf("CaptureLines = %q, want %q", lines, want)
	}
	for i := range want {
		if lines[i] != want[i] {
			t.Errorf("line %d = %q, want %q", i, lines[i], want[i])
		}
	}
}

func TestCaptureStdoutRestoresOnPanic(t *testing.T) {
	orig := os.Stdout

	defer func() {
		if recover() == nil {
			t.Error("panic did not propagate")
		}
		if os.Stdout != orig {
			t.Error("os.Stdout was not restored after panic")
		}
	}()

	CaptureStdout(t, func() { panic("boom") })
}
