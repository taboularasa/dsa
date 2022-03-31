package problems

import (
	"reflect"
	"testing"
)

func TestOnlineMedian(t *testing.T) {
	tests := []struct {
		stream []int
		want   []int
	}{
		{
			stream: []int{3, 8, 5, 2},
			want:   []int{3, 5, 5, 4},
		},
	}

	for _, tc := range tests {
		got := OnlineMedian(tc.stream)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("wanted %v, got %v", tc.want, got)
		}
	}
}
