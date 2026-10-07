package linq2go

import (
	"maps"
	"testing"
)

func Test_Count_Slice(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "Values",
			args: []int{1, 2, 3},
			want: 3,
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
			result := FromSlice(tt.args).Count()

			if result != tt.want {
				t.Errorf("Count() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_Count_Dictionary(t *testing.T) {
	tests := []struct {
		name string
		args map[string]int
		want int
	}{
		{
			name: "Values",
			args: map[string]int{"a": 1, "b": 2},
			want: 2,
		},
		{
			name: "Empty",
			args: map[string]int{},
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
			result := FromMap(tt.args).Count()

			if result != tt.want {
				t.Errorf("Count() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_Count_Group(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want map[bool]int
	}{
		{
			name: "Values",
			args: []int{1, 2, 3, 4, 6},
			want: map[bool]int{false: 2, true: 3},
		},
		{
			name: "Empty",
			args: []int{},
			want: map[bool]int{},
		},
		{
			name: "Nil",
			args: nil,
			want: map[bool]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := func(v int) bool { return v%2 == 0 }

			result := FromSlice(tt.args).Group(fn).Count().ToMap()

			if result == nil {
				t.Fatal("Count() returned nil")
			}

			if !maps.Equal(result, tt.want) {
				t.Errorf("Count() = %v, want %v", result, tt.want)
			}
		})
	}
}
