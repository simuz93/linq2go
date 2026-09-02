package linq2go

import (
	"cmp"
	"slices"
)

// Max returns the largest value selected by fn, or 0 if the slice is empty. Values are
// compared with cmp.Compare, which sorts NaN before every number, so a NaN never wins
// unless every value is one
func (s *slice[T]) Max[K Number](fn func(T) K) K {
	if len(s.values) == 0 {
		return 0
	}

	return fn(s.WhereMax(fn))
}

// WhereMax returns the element whose value selected by fn is the largest, or the zero
// value if the slice is empty. Values are compared with cmp.Compare through
// slices.MaxFunc, which sorts NaN before every number, so a NaN wins only when every
// value is one
func (s *slice[T]) WhereMax[K Number](fn func(T) K) T {
	if len(s.values) == 0 {
		return *new(T)
	}

	compare := func(e1 T, e2 T) int {
		return cmp.Compare(fn(e1), fn(e2))
	}

	return slices.MaxFunc(s.values, compare)
}

// Max returns, for each group, the largest value selected by fn. See slice.Max for how NaN compares
func (g *group[K, V]) Max[T Number](fn func(V) T) *dictionary[K, T] {
	result := make(map[K]T, len(g.values))

	for k, v := range g.values {
		result[k] = newSlice(v).Max(fn)
	}

	return newDictionary(result)
}
