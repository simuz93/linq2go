package linq2go

// Select returns a new slice with the value fn returns for each element
//
//	FromSlice([]int{1, 2}).Select(func(v int) string { return strconv.Itoa(v) }).ToSlice() // ["1" "2"]
func (s *slice[T]) Select[NewT any](fn func(T) NewT) *slice[NewT] {
	result := make([]NewT, len(s.values))

	for i, v := range s.values {
		result[i] = fn(v)
	}

	return newSlice(result)
}

// SelectMany returns a new slice concatenating the slices fn returns for each element
//
//	FromSlice([][]int{{1, 2}, {3}}).SelectMany(func(v []int) []int { return v }).ToSlice() // [1 2 3]
func (s *slice[T]) SelectMany[NewT any](fn func(T) []NewT) *slice[NewT] {
	result := []NewT{}

	for _, v := range s.values {
		result = append(result, fn(v)...)
	}

	return newSlice(result)
}

// Select returns a new map with the same keys and the value fn returns for each entry
//
//	FromMap(map[string]int{"a": 1}).Select(func(k string, v int) int { return v * 10 }).ToMap() // map[a:10]
func (m *dictionary[K, V]) Select[NewV any](fn func(key K, value V) NewV) *dictionary[K, NewV] {
	return m.Transform(func(k K, v V) (K, NewV) { return k, fn(k, v) })
}
