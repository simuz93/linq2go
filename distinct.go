package linq2go

// Distinct returns a new slice without duplicates, keeping the first element for each key selected by fn
func (s *slice[T]) Distinct[K comparable](fn func(T) K) *slice[T] {
	distinct := make(map[K]struct{})
	result := []T{}

	for _, v := range s.values {
		item := fn(v)
		if _, ok := distinct[item]; !ok {
			distinct[item] = struct{}{}
			result = append(result, v)
		}
	}

	return newSlice(result)
}
