package problems

import (
	"reflect"
	"testing"
)

func TestNChooseK(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want [][]int
	}{
		{
			n: 5,
			k: 2,
			want: [][]int{
				{1, 2},
				{1, 3},
				{1, 4},
				{1, 5},
				{2, 3},
				{2, 4},
				{2, 5},
				{3, 4},
				{3, 5},
				{4, 5},
			},
		},
	}

	for _, tc := range tests {
		got := NChooseK(tc.n, tc.k)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("wanted %v, got %v", tc.want, got)
		}
	}
}
