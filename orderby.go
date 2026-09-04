package linq2go

import (
	"cmp"
	"slices"
)

// OrderBy returns a new slice sorted ascending by the key selected by fn; the sort is
// stable — elements with equal keys keep their order — and NaN sorts before every
// number, as cmp.Compare does
//
//	FromSlice([]int{3, 1, 2}).OrderBy(func(v int) int { return v }).ToSlice() // [1 2 3]
func (s *slice[T]) OrderBy[K cmp.Ordered](fn func(T) K) *slice[T] {
	result := slices.Clone(s.values)
	slices.SortStableFunc(result, func(a, b T) int { return cmp.Compare(fn(a), fn(b)) })

	return newSlice(result)
}

// OrderByDescending returns a new slice sorted descending by the key selected by fn; the
// sort is stable — elements with equal keys keep their order — and NaN sorts after every
// number, as the reverse of cmp.Compare does
//
//	FromSlice([]int{3, 1, 2}).OrderByDescending(func(v int) int { return v }).ToSlice() // [3 2 1]
func (s *slice[T]) OrderByDescending[K cmp.Ordered](fn func(T) K) *slice[T] {
	result := slices.Clone(s.values)
	slices.SortStableFunc(result, func(a, b T) int { return cmp.Compare(fn(b), fn(a)) })

	return newSlice(result)
}
