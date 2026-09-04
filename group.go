package linq2go

// Group gathers the elements into buckets sharing the key selected by fn
//
//	FromSlice([]int{1, 2, 3, 4}).Group(func(v int) bool { return v%2 == 0 }).ToMap() // map[false:[1 3] true:[2 4]]
func (s *slice[T]) Group[K comparable](fn func(T) K) *group[K, T] {
	groups := make(map[K][]T)

	for _, v := range s.values {
		key := fn(v)
		groups[key] = append(groups[key], v)
	}

	return newGroup(groups)
}

// Group gathers the values of the map into buckets sharing the key selected by fn
//
//	FromMap(map[string]int{"a": 1, "b": 2}).Group(func(k string, v int) bool { return v%2 == 0 }).ToMap() // map[false:[1] true:[2]]
func (d *dictionary[K, V]) Group[NewK comparable](fn func(key K, value V) NewK) *group[NewK, V] {
	result := map[NewK][]V{}
	for k, v := range d.values {
		newK := fn(k, v)
		result[newK] = append(result[newK], v)
	}

	return newGroup(result)
}
