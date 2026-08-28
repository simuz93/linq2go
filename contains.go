package linq2go

import "slices"

// Concat concatenates the given slices. Shortcut for slices.Concat, left here for retrocompatibility
func (s *slice[T]) Contains(value T, fn func(T, T) bool) bool {
	return slices.ContainsFunc(s.s, func(elem T) bool { return fn(value, elem) })
}
