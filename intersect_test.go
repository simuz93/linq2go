package linq2go

import (
	"slices"
	"testing"
)

func Test_Intersect(t *testing.T) {
	tests := []struct {
		name   string
		args   []int
		values []int
		want   []int
	}{
		{
			name:   "Intersect",
			args:   []int{1, 2, 2, 3},
			values: []int{2, 3},
			want:   []int{2, 2, 3},
		},
		{
			name:   "No match",
			args:   []int{1, 2},
			values: []int{3},
			want:   []int{},
		},
		{
			name:   "Empty values",
			args:   []int{1, 2},
			values: []int{},
			want:   []int{},
		},
		{
			name:   "Empty",
			args:   []int{},
			values: []int{1},
			want:   []int{},
		},
		{
			name:   "Nil",
			args:   nil,
			values: []int{1},
			want:   []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := func(a, b int) bool { return a == b }

			result := FromSlice(tt.args).Intersect(tt.values, fn).ToSlice()

			if result == nil {
				t.Fatal("Intersect().ToSlice() = nil, want an empty slice, never nil")
			}
			if !slices.Equal(result, tt.want) {
				t.Errorf("Intersect().ToSlice() = %v, want %v", result, tt.want)
			}
		})
	}
}
