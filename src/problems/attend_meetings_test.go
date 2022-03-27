package problems

import (
	"reflect"
	"testing"
)

func TestAttendMeetings(t *testing.T) {
	tests := []struct {
		in   [][]int
		want int
	}{
		{
			in:   [][]int{{1, 5}, {5, 8}, {10, 15}},
			want: 1,
		},
		{
			in:   [][]int{{1, 5}},
			want: 1,
		},
		{
			in:   [][]int{{1, 5}, {5, 7}},
			want: 1,
		},
		{
			in:   [][]int{{1, 5}, {4, 8}},
			want: 0,
		},
		{
			in:   [][]int{{13, 56}, {1, 3}, {4, 5}, {9990, 10341}, {8, 10}, {67, 9990}},
			want: 1,
		},
	}

	for _, tc := range tests {
		res := AttendAllMeetings(tc.in)
		if !reflect.DeepEqual(res, tc.want) {
			t.Errorf("got %v, want %v", res, tc.want)
		}
	}
}
