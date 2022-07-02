package problems

import (
	"testing"
)

func TestVisibleTreeNode(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{in: "5 4 3 x x 8 x x 6 x x", want: 3},
		{in: "-100 x -500 x -50 x x", want: 2},
		{in: "9 8 11 x x 20 x x 6 x x", want: 3},
		{in: "5 8 3 x x 8 x x 6 x x", want: 4},
	}

	for _, tc := range tests {
		rootIdx := 0
		root := buildTree(splitWords(tc.in), &rootIdx)
		got := VisibleTreeNode(root)
		if got != tc.want {
			t.Errorf("got %d, wanted %d", got, tc.want)
		}
	}
}
