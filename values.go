package linq2go

// GetMapValues returns all the values of a map as a slice
func (m *dictionary[K, V]) Values() []V {
	return m.ToSlice(func(k K, v V) V { return v })
}
