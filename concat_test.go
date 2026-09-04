package linq2go

import (
	"slices"
	"testing"
)

func Test_Concat(t *testing.T) {
	tests := []struct {
		name   string
		args   []int
		values []int
		want   []int
	}{
		{
			name:   "Concat",
			args:   []int{1, 2},
			values: []int{3},
			want:   []int{1, 2, 3},
		},
		{
			name:   "Empty receiver",
			args:   []int{},
			values: []int{1, 2},
			want:   []int{1, 2},
		},
		{
			name:   "Nil values",
			args:   []int{1, 2},
			values: nil,
			want:   []int{1, 2},
		},
		{
			name:   "Both empty",
			args:   []int{},
			values: []int{},
			want:   []int{},
		},
		{
			name:   "Both nil",
			args:   nil,
			values: nil,
			want:   []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromSlice(tt.args).Concat(tt.values).ToSlice()

			if result == nil {
				t.Fatal("Concat().ToSlice() = nil, want an empty slice, never nil")
			}
			if !slices.Equal(result, tt.want) {
				t.Errorf("Concat().ToSlice() = %v, want %v", result, tt.want)
			}
		})
	}
}
