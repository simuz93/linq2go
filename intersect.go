package linq2go

// Intersect returns a new slice with the intersection of the given ones
func (s *Slice[T]) Intersect(slice []T, fn func(T, T) bool) *Slice[T] {
	result := []T{}

	values := FromSlice(slice)

	for _, elem := range s.s {
		if values.Contains(elem, fn) {
			result = append(result, elem)
		}
	}

	return FromSlice(result)
}
