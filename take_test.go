package linq2go

import (
	"slices"
	"testing"
)

func Test_Take(t *testing.T) {
	tests := []struct {
		name string
		args []int
		n    int
		want []int
	}{
		{
			name: "Take",
			args: []int{1, 2, 3},
			n:    2,
			want: []int{1, 2},
		},
		{
			name: "Zero",
			args: []int{1, 2, 3},
			n:    0,
			want: []int{},
		},
		{
			name: "Negative",
			args: []int{1, 2, 3},
			n:    -1,
			want: []int{},
		},
		{
			name: "Beyond length",
			args: []int{1, 2, 3},
			n:    10,
			want: []int{1, 2, 3},
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
			result := FromSlice(tt.args).Take(tt.n).ToSlice()

			if result == nil {
				t.Fatal("Take().ToSlice() = nil, want an empty slice, never nil")
			}
			if !slices.Equal(result, tt.want) {
				t.Errorf("Take().ToSlice() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_Take_CapsTheCapacity(t *testing.T) {
	base := FromSlice([]int{1, 2, 3})

	taken := base.Take(2).ToSlice()
	_ = append(taken, 42)

	// the capped capacity forces the append to reallocate instead of overwriting base[2]
	if got := base.ToSlice(); !slices.Equal(got, []int{1, 2, 3}) {
		t.Errorf("appending to the result wrote into the shared array: got %v, want [1 2 3]", got)
	}
}
