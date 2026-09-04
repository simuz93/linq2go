package linq2go

import (
	"fmt"
	"maps"
	"slices"
	"testing"
)

func Test_ToMap_Slice(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want map[string]int
	}{
		{
			name: "Values",
			args: []string{"a", "b"},
			want: map[string]int{"a": 0, "b": 1},
		},
		{
			name: "Empty",
			args: []string{},
			want: map[string]int{},
		},
		{
			name: "Nil",
			args: nil,
			want: map[string]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := func(elem string, index int) (string, int) { return elem, index }

			result := FromSlice(tt.args).ToMap(fn)

			if result == nil {
				t.Fatal("ToMap() = nil, want an empty map, never nil")
			}
			if !maps.Equal(result, tt.want) {
				t.Errorf("ToMap() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_ToMap_Slice_KeepsTheFirstOnKeyCollision(t *testing.T) {
	fn := func(elem string, index int) (int, string) { return len(elem), elem }

	result := FromSlice([]string{"aa", "b", "cc"}).ToMap(fn)

	want := map[int]string{2: "aa", 1: "b"}
	if !maps.Equal(result, want) {
		t.Errorf("ToMap() = %v, want %v", result, want)
	}
}

func Test_ToMap_Dictionary(t *testing.T) {
	tests := []struct {
		name string
		args map[string]int
		want map[string]int
	}{
		{
			name: "Values",
			args: map[string]int{"a": 1, "b": 2},
			want: map[string]int{"a": 1, "b": 2},
		},
		{
			name: "Empty",
			args: map[string]int{},
			want: map[string]int{},
		},
		{
			name: "Nil",
			args: nil,
			want: map[string]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromMap(tt.args).ToMap()

			if result == nil {
				t.Fatal("ToMap() = nil, want an empty map, never nil")
			}
			if !maps.Equal(result, tt.want) {
				t.Errorf("ToMap() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_ToSlice_Dictionary(t *testing.T) {
	tests := []struct {
		name string
		args map[string]int
		want []string
	}{
		{
			name: "Values",
			args: map[string]int{"a": 1, "b": 2},
			want: []string{"a1", "b2"},
		},
		{
			name: "Empty",
			args: map[string]int{},
			want: []string{},
		},
		{
			name: "Nil",
			args: nil,
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := func(k string, v int) string { return fmt.Sprintf("%s%d", k, v) }

			result := FromMap(tt.args).ToSlice(fn)

			if result == nil {
				t.Fatal("ToSlice() = nil, want an empty slice, never nil")
			}

			// the order is unspecified by contract, so compare sorted
			slices.Sort(result)
			if !slices.Equal(result, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_ToMap_Group(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want map[bool][]int
	}{
		{
			name: "Values",
			args: []int{1, 2, 3},
			want: map[bool][]int{false: {1}, true: {2, 3}},
		},
		{
			name: "Empty",
			args: []int{},
			want: map[bool][]int{},
		},
		{
			name: "Nil",
			args: nil,
			want: map[bool][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := func(v int) bool { return v > 1 }

			result := FromSlice(tt.args).Group(fn).ToMap()

			if result == nil {
				t.Fatal("ToMap() = nil, want an empty map, never nil")
			}
			if !maps.EqualFunc(result, tt.want, slices.Equal) {
				t.Errorf("ToMap() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_ToSlice_ReturnsTheInternalStorage(t *testing.T) {
	// the documented consequence of the copy contract: two wrappers derived
	// from the same From share one array, and To* exits hand it back as-is
	base := FromSlice([]int{1, 2, 3})

	base.Skip(1).ToSlice()[0] = 42

	if got := base.ToSlice(); !slices.Equal(got, []int{1, 42, 3}) {
		t.Errorf("writing through a derived ToSlice() was not visible through the base: got %v, want [1 42 3]", got)
	}
}
