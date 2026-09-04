package linq2go

// Keys returns all the keys of the map as a slice, in unspecified order
//
//	FromMap(map[string]int{"a": 1, "b": 2}).Keys() // ["a" "b"], in any order
func (d *dictionary[K, V]) Keys() []K {
	return d.ToSlice(func(k K, v V) K { return k })
}
