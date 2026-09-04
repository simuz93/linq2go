package linq2go

import (
	"maps"
	"slices"
	"testing"
)

func Test_Group_Slice(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want map[bool][]int
	}{
		{
			name: "Group",
			args: []int{1, 2, 3, 4},
			want: map[bool][]int{false: {1, 3}, true: {2, 4}},
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
			fn := func(v int) bool { return v%2 == 0 }

			result := FromSlice(tt.args).Group(fn).ToMap()

			if result == nil {
				t.Fatal("Group().ToMap() = nil, want an empty map, never nil")
			}
			if !maps.EqualFunc(result, tt.want, slices.Equal) {
				t.Errorf("Group().ToMap() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_Group_Dictionary(t *testing.T) {
	tests := []struct {
		name string
		args map[string]int
		want map[bool][]int
	}{
		{
			name: "Group",
			args: map[string]int{"a": 1, "b": 2, "c": 3},
			want: map[bool][]int{false: {1, 3}, true: {2}},
		},
		{
			name: "Empty",
			args: map[string]int{},
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
			fn := func(k string, v int) bool { return v%2 == 0 }

			result := FromMap(tt.args).Group(fn).ToMap()

			if result == nil {
				t.Fatal("Group().ToMap() = nil, want an empty map, never nil")
			}

			// buckets built from a map iteration have no order, so compare them sorted
			for _, bucket := range result {
				slices.Sort(bucket)
			}
			if !maps.EqualFunc(result, tt.want, slices.Equal) {
				t.Errorf("Group().ToMap() = %v, want %v", result, tt.want)
			}
		})
	}
}
