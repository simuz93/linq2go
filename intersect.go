package linq2go

// Intersect returns a new slice with the elements that match a value in slice according to fn. Duplicates are kept
func (s *slice[T]) Intersect(slice []T, fn func(T, T) bool) *slice[T] {
	result := []T{}

	values := newSlice(slice)

	for _, elem := range s.values {
		if values.Contains(elem, fn) {
			result = append(result, elem)
		}
	}

	return newSlice(result)
}
