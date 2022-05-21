package problems

import "testing"

func TestMinMaxWeight(t *testing.T) {
	tests := []struct {
		weights []int
		d       int
		want    int
	}{
		{
			weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			d:       5,
			want:    15,
		},
		{
			weights: []int{3, 2, 2, 4, 1, 4},
			d:       3,
			want:    6,
		},
		{
			weights: []int{1, 2, 3, 1, 1},
			d:       4,
			want:    3,
		},
		{
			weights: []int{5, 5, 5, 5, 5, 5, 5, 5, 5},
			d:       5,
			want:    10,
		},
		{
			weights: []int{5, 7, 6, 6, 3, 4, 5, 2, 3, 3, 4, 1, 9, 2},
			d:       5,
			want:    12,
		},
		{
			weights: []int{499, 377, 288, 96, 297, 107},
			d:       6,
			want:    499,
		},
		{
			weights: []int{8, 66, 143, 55, 233, 76, 101},
			d:       1,
			want:    682,
		},
	}

	for _, tc := range tests {
		got := MinMaxWeight(tc.weights, tc.d)
		if tc.want != got {
			t.Errorf("wanted %d, got %d", tc.want, got)
		}
	}
}
