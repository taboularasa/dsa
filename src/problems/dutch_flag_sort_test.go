package problems

import (
	"reflect"
	"testing"
)

func TestDutchFlagSort(t *testing.T) {
	tests := []struct {
		input []string
		want  []string
	}{
		{
			input: []string{"G", "B", "G", "G", "R", "B", "R", "G"},
			want:  []string{"R", "R", "G", "G", "G", "G", "B", "B"},
		},
		{
			input: []string{"R", "R", "G", "G", "B", "B"},
			want:  []string{"R", "R", "G", "G", "B", "B"},
		},
		{
			input: []string{"R"},
			want:  []string{"R"},
		},
		{
			input: []string{"G"},
			want:  []string{"G"},
		},
		{
			input: []string{"B"},
			want:  []string{"B"},
		},
		{
			input: []string{"R", "R"},
			want:  []string{"R", "R"},
		},
		{
			input: []string{"G", "R"},
			want:  []string{"R", "G"},
		},
		{
			input: []string{"B", "R"},
			want:  []string{"R", "B"},
		},
	}
	for _, tc := range tests {
		result := DutchFlagSort(tc.input)
		if !reflect.DeepEqual(result, tc.want) {
			t.Fatalf("expected %v, got %v", tc.want, result)
		}
	}
}
