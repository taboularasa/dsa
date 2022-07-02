package problems

import "testing"

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{in: "1 2 4 x 7 x x 5 x x 3 x 6 x x", want: true},
		{in: "1 2 4 x 7 x x 5 x x 3 x 6 8 x x x", want: false},
		{in: "1 2 3 x x 4 x 6 x x 5 x x", want: false},
		{in: "1 2 3 x x 4 x 6 x x 5 x 7 x x", want: true},
		{in: "1 2 3 x x 4 x x 5 6 x 7 x x x", want: false},
		{in: "1 2 3 7 x x x 4 x x 5 6 x x x", want: true},
		{in: "1 2 3 4 x x 4 x x 3 x x 2 x x", want: false},
	}

	for _, tc := range tests {
		rootIdx := 0
		root := buildTree(splitWords(tc.in), &rootIdx)
		got := IsBalanced(root)
		if got != tc.want {
			t.Errorf("got %t wanted %t", got, tc.want)
		}
	}
}
