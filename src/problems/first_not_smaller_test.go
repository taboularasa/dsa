package problems

import "testing"

func TestFirstNotSmaller(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		want   int
	}{
		{
			arr:    []int{1, 3, 3, 5, 8, 8, 10},
			target: 2,
			want:   1,
		},
		{
			arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target: 10,
			want:   9,
		},
	}

	for _, tc := range tests {
		got := FirstNotSmaller(tc.arr, tc.target)
		if got != tc.want {
			t.Errorf("wanted %d, got %d", tc.want, got)
		}
	}
}
