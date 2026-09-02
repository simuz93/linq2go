package linq2go

import (
	"cmp"
	"slices"
)

// Min returns the smallest value selected by fn, or 0 if the slice is empty, comparing
// with cmp.Compare, which sorts NaN before every number so a single NaN wins
//
//	FromSlice([]int{4, 1, 3}).Min(func(v int) int { return v }) // 1
func (s *slice[T]) Min[V Number](fn func(T) V) V {
	if len(s.values) == 0 {
		return 0
	}

	return fn(s.WhereMin(fn))
}

// WhereMin returns the element whose value selected by fn is the smallest, or the zero
// value if the slice is empty, comparing through slices.MinFunc with cmp.Compare, which
// sorts NaN before every number so an element whose value is NaN wins
//
//	FromSlice([]string{"go", "linq"}).WhereMin(func(s string) int { return len(s) }) // "go"
func (s *slice[T]) WhereMin[V Number](fn func(T) V) T {
	if len(s.values) == 0 {
		return *new(T)
	}

	compare := func(e1 T, e2 T) int {
		return cmp.Compare(fn(e1), fn(e2))
	}

	return slices.MinFunc(s.values, compare)
}

// Min returns, for each group, the smallest value selected by fn, comparing NaN as slice.Min does
//
//	FromSlice([]int{1, 2, 3, 4}).Group(func(v int) bool { return v%2 == 0 }).Min(func(v int) int { return v }).ToMap() // map[false:1 true:2]
func (g *group[K, V]) Min[NewV Number](fn func(V) NewV) *dictionary[K, NewV] {
	result := make(map[K]NewV, len(g.values))

	for k, v := range g.values {
		result[k] = newSlice(v).Min(fn)
	}

	return newDictionary(result)
}
