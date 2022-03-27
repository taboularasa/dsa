package problems

import (
	"reflect"
	"sort"
	"testing"
)

func TestFrequentElements(t *testing.T) {
	tests := []struct {
		in   []int
		k    int
		want []int
	}{
		{
			in:   []int{1, 2, 3, 2, 4, 3, 1},
			k:    2,
			want: []int{1, 2},
		},
		{
			in:   []int{4, 4, 3, 5, 5, 1},
			k:    2,
			want: []int{4, 5},
		},
	}

	for _, tc := range tests {
		got := FrequentElements(tc.in, tc.k)
		sort.Slice(got, func(i, j int) bool { return got[i] < got[j] })
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("want %v, got %v", tc.want, got)
		}
	}
}
