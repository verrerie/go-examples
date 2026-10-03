package main

import "testing"

func TestIntMin(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		a    int
		b    int
		want int
	}{
		{"basic", 10, 20, 10},
		{"equal", 10, 10, 10},
		{"negative", -1, -2, -2},
		{"0 case 1", 0, 1, 0},
		{"0 case 2", 0, -1, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intMin(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("intMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkInitMain(b *testing.B) {
	for b.Loop() {
		intMin(10304083450, 3048503)
	}
}
