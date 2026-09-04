package linq2go

import (
	"maps"
	"math"
	"testing"
)

func Test_Min(t *testing.T) {
	type testArg struct {
		name  string
		value int
	}

	tests := []struct {
		name      string
		args      []testArg
		wantIdx   int
		wantValue int
	}{
		{
			name: "Min value",
			args: []testArg{
				{name: "three", value: 3},
				{name: "one", value: 1},
				{name: "four", value: 4},
			},
			wantIdx:   1,
			wantValue: 1,
		},
		{
			name: "First of equal minima",
			args: []testArg{
				{name: "first", value: 1},
				{name: "second", value: 1},
			},
			wantIdx:   0,
			wantValue: 1,
		},
		{
			name:      "Empty",
			args:      []testArg{},
			wantIdx:   -1,
			wantValue: 0,
		},
		{
			name:      "Nil",
			args:      nil,
			wantIdx:   -1,
			wantValue: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := func(a testArg) int { return a.value }

			idx, value := FromSlice(tt.args).Min(fn)

			if idx != tt.wantIdx || value != tt.wantValue {
				t.Errorf("Min() = (%v, %v), want (%v, %v)", idx, value, tt.wantIdx, tt.wantValue)
			}
		})
	}
}

func Test_Min_NaNWins(t *testing.T) {
	fn := func(v float64) float64 { return v }

	idx, value := FromSlice([]float64{1, math.NaN(), 2}).Min(fn)

	// cmp.Compare sorts NaN below every number, so a single NaN wins a minimum
	if idx != 1 || !math.IsNaN(value) {
		t.Errorf("Min() = (%v, %v), want (1, NaN)", idx, value)
	}
}

func Test_WhereMin(t *testing.T) {
	type testArg struct {
		name  string
		value int
	}

	tests := []struct {
		name     string
		args     []testArg
		wantIdx  int
		wantElem testArg
	}{
		{
			name: "Min element",
			args: []testArg{
				{name: "three", value: 3},
				{name: "one", value: 1},
				{name: "four", value: 4},
			},
			wantIdx:  1,
			wantElem: testArg{name: "one", value: 1},
		},
		{
			name:     "Empty",
			args:     []testArg{},
			wantIdx:  -1,
			wantElem: testArg{},
		},
		{
			name:     "Nil",
			args:     nil,
			wantIdx:  -1,
			wantElem: testArg{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := func(a testArg) int { return a.value }

			idx, elem := FromSlice(tt.args).WhereMin(fn)

			if idx != tt.wantIdx || elem != tt.wantElem {
				t.Errorf("WhereMin() = (%v, %+v), want (%v, %+v)", idx, elem, tt.wantIdx, tt.wantElem)
			}
		})
	}
}

func Test_Min_Group(t *testing.T) {
	groupBy := func(v int) bool { return v%2 == 0 }
	fn := func(v int) int { return v }

	result := FromSlice([]int{1, 2, 3, 4}).Group(groupBy).Min(fn).ToMap()

	want := map[bool]int{false: 1, true: 2}
	if !maps.Equal(result, want) {
		t.Errorf("Group().Min().ToMap() = %v, want %v", result, want)
	}
}
