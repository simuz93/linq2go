package linq2go

// Where returns a new slice with the elements satisfying fn
//
//	FromSlice([]int{1, 2, 3}).Where(func(v int) bool { return v > 1 }).ToSlice() // [2 3]
func (s *slice[T]) Where(fn func(T) bool) *slice[T] {
	result := []T{}

	for _, v := range s.values {
		if fn(v) {
			result = append(result, v)
		}
	}

	return newSlice(result)
}

// Where returns a new dictionary with the elements satisfying fn
//
//	FromMap(map[string]int{"a": 1, "b": 2}).Where(func(k string, v int) bool { return v%2 == 0 }).ToMap() // map[b:2]
func (d *dictionary[K, V]) Where(fn func(K, V) bool) *dictionary[K, V] {
	result := map[K]V{}

	for k, v := range d.values {
		if fn(k, v) {
			result[k] = v
		}
	}

	return newDictionary(result)
}

// Where returns a new group with the buckets satisfying fn
//
//	FromSlice([]int{1, 2, 3}).Group(func(v int) bool { return v > 1 }).Where(func(k bool, bucket []int) bool { return len(bucket) > 1 }).ToMap() // map[true:[2 3]]
func (g *group[K, V]) Where(fn func(K, []V) bool) *group[K, V] {
	result := map[K][]V{}

	for k, v := range g.values {
		if fn(k, v) {
			result[k] = v
		}
	}

	return newGroup(result)
}
