package problems

import (
	"reflect"
	"testing"
)

func TestLCM(t *testing.T) {
	tests := []struct {
		in   string
		want []string
	}{
		{
			in:   "a1z",
			want: []string{"A1Z", "A1z", "a1Z", "a1z"},
		},
	}

	for _, tc := range tests {
		got := LCP(tc.in)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("wanted %v, got %v", tc.want, got)
		}
	}
}
