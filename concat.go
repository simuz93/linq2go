package linq2go

import "slices"

// Concat concatenates the given slices. Shortcut for slices.Concat, left here for retrocompatibility
func (s *slice[T]) Concat(values []T) *slice[T] {
	return newSlice(slices.Concat(s.s, values))
}
