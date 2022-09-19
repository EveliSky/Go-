// fib_test.go
package fibonacci

import (
	"testing"
)

func TestFib(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{-1, 0},
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{6, 8},
		{10, 55},
	}

	for _, tt := range cases {
		got := fib(tt.n)
		if got != tt.want {
			t.Errorf("fib(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
