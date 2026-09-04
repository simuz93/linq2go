package linq2go

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"testing"
)

func Test_Select_Slice(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []string
	}{
		{
			name: "Select string",
			args: []int{1, 2, 3, 4},
			want: []string{"1", "2", "3", "4"},
		},
		{
			name: "Empty",
			args: []int{},
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
			fn := func(v int) string { return strconv.Itoa(v) }

			result := FromSlice(tt.args).Select(fn).ToSlice()

			if result == nil {
				t.Fatal("Select().ToSlice() = nil, want an empty slice, never nil")
			}
			if !slices.Equal(result, tt.want) {
				t.Errorf("Select().ToSlice() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_SelectMany(t *testing.T) {
	type testArg struct {
		id     string
		values []int
	}

	tests := []struct {
		name string
		args []testArg
		want []string
	}{
		{
			name: "Select many",
			args: []testArg{
				{id: "A", values: []int{1, 2, 3}},
				{id: "B", values: []int{1, 2}},
				{id: "C", values: nil},
			},
			want: []string{"A1", "A2", "A3", "B1", "B2"},
		},
		{
			name: "Empty",
			args: []testArg{},
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
			fn := func(a testArg) []string {
				result := make([]string, 0, len(a.values))
				for _, v := range a.values {
					result = append(result, fmt.Sprintf("%s%d", a.id, v))
				}
				return result
			}

			result := FromSlice(tt.args).SelectMany(fn).ToSlice()

			if result == nil {
				t.Fatal("SelectMany().ToSlice() = nil, want an empty slice, never nil")
			}
			if !slices.Equal(result, tt.want) {
				t.Errorf("SelectMany().ToSlice() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_Select_Dictionary(t *testing.T) {
	tests := []struct {
		name string
		args map[string]int
		want map[string]int
	}{
		{
			name: "Select",
			args: map[string]int{"a": 1, "b": 2},
			want: map[string]int{"a": 10, "b": 20},
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
			fn := func(k string, v int) int { return v * 10 }

			result := FromMap(tt.args).Select(fn).ToMap()

			if result == nil {
				t.Fatal("Select().ToMap() = nil, want an empty map, never nil")
			}
			if !maps.Equal(result, tt.want) {
				t.Errorf("Select().ToMap() = %v, want %v", result, tt.want)
			}
		})
	}
}
