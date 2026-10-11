package main

import (
	"regexp"
	"testing"

	"go-examples/internal/testutil"
)

// main prints one line per tick, each stamped with time.Now(), so the test
// matches a shape rather than exact text.
var tickLine = regexp.MustCompile(`^ticker at .+$`)

func TestMainFunc(t *testing.T) {
	lines := testutil.CaptureLines(t, main)

	if len(lines) == 0 {
		t.Fatal("main produced no output")
	}

	// The ticker runs at 200ms for a 1s sleep, so 5 ticks are scheduled but the
	// last one races with the sleep expiring. Timer granularity and machine load
	// move this around, so assert a range instead of an exact count.
	const minTicks, maxTicks = 3, 6

	var ticks int
	var sawCompleted bool
	for _, line := range lines {
		switch {
		case tickLine.MatchString(line):
			ticks++
		case line == "completed":
			sawCompleted = true
		case line == "work done":
			// Printed by the goroutine after it receives on done, racing with
			// main's own "completed". May land after main returns and therefore
			// after capture ends, so its presence is not asserted.
		default:
			t.Errorf("unexpected output line %q", line)
		}
	}

	if ticks < minTicks || ticks > maxTicks {
		t.Errorf("got %d ticker lines, want between %d and %d; full output: %q",
			ticks, minTicks, maxTicks, lines)
	}
	if !sawCompleted {
		t.Errorf(`missing "completed" line; full output: %q`, lines)
	}
}
