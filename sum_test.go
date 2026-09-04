package linq2go

import (
	"maps"
	"math"
	"testing"
)

func Test_Sum(t *testing.T) {
	type testArg struct {
		name  string
		value int
	}

	tests := []struct {
		name string
		args []testArg
		want int
	}{
		{
			name: "Sum value",
			args: []testArg{
				{name: "one", value: 1},
				{name: "two", value: 2},
				{name: "three", value: 3},
				{name: "four", value: 4},
			},
			want: 10,
		},
		{
			name: "Empty",
			args: []testArg{},
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
			fn := func(a testArg) int { return a.value }

			result := FromSlice(tt.args).Sum(fn)

			if result != tt.want {
				t.Errorf("Sum() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_Sum_AccumulatesInTheSelectedType(t *testing.T) {
	fn := func(v int8) int8 { return v }

	// the sum accumulates in V, so a narrow integer overflows silently: 100+100 wraps to -56
	result := FromSlice([]int8{100, 100}).Sum(fn)

	if result != -56 {
		t.Errorf("Sum() = %v, want -56", result)
	}
}

func Test_Sum_NaNPoisons(t *testing.T) {
	fn := func(v float64) float64 { return v }

	result := FromSlice([]float64{1, math.NaN(), 2}).Sum(fn)

	if !math.IsNaN(result) {
		t.Errorf("Sum() = %v, want NaN", result)
	}
}

func Test_Sum_Group(t *testing.T) {
	groupBy := func(v int) bool { return v%2 == 0 }
	fn := func(v int) int { return v }

	result := FromSlice([]int{1, 2, 3, 4}).Group(groupBy).Sum(fn).ToMap()

	want := map[bool]int{false: 4, true: 6}
	if !maps.Equal(result, want) {
		t.Errorf("Group().Sum().ToMap() = %v, want %v", result, want)
	}
}
