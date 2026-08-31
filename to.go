package linq2go

func (s *slice[T]) ToSlice() []T {
	return s.s
}

// MapFromSlice creates a map from a slice by applying the given function to each element of the slice
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

func (m *dictionary[K, V]) ToMap() map[K]V {
	return m.m
}

// MapToSlice converts a map to a slice by applying the given function to each element of the map. The resulting slice will be unordered
func (m *dictionary[K, V]) ToSlice[T any](fn func(key K, value V) T) []T {
	s := []T{}

	for key := range m.m {
		s = append(s, fn(key, m.m[key]))
	}
	return s
}

func (g *group[K, V]) ToMap() map[K][]V {
	return g.m
}
