package linq2go

import (
	"maps"
	"math"
	"testing"
)

func Test_Max(t *testing.T) {
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
			name: "Max value",
			args: []testArg{
				{name: "three", value: 3},
				{name: "four", value: 4},
				{name: "one", value: 1},
			},
			wantIdx:   1,
			wantValue: 4,
		},
		{
			name: "First of equal maxima",
			args: []testArg{
				{name: "first", value: 4},
				{name: "second", value: 4},
			},
			wantIdx:   0,
			wantValue: 4,
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

			idx, value := FromSlice(tt.args).Max(fn)

			if idx != tt.wantIdx || value != tt.wantValue {
				t.Errorf("Max() = (%v, %v), want (%v, %v)", idx, value, tt.wantIdx, tt.wantValue)
			}
		})
	}
}

func Test_Max_NaNNeverWins(t *testing.T) {
	fn := func(v float64) float64 { return v }

	// cmp.Compare sorts NaN below every number, so a NaN never wins a maximum...
	idx, value := FromSlice([]float64{1, math.NaN(), 2}).Max(fn)
	if idx != 2 || value != 2 {
		t.Errorf("Max() = (%v, %v), want (2, 2)", idx, value)
	}

	// ...unless every value is one
	idx, value = FromSlice([]float64{math.NaN(), math.NaN()}).Max(fn)
	if idx != 0 || !math.IsNaN(value) {
		t.Errorf("Max() = (%v, %v), want (0, NaN)", idx, value)
	}
}

func Test_WhereMax(t *testing.T) {
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
			name: "Max element",
			args: []testArg{
				{name: "three", value: 3},
				{name: "four", value: 4},
				{name: "one", value: 1},
			},
			wantIdx:  1,
			wantElem: testArg{name: "four", value: 4},
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

			idx, elem := FromSlice(tt.args).WhereMax(fn)

			if idx != tt.wantIdx || elem != tt.wantElem {
				t.Errorf("WhereMax() = (%v, %+v), want (%v, %+v)", idx, elem, tt.wantIdx, tt.wantElem)
			}
		})
	}
}

func Test_Max_Group(t *testing.T) {
	groupBy := func(v int) bool { return v%2 == 0 }
	fn := func(v int) int { return v }

	result := FromSlice([]int{1, 2, 3, 4}).Group(groupBy).Max(fn).ToMap()

	want := map[bool]int{false: 3, true: 4}
	if !maps.Equal(result, want) {
		t.Errorf("Group().Max().ToMap() = %v, want %v", result, want)
	}
}
