package linq2go

import (
	"slices"
	"testing"
)

func Test_Except(t *testing.T) {
	tests := []struct {
		name   string
		args   []int
		values []int
		want   []int
	}{
		{
			name:   "Except",
			args:   []int{1, 2, 2, 3},
			values: []int{2},
			want:   []int{1, 3},
		},
		{
			name:   "Duplicates kept",
			args:   []int{1, 1, 3},
			values: []int{3},
			want:   []int{1, 1},
		},
		{
			name:   "Empty values",
			args:   []int{1, 2},
			values: []int{},
			want:   []int{1, 2},
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

			result := FromSlice(tt.args).Except(tt.values, fn).ToSlice()

			if result == nil {
				t.Fatal("Except().ToSlice() = nil, want an empty slice, never nil")
			}
			if !slices.Equal(result, tt.want) {
				t.Errorf("Except().ToSlice() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_Except_CallsFnWithTheElementFirst(t *testing.T) {
	// an asymmetric fn pins the documented argument order: fn(element, value)
	fn := func(element, value int) bool { return element == value*10 }

	result := FromSlice([]int{10, 2, 30}).Except([]int{1, 3}, fn).ToSlice()

	if want := []int{2}; !slices.Equal(result, want) {
		t.Errorf("Except().ToSlice() = %v, want %v", result, want)
	}
}
