package linq2go

// ExceptBy returns a new slice with the elements included in the first slice given but not included in the others
func (s *Slice[T]) Except(slice []T, fn func(T, T) bool) *Slice[T] {
	result := []T{}

	for _, v1 := range slice {
		// All the slices must satisfy the fn condition: this means that elem isn't contained in any of the other slices
		if !s.Contains(v1, fn) {
			result = append(result, v1)
		}
	}

	return FromSlice(result)
}
