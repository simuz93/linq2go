package linq2go

// Count returns the length of the slice
//
//	FromSlice([]int{1, 2, 3}).Count() // 3
func (s *slice[T]) Count() int {
	return len(s.values)
}

// Count returns the number of key-value pairs in the map
//
//	FromMap(map[string]int{"a": 1, "b": 2}).Count() // 2
func (d *dictionary[K, V]) Count() int {
	return len(d.values)
}

// Count returns, for each group, the number of values in it; chain a second Count
// for the number of groups
//
//	FromSlice([]int{1, 2, 3, 4, 6}).Group(func(v int) bool { return v%2 == 0 }).Count().ToMap() // map[false:2 true:3]
func (g *group[K, V]) Count() *dictionary[K, int] {
	result := make(map[K]int, len(g.values))

	for k, v := range g.values {
		result[k] = newSlice(v).Count()
	}

	return newDictionary(result)
}
