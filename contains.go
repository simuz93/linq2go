package linq2go

import "slices"

// Contains reports whether any element matches value according to fn, which is called as fn(value, element)
func (s *slice[T]) Contains(value T, fn func(T, T) bool) bool {
	return slices.ContainsFunc(s.values, func(elem T) bool { return fn(value, elem) })
}
