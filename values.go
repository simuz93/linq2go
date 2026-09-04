package linq2go

// Values returns all the values of the map as a slice, in unspecified order
//
//	FromMap(map[string]int{"a": 1, "b": 2}).Values() // [1 2], in any order
func (d *dictionary[K, V]) Values() []V {
	return d.ToSlice(func(k K, v V) V { return v })
}
