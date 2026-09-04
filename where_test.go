package linq2go

import (
	"maps"
	"slices"
	"testing"
)

func Test_Where_Slice(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "Where",
			args: []int{1, 2, 3, 4},
			want: []int{3, 4},
		},
		{
			name: "No match",
			args: []int{1, 2},
			want: []int{},
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
			fn := func(v int) bool { return v > 2 }

			result := FromSlice(tt.args).Where(fn).ToSlice()

			if result == nil {
				t.Fatal("Where().ToSlice() = nil, want an empty slice, never nil")
			}
			if !slices.Equal(result, tt.want) {
				t.Errorf("Where().ToSlice() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_Where_Dictionary(t *testing.T) {
	tests := []struct {
		name string
		args map[string]int
		want map[string]int
	}{
		{
			name: "Where",
			args: map[string]int{"a": 1, "b": 2, "c": 4},
			want: map[string]int{"b": 2, "c": 4},
		},
		{
			name: "No match",
			args: map[string]int{"a": 1},
			want: map[string]int{},
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
			fn := func(k string, v int) bool { return v%2 == 0 }

			result := FromMap(tt.args).Where(fn).ToMap()

			if result == nil {
				t.Fatal("Where().ToMap() = nil, want an empty map, never nil")
			}
			if !maps.Equal(result, tt.want) {
				t.Errorf("Where().ToMap() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_Where_Group(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want map[bool][]int
	}{
		{
			name: "Where",
			args: []int{1, 2, 3},
			want: map[bool][]int{true: {2, 3}},
		},
		{
			name: "No match",
			args: []int{1, 2},
			want: map[bool][]int{},
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
			groupBy := func(v int) bool { return v > 1 }
			fn := func(k bool, bucket []int) bool { return len(bucket) > 1 }

			result := FromSlice(tt.args).Group(groupBy).Where(fn).ToMap()

			if result == nil {
				t.Fatal("Where().ToMap() = nil, want an empty map, never nil")
			}
			if !maps.EqualFunc(result, tt.want, slices.Equal) {
				t.Errorf("Where().ToMap() = %v, want %v", result, tt.want)
			}
		})
	}
}
