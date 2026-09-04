package linq2go

import (
	"slices"
	"testing"
)

func Test_OrderBy(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "Sorts ascending",
			args: []int{3, 1, 2},
			want: []int{1, 2, 3},
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
			fn := func(v int) int { return v }

			result := FromSlice(tt.args).OrderBy(fn).ToSlice()

			if result == nil {
				t.Fatal("OrderBy().ToSlice() = nil, want an empty slice, never nil")
			}
			if !slices.Equal(result, tt.want) {
				t.Errorf("OrderBy().ToSlice() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_OrderBy_IsStable(t *testing.T) {
	// elements with equal keys keep their relative order
	type pair struct{ key, ord int }
	args := []pair{{2, 0}, {1, 1}, {2, 2}, {1, 3}}

	result := FromSlice(args).OrderBy(func(p pair) int { return p.key }).ToSlice()

	want := []pair{{1, 1}, {1, 3}, {2, 0}, {2, 2}}
	if !slices.Equal(result, want) {
		t.Errorf("OrderBy().ToSlice() = %v, want %v", result, want)
	}
}

func Test_OrderBy_DoesNotMutateTheReceiver(t *testing.T) {
	base := FromSlice([]int{3, 1, 2})

	base.OrderBy(func(v int) int { return v })

	if got := base.ToSlice(); !slices.Equal(got, []int{3, 1, 2}) {
		t.Errorf("OrderBy() reordered the receiver: got %v, want [3 1 2]", got)
	}
}

func Test_OrderByDescending(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "Sorts descending",
			args: []int{3, 1, 2},
			want: []int{3, 2, 1},
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
			fn := func(v int) int { return v }

			result := FromSlice(tt.args).OrderByDescending(fn).ToSlice()

			if result == nil {
				t.Fatal("OrderByDescending().ToSlice() = nil, want an empty slice, never nil")
			}
			if !slices.Equal(result, tt.want) {
				t.Errorf("OrderByDescending().ToSlice() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_OrderByDescending_IsStable(t *testing.T) {
	// elements with equal keys keep their relative order
	type pair struct{ key, ord int }
	args := []pair{{2, 0}, {1, 1}, {2, 2}, {1, 3}}

	result := FromSlice(args).OrderByDescending(func(p pair) int { return p.key }).ToSlice()

	want := []pair{{2, 0}, {2, 2}, {1, 1}, {1, 3}}
	if !slices.Equal(result, want) {
		t.Errorf("OrderByDescending().ToSlice() = %v, want %v", result, want)
	}
}

func Test_OrderByDescending_DoesNotMutateTheReceiver(t *testing.T) {
	base := FromSlice([]int{3, 1, 2})

	base.OrderByDescending(func(v int) int { return v })

	if got := base.ToSlice(); !slices.Equal(got, []int{3, 1, 2}) {
		t.Errorf("OrderByDescending() reordered the receiver: got %v, want [3 1 2]", got)
	}
}
