package linq2go

import "slices"

// Any reports whether at least one element satisfies fn. An empty slice returns false
func (s *slice[T]) Any(fn func(T) bool) bool {
	return slices.ContainsFunc(s.s, fn)
}
