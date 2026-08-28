package linq2go

// ExceptBy returns a new slice with the elements included in the first slice given but not included in the others
func (s *slice[T]) Except(slice []T, fn func(T, T) bool) *slice[T] {
	result := []T{}

	values := FromSlice(slice)

	for _, elem := range s.s {
		// All the slices must satisfy the fn condition: this means that elem isn't contained in any of the other slices
		if !values.Contains(elem, fn) {
			result = append(result, elem)
		}
	}

	return FromSlice(result)
}
