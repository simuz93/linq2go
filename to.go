package linq2go

// ToSlice returns the wrapped slice itself, not a copy: writing to it also changes the wrapper
func (s *slice[T]) ToSlice() []T {
	return s.values
}

// ToMap builds a map from the key-value pairs fn returns for each element. On a key collision the first element wins
func (s *slice[T]) ToMap[K comparable, V any](fn func(elem T, index int) (K, V)) map[K]V {
	m := map[K]V{}
	for i := range s.values {
		key, value := fn(s.values[i], i)
		if _, ok := m[key]; !ok {
			m[key] = value
		}
	}
	return m
}

// ToMap returns the wrapped map itself, not a copy: writing to it also changes the wrapper
func (m *dictionary[K, V]) ToMap() map[K]V {
	return m.values
}

// ToSlice returns a slice with the value fn returns for each entry, in unspecified order
func (m *dictionary[K, V]) ToSlice[T any](fn func(key K, value V) T) []T {
	s := make([]T, 0, len(m.values))

	for key, value := range m.values {
		s = append(s, fn(key, value))
	}

	return s
}

// ToMap returns the wrapped map of buckets itself, not a copy: writing to it also changes the wrapper
func (g *group[K, V]) ToMap() map[K][]V {
	return g.values
}
