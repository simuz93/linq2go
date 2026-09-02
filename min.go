package linq2go

import (
	"cmp"
	"slices"
)

// Min returns the smallest value selected by fn, or 0 if the slice is empty. Values are
// compared with cmp.Compare, which sorts NaN before every number, so a single NaN wins
func (s *slice[T]) Min[K Number](fn func(T) K) K {
	if len(s.s) == 0 {
		return 0
	}

	return fn(s.WhereMin(fn))
}

// WhereMin returns the element whose value selected by fn is the smallest, or the zero
// value if the slice is empty. Values are compared with cmp.Compare through
// slices.MinFunc, which sorts NaN before every number, so an element whose value is NaN wins
func (s *slice[T]) WhereMin[K Number](fn func(T) K) T {
	if len(s.s) == 0 {
		return *new(T)
	}

	compare := func(e1 T, e2 T) int {
		return cmp.Compare(fn(e1), fn(e2))
	}

	return slices.MinFunc(s.s, compare)
}

// Min returns, for each group, the smallest value selected by fn. See slice.Min for how NaN compares
func (g *group[K, V]) Min[T Number](fn func(V) T) *dictionary[K, T] {
	result := make(map[K]T, len(g.m))

	for k, v := range g.m {
		result[k] = newSlice(v).Min(fn)
	}

	return newDictionary(result)
}
