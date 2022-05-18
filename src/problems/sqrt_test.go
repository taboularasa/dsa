package problems

import "testing"

func TestSqrt(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 4, want: 2},
		{n: 8, want: 2},
		{n: 10, want: 3},
	}

	for _, tc := range tests {
		got := Sqrt(tc.n)
		if got != tc.want {
			t.Errorf("wanted %d, got %d", tc.want, got)
		}
	}
}
