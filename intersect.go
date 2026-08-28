package linq2go

// Intersect returns a new slice with the intersection of the given ones
func (s *Slice[T]) Intersect(slice []T, fn func(T, T) bool) *Slice[T] {
	result := []T{}

	for _, elem := range s.s {
		if FromSlice(slice).Contains(elem, fn) {
			result = append(result, elem)
		}
	}

	return FromSlice(result)
}
