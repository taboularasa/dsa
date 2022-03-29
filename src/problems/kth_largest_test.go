package problems

import (
	"reflect"
	"testing"
)

func TestKthLargest(t *testing.T) {
	tests := []struct {
		k             int
		initialStream []int
		appendStream  []int
		want          []int
	}{
		{
			k:             2,
			initialStream: []int{4, 6},
			appendStream:  []int{5, 2, 20},
			want:          []int{5, 5, 6},
		},
		{
			k:             2,
			initialStream: []int{1, 1},
			appendStream:  []int{2, 3, 4, 5},
			want:          []int{1, 2, 3, 4},
		},
		{
			k:             2,
			initialStream: []int{2, 1},
			appendStream:  []int{1, 1, 1, 1},
			want:          []int{1, 1, 1, 1},
		},
		{
			k:             3,
			initialStream: []int{3, 2, 1},
			appendStream:  []int{4, 4, 4},
			want:          []int{2, 3, 4},
		},
		{
			k:             3,
			initialStream: []int{4, 5, 6, 7},
			appendStream:  []int{5, 6, 4},
			want:          []int{5, 6, 6},
		},
		{
			k:             2,
			initialStream: []int{1000000000},
			appendStream:  []int{100000000},
			want:          []int{100000000},
		},
		{
			k:             1,
			initialStream: []int{1000000000},
			appendStream:  []int{100000000},
			want:          []int{1000000000},
		},
	}

	for _, tc := range tests {
		got := KthLargest(tc.k, tc.initialStream, tc.appendStream)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("wanted %v, got %v", tc.want, got)
		}
	}
}
