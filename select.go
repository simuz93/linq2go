package linq2go

// Select creates a new slice by applying the given function to each element of the given slice
func (s *slice[T]) Select[K any](fn func(T) K) *slice[K] {
	result := []K{}

	for _, v := range s.s {
		result = append(result, fn(v))
	}

	return newSlice(result)
}

// SelectMany creates a new slice by applying the given function to each element of the given slice and concatenating the results of the function
func (s *slice[T]) SelectMany[K any](fn func(T) []K) *slice[K] {
	result := []K{}

	for _, v := range s.s {
		result = append(result, fn(v)...)
	}

	return newSlice(result)
}

// SelectMapValues creates a new map from the given one by selecting a new element for each key according to the given function
func (m *dictionary[K, V]) Select[NewV any](fn func(key K, value V) NewV) *dictionary[K, NewV] {
	return m.Transform(func(k K, v V) (K, NewV) { return k, fn(k, v) })
}
