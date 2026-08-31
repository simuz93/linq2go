package linq2go

// ToSlice returns the wrapped slice itself, not a copy: writing to it also changes the wrapper
func (s *slice[T]) ToSlice() []T {
	return s.s
}

// ToMap builds a map from the key-value pairs fn returns for each element. On a key collision the first element wins
func (s *slice[T]) ToMap[K comparable, V any](fn func(elem T, index int) (K, V)) map[K]V {
	m := map[K]V{}
	for i := range s.s {
		key, value := fn(s.s[i], i)
		if _, ok := m[key]; !ok {
			m[key] = value
		}
	}
	return m
}

// ToMap returns the wrapped map itself, not a copy: writing to it also changes the wrapper
func (m *dictionary[K, V]) ToMap() map[K]V {
	return m.m
}

// ToSlice returns a slice with the value fn returns for each entry, in unspecified order
func (m *dictionary[K, V]) ToSlice[T any](fn func(key K, value V) T) []T {
	s := []T{}

	for key := range m.m {
		s = append(s, fn(key, m.m[key]))
	}
	return s
}

// ToMap returns the wrapped map of buckets itself, not a copy: writing to it also changes the wrapper
func (g *group[K, V]) ToMap() map[K][]V {
	return g.m
}
