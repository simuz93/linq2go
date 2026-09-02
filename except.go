package linq2go

// Except returns a new slice with the elements that match no value in slice according to fn. Duplicates are kept
func (s *slice[T]) Except(slice []T, fn func(T, T) bool) *slice[T] {
	result := []T{}

	values := newSlice(slice)

	for _, elem := range s.values {
		// All the slices must satisfy the fn condition: this means that elem isn't contained in any of the other slices
		if !values.Contains(elem, fn) {
			result = append(result, elem)
		}
	}

	return newSlice(result)
}
