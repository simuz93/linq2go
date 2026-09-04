package linq2go

import (
	"slices"
	"testing"
)

func Test_Skip(t *testing.T) {
	tests := []struct {
		name string
		args []int
		n    int
		want []int
	}{
		{
			name: "Skip",
			args: []int{1, 2, 3},
			n:    1,
			want: []int{2, 3},
		},
		{
			name: "Zero",
			args: []int{1, 2, 3},
			n:    0,
			want: []int{1, 2, 3},
		},
		{
			name: "Negative",
			args: []int{1, 2, 3},
			n:    -1,
			want: []int{1, 2, 3},
		},
		{
			name: "Beyond length",
			args: []int{1, 2, 3},
			n:    10,
			want: []int{},
		},
		{
			name: "Empty",
			args: []int{},
			n:    2,
			want: []int{},
		},
		{
			name: "Nil",
			args: nil,
			n:    2,
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromSlice(tt.args).Skip(tt.n).ToSlice()

			if result == nil {
				t.Fatal("Skip().ToSlice() = nil, want an empty slice, never nil")
			}
			if !slices.Equal(result, tt.want) {
				t.Errorf("Skip().ToSlice() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_SkipWhile(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "Keeps from the first that does not match",
			args: []int{1, 2, 3, 1},
			want: []int{3, 1},
		},
		{
			name: "All match",
			args: []int{1, 2},
			want: []int{},
		},
		{
			name: "First does not match",
			args: []int{3, 1, 2},
			want: []int{3, 1, 2},
		},
		{
			name: "Empty",
			args: []int{},
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
			fn := func(v int) bool { return v < 3 }

			result := FromSlice(tt.args).SkipWhile(fn).ToSlice()

			if result == nil {
				t.Fatal("SkipWhile().ToSlice() = nil, want an empty slice, never nil")
			}
			if !slices.Equal(result, tt.want) {
				t.Errorf("SkipWhile().ToSlice() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_SkipWhile_CapsTheCapacity(t *testing.T) {
	base := FromSlice([]int{1, 2, 3})

	skipped := base.SkipWhile(func(v int) bool { return v < 4 }).ToSlice()
	_ = append(skipped, 42)

	// the capped capacity forces the append to reallocate instead of writing past base's length
	if got := base.ToSlice(); !slices.Equal(got, []int{1, 2, 3}) {
		t.Errorf("appending to the result wrote into the shared array: got %v, want [1 2 3]", got)
	}
}

func Test_Skip_CapsTheCapacity(t *testing.T) {
	base := FromSlice([]int{1, 2, 3})

	skipped := base.Skip(3).ToSlice()
	_ = append(skipped, 42)

	// the capped capacity forces the append to reallocate instead of writing past base's length
	if got := base.ToSlice(); !slices.Equal(got, []int{1, 2, 3}) {
		t.Errorf("appending to the result wrote into the shared array: got %v, want [1 2 3]", got)
	}
}
