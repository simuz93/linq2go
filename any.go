package linq2go

import "slices"

/*
Any returns true if any of the elements in the slice satisfy the fn condition.
If the slice has no elements, Any returns false.
Wraps slices.ContainsFunc
*/
func (s *slice[T]) Any(fn func(T) bool) bool {
	return slices.ContainsFunc(s.s, fn)
}
