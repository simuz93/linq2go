package linq2go

import (
	"slices"
	"testing"
)

func Test_Distinct(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "Distinct",
			args: []int{1, 1, 2, 2, 3, 4},
			want: []int{1, 2, 3, 4},
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

			result := FromSlice(tt.args).Distinct(fn).ToSlice()

			if result == nil {
				t.Fatal("Distinct().ToSlice() = nil, want an empty slice, never nil")
			}
			if !slices.Equal(result, tt.want) {
				t.Errorf("Distinct().ToSlice() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_Distinct_KeepsTheFirstPerKey(t *testing.T) {
	type testArg struct {
		name  string
		value int
	}

	args := []testArg{
		{name: "first one", value: 1},
		{name: "second one", value: 1},
		{name: "first two", value: 2},
		{name: "second two", value: 2},
	}
	fn := func(a testArg) int { return a.value }

	result := FromSlice(args).Distinct(fn).ToSlice()

	want := []testArg{
		{name: "first one", value: 1},
		{name: "first two", value: 2},
	}
	if !slices.Equal(result, want) {
		t.Errorf("Distinct().ToSlice() = %+v, want %+v", result, want)
	}
}
