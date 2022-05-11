package problems

import (
	"reflect"
	"testing"
)

func TestGetPermute(t *testing.T) {
	tests := []struct {
		in   []int
		want [][]int
	}{
		{
			in: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 2, 1},
				{3, 1, 2},
			},
		},
	}

	for _, tc := range tests {
		got := GetPermutations(tc.in)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("wanted %v, got %v", tc.want, got)
		}
	}
}
