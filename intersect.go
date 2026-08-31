package linq2go

// Intersect returns a new slice with the intersection of the given ones
func (s *slice[T]) Intersect(slice []T, fn func(T, T) bool) *slice[T] {
	result := []T{}

	values := newSlice(slice)

	for _, elem := range s.s {
		if values.Contains(elem, fn) {
			result = append(result, elem)
		}
	}

	return newSlice(result)
}
