package main

import (
	"math/rand/v2"
	"slices"
	"testing"
)

// The sorted oracle catches wrong rank direction and mishandled duplicates.
func TestKthLargestAgainstSorting(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	inputs := [][]int{
		{7}, {3, 2, 3, 1, 2, 4, 5, 5, 6}, {5, 5, 5, 5},
		{-7, -1, -3, -1}, {1, 2, 3, 4, 5}, {5, 4, 3, 2, 1},
		{-maxInt - 1, maxInt, 0, maxInt},
	}
	rng := rand.New(rand.NewPCG(19, 42))
	for range 200 {
		nums := make([]int, 1+rng.IntN(32))
		for i := range nums {
			nums[i] = rng.IntN(21) - 10
		}
		inputs = append(inputs, nums)
	}
	for _, nums := range inputs {
		sorted := slices.Clone(nums)
		slices.Sort(sorted)
		for k := 1; k <= len(nums); k++ {
			want := sorted[len(nums)-k]
			if got := findKthLargest(slices.Clone(nums), k); got != want {
				t.Fatalf("quickselect(%v, %d) = %d; want %d", nums, k, got, want)
			}
			input := slices.Clone(nums)
			if got := findKthLargestHeap(input, k); got != want {
				t.Fatalf("heap(%v, %d) = %d; want %d", nums, k, got, want)
			}
			if !slices.Equal(input, nums) {
				t.Fatalf("heap changed input: got %v; want %v", input, nums)
			}
		}
	}
}

func TestKthLargestHeapRejectsInvalidRank(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
	}{{nil, 1}, {[]int{1, 2}, 0}, {[]int{1, 2}, -1}, {[]int{1, 2}, 3}} {
		func() {
			defer func() {
				if recover() == nil {
					t.Errorf("heap(%v, %d) did not panic", tc.nums, tc.k)
				}
			}()
			findKthLargestHeap(tc.nums, tc.k)
		}()
	}
}
