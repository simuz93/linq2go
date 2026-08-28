package linq2go

/*
Group groups the given slice using the key selected by the groupBy function. Then it applies the fn function to each group.
*/
func (s *Slice[T]) Group[K comparable](fn func(T) K) *Group[K, T] {
	groups := make(map[K][]T)

	for _, v := range s.s {
		key := fn(v)

		if _, ok := groups[key]; !ok {
			groups[key] = []T{v}
		} else {
			groups[key] = append(groups[key], v)
		}
	}

	return &Group[K, T]{m: groups}
}

// GroupMap creates a new map by grouping every element of the given one into a slice sharing the same key, which is defined by the given function
func (m *Map[K, V]) Group[NewK comparable](fn func(key K, value V) NewK) *Group[NewK, V] {
	result := map[NewK][]V{}
	for k, v := range m.m {
		newK := fn(k, v)
		result[newK] = append(result[newK], v)
	}

	return &Group[NewK, V]{m: result}
}
