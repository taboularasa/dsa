package problems

import (
	"reflect"
	"testing"
)

func TestFindIntersection(t *testing.T) {
	tests := []struct {
		arr1 []int
		arr2 []int
		arr3 []int
		want []int
	}{
		{
			arr1: []int{2, 5, 10},
			arr2: []int{2, 3, 4, 10},
			arr3: []int{2, 4, 10},
			want: []int{2, 10},
		},
		{
			arr1: []int{1, 2, 2, 2, 9},
			arr2: []int{1, 1, 2, 2},
			arr3: []int{1, 1, 1, 2, 2, 2},
			want: []int{1, 2, 2},
		},
	}
	for _, tc := range tests {
		got := FindIntersection(tc.arr1, tc.arr2, tc.arr3)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("wanted %v, got %v", tc.want, got)
		}
	}
}
