package linq2go

import (
	"maps"
	"math"
	"testing"
)

func Test_Avg(t *testing.T) {
	type testArg struct {
		name  string
		value int
	}

	tests := []struct {
		name string
		args []testArg
		want float64
	}{
		{
			name: "Avg value",
			args: []testArg{
				{name: "one", value: 1},
				{name: "two", value: 2},
				{name: "three", value: 3},
				{name: "four", value: 4},
			},
			want: 2.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := func(a testArg) int { return a.value }

			result := FromSlice(tt.args).Avg(fn)

			if result != tt.want {
				t.Errorf("Avg() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_Avg_EmptyIsNaN(t *testing.T) {
	tests := []struct {
		name string
		args []int
	}{
		{
			name: "Empty",
			args: []int{},
		},
		{
			name: "Nil",
			args: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := func(v int) int { return v }

			result := FromSlice(tt.args).Avg(fn)

			if !math.IsNaN(result) {
				t.Errorf("Avg() = %v, want NaN", result)
			}
		})
	}
}

func Test_Avg_Group(t *testing.T) {
	groupBy := func(v int) bool { return v%2 == 0 }
	fn := func(v int) int { return v }

	result := FromSlice([]int{1, 2, 3, 4}).Group(groupBy).Avg(fn).ToMap()

	want := map[bool]float64{false: 2, true: 3}
	if !maps.Equal(result, want) {
		t.Errorf("Group().Avg().ToMap() = %v, want %v", result, want)
	}
}
