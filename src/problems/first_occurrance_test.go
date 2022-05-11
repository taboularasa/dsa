package problems

import "testing"

func TestFirstOccurance(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		want   int
	}{
		{
			arr:    []int{1, 3, 3, 3, 3, 4, 6, 10, 10, 10, 100},
			target: 3,
			want:   1,
		},
	}

	for _, tc := range tests {
		got := FirstOccurance(tc.arr, tc.target)
		if got != tc.want {
			t.Errorf("wanted %d, got %d", tc.want, got)
		}
	}
}
