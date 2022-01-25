package sort

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr = Merge(arr)
	if !reflect.DeepEqual(arr, expected) {
		t.Fatalf("array not sorted %v", arr)
	}
}
