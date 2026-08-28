package linq2go

// ExceptBy returns a new slice with the elements included in the first slice given but not included in the others
func (s *Slice[T]) Except(slice []T, fn func(T, T) bool) *Slice[T] {
	result := []T{}

	for _, elem := range s.s {
		// All the slices must satisfy the fn condition: this means that elem isn't contained in any of the other slices
		if !FromSlice(slice).Contains(elem, fn) {
			result = append(result, elem)
		}
	}

	return FromSlice(result)
}
