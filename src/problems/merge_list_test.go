package problems

import (
	"reflect"
	"testing"
)

func listFromArray(in []int) *LinkedListNode {
	var next *LinkedListNode
	for i := len(in) - 1; i >= 0; i-- {
		e := LinkedListNode{
			value: in[i],
			next:  next,
		}
		if i != len(in) {
			next = &e
		}
	}
	return next
}

func arrayFromList(in *LinkedListNode) []int {
	res := []int{}
	node := in
	for node != nil {
		res = append(res, node.value)
		node = node.next
	}
	return res
}
func TestMergeLists(t *testing.T) {
	tests := []struct {
		input [][]int
		want  []int
	}{
		{
			input: [][]int{{1, 3, 5}, {3, 4}, {7}},
			want:  []int{1, 3, 3, 4, 5, 7},
		},
		{
			input: [][]int{},
			want:  []int{},
		},
		{
			input: [][]int{{1, 3, 5, 7, 8}},
			want:  []int{1, 3, 5, 7, 8},
		},
		{
			input: [][]int{{}, {}, {}, {}, {}},
			want:  []int{},
		},
		{
			input: [][]int{{1, 13, 22, 23}, {2, 12, 14}, {3, 11, 15, 21}, {4, 10, 16}, {5, 9, 17, 20}, {6, 8, 18}, {7, 19}},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
		},
		{
			input: [][]int{{-10, 5, 9, 11, 13}, {-10, 10}, {}, {-11, 2, 6, 10, 12}, {}, {-100}, {100}},
			want:  []int{-100, -11, -10, -10, 2, 5, 6, 9, 10, 10, 11, 12, 13, 100},
		},
		{
			input: [][]int{{-100}},
			want:  []int{-100},
		},
		{
			input: [][]int{{100}},
			want:  []int{100},
		},
	}

	for _, tc := range tests {
		var input []*LinkedListNode
		for _, e := range tc.input {
			input = append(input, listFromArray(e))
		}
		res := MergeLists(input)
		resAry := arrayFromList(res)
		if !reflect.DeepEqual(resAry, tc.want) {
			t.Errorf("wanted %v, got %v", tc.want, resAry)
		}
	}
}
