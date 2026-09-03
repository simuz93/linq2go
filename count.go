package linq2go

// Count returns the length of the slice
//
//	FromSlice([]int{1, 2, 3}).Count() // 3
func (s *slice[T]) Count() int {
	return len(s.values)
}

// Count returns the number of key-value pairs in the map
//
//	FromMap(map[string]int{"a": 1, "b": 2}).Count()) // 2
func (d *dictionary[K, V]) Count() int {
	return len(d.values)
}

// Count returns the number of key-value pairs in the group
//
//	FromMap(map[string]int{"a": 2, "b": 4}).Group(func(k string, v int) bool { return v%2 == 0 }).Count()) // 1
func (g *group[K, V]) Count() int {
	return len(g.values)
}
