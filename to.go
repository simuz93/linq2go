package linq2go

// ToSlice returns the wrapped slice itself, not a copy: writing to it also changes the wrapper
//
//	FromSlice([]int{1, 2, 3}).Where(func(v int) bool { return v > 1 }).ToSlice() // [2 3]
func (s *slice[T]) ToSlice() []T {
	return s.values
}

// ToMap builds a map from the key-value pairs fn returns for each element; on a key collision the first element wins
//
//	FromSlice([]string{"a", "b"}).ToMap(func(s string, i int) (string, int) { return s, i }) // map[a:0 b:1]
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
//
//	FromMap(map[string]int{"a": 1}).ToMap() // map[a:1]
func (m *dictionary[K, V]) ToMap() map[K]V {
	return m.values
}

// ToSlice returns a slice with the value fn returns for each entry, in unspecified order
//
//	FromMap(map[string]int{"a": 1, "b": 2}).ToSlice(func(k string, v int) int { return v }) // [1 2], in any order
func (m *dictionary[K, V]) ToSlice[T any](fn func(key K, value V) T) []T {
	s := make([]T, 0, len(m.values))

	for key, value := range m.values {
		s = append(s, fn(key, value))
	}

	return s
}

// ToMap returns the wrapped map of buckets itself, not a copy: writing to it also changes the wrapper
//
//	FromSlice([]int{1, 2, 3}).Group(func(v int) bool { return v > 1 }).ToMap() // map[false:[1] true:[2 3]]
func (g *group[K, V]) ToMap() map[K][]V {
	return g.values
}
