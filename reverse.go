package linq2go

import "slices"

// Reverse returns a new slice with the elements in the opposite order; the receiver keeps its own,
// since slices.Reverse works in place and is given a fresh clone
//
//	FromSlice([]int{1, 2, 3}).Reverse().ToSlice() // [3 2 1]
func (s *slice[T]) Reverse() *slice[T] {
	result := slices.Clone(s.values)
	slices.Reverse(result)

	return newSlice(result)
}
