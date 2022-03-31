package problems

import (
	"reflect"
	"testing"
)

func TestKthLargestStream(t *testing.T) {
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
		got := KthLargestStream(tc.k, tc.initialStream, tc.appendStream)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("wanted %v, got %v", tc.want, got)
		}
	}
}

func TestKthLargestArray(t *testing.T) {
	tests := []struct {
		numbers []int
		k       int
		want    int
	}{
		{
			numbers: []int{5, 1, 10, 3, 2},
			k:       2,
			want:    5,
		},
		{
			numbers: []int{4, 1, 2, 2, 3},
			k:       4,
			want:    2,
		},
	}

	for _, tc := range tests {
		got := KthLargestArray(tc.numbers, tc.k)
		if got != tc.want {
			t.Errorf("wanted %d, got %d", tc.want, got)
		}
	}
}
