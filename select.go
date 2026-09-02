package linq2go

// Select returns a new slice with the value fn returns for each element
func (s *slice[T]) Select[K any](fn func(T) K) *slice[K] {
	result := make([]K, len(s.s))

	for i, v := range s.s {
		result[i] = fn(v)
	}

	return newSlice(result)
}

// SelectMany returns a new slice concatenating the slices fn returns for each element
func (s *slice[T]) SelectMany[K any](fn func(T) []K) *slice[K] {
	result := []K{}

	for _, v := range s.s {
		result = append(result, fn(v)...)
	}

	return newSlice(result)
}

// Select returns a new map with the same keys and the value fn returns for each entry
func (m *dictionary[K, V]) Select[NewV any](fn func(key K, value V) NewV) *dictionary[K, NewV] {
	return m.Transform(func(k K, v V) (K, NewV) { return k, fn(k, v) })
}
