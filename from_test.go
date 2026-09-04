package linq2go

import (
	"iter"
	"maps"
	"slices"
	"testing"
)

func Test_FromSlice(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "Values",
			args: []int{1, 2, 3},
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
			result := FromSlice(tt.args).ToSlice()

			if result == nil {
				t.Fatal("FromSlice().ToSlice() = nil, want an empty slice, never nil")
			}
			if !slices.Equal(result, tt.want) {
				t.Errorf("FromSlice().ToSlice() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_FromSlice_CopiesTheInput(t *testing.T) {
	src := []int{1, 2, 3}
	wrapped := FromSlice(src)

	src[0] = 42

	if got := wrapped.ToSlice(); !slices.Equal(got, []int{1, 2, 3}) {
		t.Errorf("mutating the source changed the pipeline: got %v, want [1 2 3]", got)
	}
}

func Test_FromMap(t *testing.T) {
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
				t.Fatal("FromMap().ToMap() = nil, want an empty map, never nil")
			}
			if !maps.Equal(result, tt.want) {
				t.Errorf("FromMap().ToMap() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_FromMap_CopiesTheInput(t *testing.T) {
	src := map[string]int{"a": 1}
	wrapped := FromMap(src)

	src["a"] = 42

	if got := wrapped.ToMap(); !maps.Equal(got, map[string]int{"a": 1}) {
		t.Errorf("mutating the source changed the pipeline: got %v, want map[a:1]", got)
	}
}

func Test_FromSeq(t *testing.T) {
	tests := []struct {
		name string
		args iter.Seq[int]
		want []int
	}{
		{
			name: "Values",
			args: slices.Values([]int{1, 2, 3}),
			want: []int{1, 2, 3},
		},
		{
			name: "Empty",
			args: slices.Values([]int{}),
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
			result := FromSeq(tt.args).ToSlice()

			if result == nil {
				t.Fatal("FromSeq().ToSlice() = nil, want an empty slice, never nil")
			}
			if !slices.Equal(result, tt.want) {
				t.Errorf("FromSeq().ToSlice() = %v, want %v", result, tt.want)
			}
		})
	}
}
