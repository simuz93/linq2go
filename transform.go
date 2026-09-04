package linq2go

// Transform returns a new map with the key-value pair fn returns for each entry; on a key collision which entry survives is not predictable
//
//	FromMap(map[string]int{"a": 1}).Transform(func(k string, v int) (int, string) { return v, k }).ToMap() // map[1:a]
func (d *dictionary[K, V]) Transform[NewK comparable, NewV any](fn func(key K, value V) (NewK, NewV)) *dictionary[NewK, NewV] {
	result := make(map[NewK]NewV, len(d.values))
	for k, v := range d.values {
		newK, newV := fn(k, v)
		result[newK] = newV
	}

	return newDictionary(result)
}

// ChangeKey returns a new map with the same values re-keyed by fn; on a key collision which entry survives is not predictable
//
//	FromMap(map[string]int{"a": 1}).ChangeKey(func(k string, v int) string { return strings.ToUpper(k) }).ToMap() // map[A:1]
func (d *dictionary[K, V]) ChangeKey[NewK comparable](fn func(key K, value V) NewK) *dictionary[NewK, V] {
	return d.Transform(func(k K, v V) (NewK, V) { return fn(k, v), v })
}
