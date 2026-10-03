package main

import (
	"fmt"
	"regexp"
	"testing"

	"go-examples/internal/testutil"
)

// main spawns one worker per job.
const numJobs = 5

var (
	callingRe  = regexp.MustCompile(`^calling (superman\d+)$`)
	fightingRe = regexp.MustCompile(`^(superman\d+) is fighting against (bad-guy\d+)$`)
	wonRe      = regexp.MustCompile(`^(superman\d+) won, (bad-guy\d+) defeated$`)
	justiceRe  = regexp.MustCompile(`^Justice: (superman\d+) beated (bad-guy\d+)$`)
)

// record stores a job -> worker pairing, complaining if the job shows up twice
// at the same stage.
func record(t *testing.T, seen map[string]string, stage, job, worker string) {
	t.Helper()

	if prev, dup := seen[job]; dup {
		t.Errorf("%s: %s reported twice, by %s and %s", stage, job, prev, worker)
	}
	seen[job] = worker
}

func TestMainFunc(t *testing.T) {
	lines := testutil.CaptureLines(t, main)

	// Which worker picks up which job is the scheduler's choice, and the four
	// kinds of line interleave freely. So bucket them into job -> worker maps
	// and assert on the sets instead of on order.
	fighting := map[string]string{}
	won := map[string]string{}
	justice := map[string]string{}
	called := map[string]bool{}

	for _, line := range lines {
		switch {
		case callingRe.MatchString(line):
			m := callingRe.FindStringSubmatch(line)
			if called[m[1]] {
				t.Errorf("worker %s announced itself twice", m[1])
			}
			called[m[1]] = true
		case fightingRe.MatchString(line):
			m := fightingRe.FindStringSubmatch(line)
			record(t, fighting, "fighting", m[2], m[1])
		case wonRe.MatchString(line):
			m := wonRe.FindStringSubmatch(line)
			record(t, won, "won", m[2], m[1])
		case justiceRe.MatchString(line):
			m := justiceRe.FindStringSubmatch(line)
			record(t, justice, "Justice", m[2], m[1])
		default:
			t.Errorf("unexpected output line %q", line)
		}
	}

	// Every job is fought, won, and credited exactly once. This is guaranteed
	// rather than timing-dependent: a worker prints both of its lines before
	// sending its result, and main reads all numJobs results before returning.
	for _, stage := range []struct {
		name string
		seen map[string]string
	}{
		{"fighting", fighting},
		{"won", won},
		{"Justice", justice},
	} {
		if len(stage.seen) != numJobs {
			t.Errorf("%s: got %d jobs %v, want %d", stage.name, len(stage.seen), stage.seen, numJobs)
		}
		for i := range numJobs {
			job := fmt.Sprintf("bad-guy%d", i)
			if _, ok := stage.seen[job]; !ok {
				t.Errorf("%s: no line for %s", stage.name, job)
			}
		}
	}

	// A job must be carried by one worker from start to credit.
	for job, worker := range fighting {
		if won[job] != worker {
			t.Errorf("%s: fought by %s but won by %s", job, worker, won[job])
		}
		if justice[job] != worker {
			t.Errorf("%s: fought by %s but credited to %s", job, worker, justice[job])
		}
	}

	valid := map[string]bool{}
	for i := range numJobs {
		valid[fmt.Sprintf("superman%d", i)] = true
	}
	for worker := range called {
		if !valid[worker] {
			t.Errorf("unknown worker %q announced itself", worker)
		}
	}
	for job, worker := range fighting {
		if !valid[worker] {
			t.Errorf("%s handled by unknown worker %q", job, worker)
		}
	}

	// "calling" is the one line NOT guaranteed numJobs times: a fast worker can
	// take several jobs, leaving another still unscheduled when main returns and
	// capture ends. At least one must have run to accept a job.
	if len(called) < 1 || len(called) > numJobs {
		t.Errorf("got %d workers announcing themselves %v, want between 1 and %d",
			len(called), called, numJobs)
	}
}
