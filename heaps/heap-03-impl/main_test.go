package main

import (
	"math/rand"
	"slices"
	"testing"
)

func checkHeap(t *testing.T, h *Heap, want []int) {
	t.Helper()
	if h.size != len(want) || len(h.minHeap) != h.size+1 {
		t.Fatalf("size=%d, slice length=%d, want %d items", h.size, len(h.minHeap), len(want))
	}
	for child := 2; child <= h.size; child++ {
		if h.minHeap[child/2] > h.minHeap[child] {
			t.Fatalf("heap order broken: %v", h.minHeap[1:])
		}
	}
	got := slices.Clone(h.minHeap[1:])
	slices.Sort(got)
	if !slices.Equal(got, want) {
		t.Fatalf("heap contains %v, want %v", got, want)
	}
	if len(want) > 0 && h.peek() != want[0] {
		t.Fatalf("peek=%d, want %d", h.peek(), want[0])
	}
}

func TestHeapDrain(t *testing.T) {
	cases := [][]int{nil, {1}, {1, 2, 3}, {7, 2, 1, 6, 4, 0, 5, 3}, {2, 2, 1}, {0, -1, -2}, {-1}}
	rng := rand.New(rand.NewSource(1))
	for n := 0; n < 30; n++ {
		for trial := 0; trial < 100; trial++ {
			cases = append(cases, rng.Perm(n))
		}
	}
	for _, input := range cases {
		h := NewHeap(0)
		want := []int{}
		for _, v := range input {
			h.add(v)
			want = append(want, v)
			slices.Sort(want)
			checkHeap(t, h, want)
		}
		for len(want) > 0 {
			if got := h.pop(); got != want[0] {
				t.Fatalf("input %v: pop=%d, want %d", input, got, want[0])
			}
			want = want[1:]
			checkHeap(t, h, want)
		}
	}
}

func TestHeapMixedOperations(t *testing.T) {
	h := NewHeap(1)
	h.add(1)
	h.add(2)
	if got := h.pop(); got != 1 {
		t.Fatal(got)
	}
	h.add(0)
	checkHeap(t, h, []int{0, 2})
	if h.pop() != 0 || h.pop() != 2 {
		t.Fatal("pop/add regression")
	}
	// Reuse the drained heap, comparing every operation with a sorted model.
	rng := rand.New(rand.NewSource(42))
	want := []int{}
	for step := 0; step < 2000; step++ {
		if len(want) == 0 || rng.Intn(2) == 0 {
			v := rng.Intn(21) - 10
			h.add(v)
			want = append(want, v)
			slices.Sort(want)
		} else {
			if got := h.pop(); got != want[0] {
				t.Fatalf("step %d: pop=%d, want %d", step, got, want[0])
			}
			want = want[1:]
		}
		checkHeap(t, h, want)
	}
}

func TestEmptyHeapPanics(t *testing.T) {
	for _, drained := range []bool{false, true} {
		for _, operation := range []string{"peek", "pop"} {
			t.Run(operation+map[bool]string{false: "/new", true: "/drained"}[drained], func(t *testing.T) {
				h := NewHeap(0)
				if drained {
					h.add(-1)
					h.pop()
				}
				defer func() {
					if recover() == nil {
						t.Fatal("expected panic on empty heap")
					}
					checkHeap(t, h, nil)
				}()
				if operation == "peek" {
					h.peek()
				} else {
					h.pop()
				}
			})
		}
	}
}
