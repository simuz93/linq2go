package linq2go

import (
	"slices"
	"testing"
)

func Test_Values(t *testing.T) {
	tests := []struct {
		name string
		args map[string]int
		want []int
	}{
		{
			name: "Values",
			args: map[string]int{"a": 1, "b": 2},
			want: []int{1, 2},
		},
		{
			name: "Empty",
			args: map[string]int{},
			want: []int{},
		},
		{
			name: "Nil",
			args: nil,
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromMap(tt.args).Values()

			if result == nil {
				t.Fatal("Values() = nil, want an empty slice, never nil")
			}

			// the order is unspecified by contract, so compare sorted
			slices.Sort(result)
			if !slices.Equal(result, tt.want) {
				t.Errorf("Values() = %v, want %v", result, tt.want)
			}
		})
	}
}
