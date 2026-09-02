package linq2go

// Transform returns a new map with the key-value pair fn returns for each entry. On a key collision which entry survives is not predictable
func (m *dictionary[K, V]) Transform[NewK comparable, NewV any](fn func(key K, value V) (NewK, NewV)) *dictionary[NewK, NewV] {
	result := make(map[NewK]NewV, len(m.values))
	for k, v := range m.values {
		newK, newV := fn(k, v)
		result[newK] = newV
	}

	return newDictionary(result)
}

// ChangeKey returns a new map with the same values re-keyed by fn. On a key collision which entry survives is not predictable
func (m *dictionary[K, V]) ChangeKey[NewK comparable](fn func(key K, value V) NewK) *dictionary[NewK, V] {
	return m.Transform(func(k K, v V) (NewK, V) { return fn(k, v), v })
}
