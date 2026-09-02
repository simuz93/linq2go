package linq2go

import "slices"

// Any reports whether at least one element satisfies fn; an empty slice returns false
//
//	FromSlice([]int{1, 2, 3}).Any(func(v int) bool { return v > 2 }) // true
func (s *slice[T]) Any(fn func(T) bool) bool {
	return slices.ContainsFunc(s.values, fn)
}
