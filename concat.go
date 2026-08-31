package linq2go

import "slices"

// Concat returns a new slice with values appended after the elements of this one
func (s *slice[T]) Concat(values []T) *slice[T] {
	return newSlice(slices.Concat(s.s, values))
}
