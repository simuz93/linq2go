package linq2go

// Intersect returns a new slice with the intersection of the given ones
func (s *Slice[T]) IntersectBy(slice []T, fn func(T, T) bool) *Slice[T] {
	result := []T{}

	for _, v1 := range slice {
		if s.Contains(v1, fn) {
			result = append(result, v1)
		}
	}

	return FromSlice(result)
}
