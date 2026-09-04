package linq2go

import (
	"slices"
	"testing"
)

func Test_Reverse(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "Reverse",
			args: []int{1, 2, 3},
			want: []int{3, 2, 1},
		},
		{
			name: "Single element",
			args: []int{1},
			want: []int{1},
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
			result := FromSlice(tt.args).Reverse().ToSlice()

			if result == nil {
				t.Fatal("Reverse().ToSlice() = nil, want an empty slice, never nil")
			}
			if !slices.Equal(result, tt.want) {
				t.Errorf("Reverse().ToSlice() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_Reverse_DoesNotMutateTheReceiver(t *testing.T) {
	base := FromSlice([]int{1, 2, 3})

	base.Reverse()

	if got := base.ToSlice(); !slices.Equal(got, []int{1, 2, 3}) {
		t.Errorf("Reverse() reordered the receiver: got %v, want [1 2 3]", got)
	}
}
