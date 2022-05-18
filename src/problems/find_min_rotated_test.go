package problems

import "testing"

func TestFindMinRotated(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{in: []int{30, 40, 50, 10, 20}, want: 3},
		{in: []int{0, 1, 2, 3, 4, 5}, want: 0},
		{in: []int{0}, want: 0},
	}

	for _, tc := range tests {
		got := FindMinRotated(tc.in)
		if got != tc.want {
			t.Errorf("wanted %d, got %d", tc.want, got)
		}
	}
}
