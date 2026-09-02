package linq2go

// Avg returns the average of the values selected by fn, or 0 if the slice is empty,
// building on Sum and inheriting both its overflow and its NaN
//
//	FromSlice([]int{1, 2, 6}).Avg(func(v int) int { return v }) // 3
func (s *slice[T]) Avg[V Number](fn func(T) V) float64 {
	l := len(s.values)
	if l == 0 {
		return 0
	}
	sum := s.Sum(fn)

	return float64(sum) / float64(l)
}

// Avg returns, for each group, the average of the values selected by fn, with the overflow
// and NaN caveats of slice.Avg
//
//	FromSlice([]int{1, 2, 3, 4}).Group(func(v int) bool { return v%2 == 0 }).Avg(func(v int) int { return v }).ToMap() // map[false:2 true:3]
func (g *group[K, V]) Avg[NewV Number](fn func(V) NewV) *dictionary[K, float64] {
	result := make(map[K]float64, len(g.values))

	for k, v := range g.values {
		result[k] = newSlice(v).Avg(fn)
	}

	return newDictionary(result)
}
