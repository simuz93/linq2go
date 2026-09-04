package linq2go

import (
	"slices"
	"testing"
)

func Test_FirstOrNil(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want *int
	}{
		{
			name: "Found",
			args: []int{1, 2, 3, 4},
			want: func() *int { v := 3; return &v }(),
		},
		{
			name: "Not found",
			args: []int{1, 1, 2, 2},
			want: nil,
		},
		{
			name: "Empty",
			args: []int{},
			want: nil,
		},
		{
			name: "Nil",
			args: nil,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := func(v int) bool { return v > 2 }

			result := FromSlice(tt.args).FirstOrNil(fn)

			switch {
			case tt.want == nil && result != nil:
				t.Errorf("FirstOrNil() = %v, want nil", *result)
			case tt.want != nil && result == nil:
				t.Errorf("FirstOrNil() = nil, want %v", *tt.want)
			case tt.want != nil && *result != *tt.want:
				t.Errorf("FirstOrNil() = %v, want %v", *result, *tt.want)
			}
		})
	}
}

func Test_FirstOrNil_ReturnsAPointerToACopy(t *testing.T) {
	base := FromSlice([]int{1, 2, 3})

	result := base.FirstOrNil(func(v int) bool { return v > 1 })
	*result = 42

	if got := base.ToSlice(); !slices.Equal(got, []int{1, 2, 3}) {
		t.Errorf("writing through the pointer changed the pipeline: got %v, want [1 2 3]", got)
	}
}

func Test_FirstOrDefault(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "Found",
			args: []int{1, 2, 3, 4},
			want: 3,
		},
		{
			name: "Not found",
			args: []int{1, 1, 2, 2},
			want: 0,
		},
		{
			name: "Empty",
			args: []int{},
			want: 0,
		},
		{
			name: "Nil",
			args: nil,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := func(v int) bool { return v > 2 }

			result := FromSlice(tt.args).FirstOrDefault(fn)

			if result != tt.want {
				t.Errorf("FirstOrDefault() = %v, want %v", result, tt.want)
			}
		})
	}
}
