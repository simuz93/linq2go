package linq2go

import (
	"cmp"
	"slices"
)

// Max returns the largest value selected by fn, or 0 if the slice is empty, comparing
// with cmp.Compare, which sorts NaN before every number so a NaN never wins unless every
// value is one
//
//	FromSlice([]int{1, 5, 3}).Max(func(v int) int { return v }) // 5
func (s *slice[T]) Max[V Number](fn func(T) V) V {
	if len(s.values) == 0 {
		return 0
	}

	return fn(s.WhereMax(fn))
}

// WhereMax returns the element whose value selected by fn is the largest, or the zero
// value if the slice is empty, comparing through slices.MaxFunc with cmp.Compare, which
// sorts NaN before every number so a NaN wins only when every value is one
//
//	FromSlice([]string{"go", "linq"}).WhereMax(func(s string) int { return len(s) }) // "linq"
func (s *slice[T]) WhereMax[V Number](fn func(T) V) T {
	if len(s.values) == 0 {
		return *new(T)
	}

	compare := func(e1 T, e2 T) int {
		return cmp.Compare(fn(e1), fn(e2))
	}

	return slices.MaxFunc(s.values, compare)
}

// Max returns, for each group, the largest value selected by fn, comparing NaN as slice.Max does
//
//	FromSlice([]int{1, 2, 3, 4}).Group(func(v int) bool { return v%2 == 0 }).Max(func(v int) int { return v }).ToMap() // map[false:3 true:4]
func (g *group[K, V]) Max[NewV Number](fn func(V) NewV) *dictionary[K, NewV] {
	result := make(map[K]NewV, len(g.values))

	for k, v := range g.values {
		result[k] = newSlice(v).Max(fn)
	}

	return newDictionary(result)
}
