package linq2go

// Group gathers the elements into buckets sharing the key selected by fn
func (s *slice[T]) Group[K comparable](fn func(T) K) *group[K, T] {
	groups := make(map[K][]T)

	for _, v := range s.s {
		key := fn(v)
		groups[key] = append(groups[key], v)
	}

	return newGroup(groups)
}

// Group gathers the values of the map into buckets sharing the key selected by fn
func (m *dictionary[K, V]) Group[NewK comparable](fn func(key K, value V) NewK) *group[NewK, V] {
	result := map[NewK][]V{}
	for k, v := range m.m {
		newK := fn(k, v)
		result[newK] = append(result[newK], v)
	}

	return newGroup(result)
}
