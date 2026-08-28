package linq2go

// TransformMap creates a new map from the given one by selecting a new key-value pair for each existing one
func (m *dictionary[K, V]) Transform[NewK comparable, NewV any](fn func(key K, value V) (NewK, NewV)) *dictionary[NewK, NewV] {
	result := map[NewK]NewV{}
	for k, v := range m.m {
		newK, newV := fn(k, v)
		result[newK] = newV
	}

	return FromMap(result)
}

// ChangeMapKey creates a new map from the given one assigning each element to the new key created by the given function
func (m *dictionary[K, V]) ChangeKey[NewK comparable](fn func(key K, value V) NewK) *dictionary[NewK, V] {
	return m.Transform(func(k K, v V) (NewK, V) { return fn(k, v), v })
}
