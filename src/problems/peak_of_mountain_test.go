package problems

import "testing"

func TestPeakOfMountain(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{arr: []int{0, 1, 2, 3, 2, 1, 0}, want: 3},
		{arr: []int{0, 10, 3, 2, 1, 0}, want: 1},
		{arr: []int{0, 10, 0}, want: 1},
		{arr: []int{0, 1, 2, 12, 22, 32, 42, 52, 62, 72, 82, 92, 102, 112, 122, 132, 133, 132, 111, 0}, want: 16},
	}

	for _, tc := range tests {
		got := PeakOfMountain(tc.arr)
		if got != tc.want {
			t.Errorf("wanted %d, got %d", tc.want, got)
		}
	}
}
