package linq2go

import "slices"

// Concat returns a new slice with values appended after the elements of this one
func (s *slice[T]) Concat(values []T) *slice[T] {
	result := slices.Concat(s.s, values)

	// slices.Concat returns nil when both slices are empty
	if result == nil {
		result = []T{}
	}

	return newSlice(result)
}
