package linq2go

import (
	"slices"
	"testing"
)

func Test_Self(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{
			name: "Value",
			args: 42,
			want: 42,
		},
		{
			name: "Zero value",
			args: 0,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := Self(tt.args); result != tt.want {
				t.Errorf("Self() = %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_Equal(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want bool
	}{
		{
			name: "Equal",
			a:    "a",
			b:    "a",
			want: true,
		},
		{
			name: "Different",
			a:    "a",
			b:    "b",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := Equal(tt.a, tt.b); result != tt.want {
				t.Errorf("Equal() = %v, want %v", result, tt.want)
			}
		})
	}
}

// what the two are for: the inference has to hold at every operator that takes one of the
// two fn shapes, so every such call site is compiled here
func Test_Self_AtTheSelectorOperators(t *testing.T) {
	if got := FromSlice([]int{1, 2, 2, 3}).Distinct(Self).ToSlice(); !slices.Equal(got, []int{1, 2, 3}) {
		t.Errorf("Distinct(Self) = %v, want [1 2 3]", got)
	}
	if got := FromSlice([]int{3, 1, 2}).OrderBy(Self).ToSlice(); !slices.Equal(got, []int{1, 2, 3}) {
		t.Errorf("OrderBy(Self) = %v, want [1 2 3]", got)
	}
	if got := FromSlice([]int{3, 1, 2}).OrderByDescending(Self).ToSlice(); !slices.Equal(got, []int{3, 2, 1}) {
		t.Errorf("OrderByDescending(Self) = %v, want [3 2 1]", got)
	}
	if got := FromSlice([]int{1, 2, 3}).Sum(Self); got != 6 {
		t.Errorf("Sum(Self) = %v, want 6", got)
	}
	if got := FromSlice([]int{1, 2, 6}).Avg(Self); got != 3 {
		t.Errorf("Avg(Self) = %v, want 3", got)
	}
	if idx, got := FromSlice([]int{9, 4, 6}).Min(Self); idx != 1 || got != 4 {
		t.Errorf("Min(Self) = %v, %v, want 1, 4", idx, got)
	}
	if idx, got := FromSlice([]int{1, 5, 3}).Max(Self); idx != 1 || got != 5 {
		t.Errorf("Max(Self) = %v, %v, want 1, 5", idx, got)
	}
	if got := FromSlice([]int{1, 1, 2}).Group(Self).Count(); got != 2 {
		t.Errorf("Group(Self).Count() = %v, want 2", got)
	}
}

func Test_Equal_AtTheComparisonOperators(t *testing.T) {
	if got := FromSlice([]int{1, 2, 3}).Contains(2, Equal); !got {
		t.Error("Contains(2, Equal) = false, want true")
	}
	if got := FromSlice([]int{1, 2, 2, 3}).Except([]int{2}, Equal).ToSlice(); !slices.Equal(got, []int{1, 3}) {
		t.Errorf("Except(_, Equal) = %v, want [1 3]", got)
	}
	if got := FromSlice([]int{1, 2, 2, 3}).Intersect([]int{2, 3}, Equal).ToSlice(); !slices.Equal(got, []int{2, 2, 3}) {
		t.Errorf("Intersect(_, Equal) = %v, want [2 2 3]", got)
	}
}
