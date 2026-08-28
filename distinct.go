package linq2go

// Distinct use the selection function to create a new slice without duplicates. The first occurrence of each distinct element will be returned
func (s *slice[T]) Distinct[K comparable](fn func(T) K) *slice[T] {
	distinct := make(map[K]any)
	result := []T{}

	for _, v := range s.s {
		item := fn(v)
		if _, ok := distinct[item]; !ok {
			distinct[item] = nil
			result = append(result, v)
		}
	}

	return FromSlice(result)
}
