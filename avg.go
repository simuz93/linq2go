package linq2go

// Avg returns the average of the values selected by fn, or 0 if the slice is empty.
// It builds on Sum, so it inherits both its overflow and its NaN
func (s *slice[T]) Avg[K Number](fn func(T) K) float64 {
	l := len(s.s)
	if l == 0 {
		return 0
	}
	sum := s.Sum(fn)

	return float64(sum) / float64(l)
}

// Avg returns, for each group, the average of the values selected by fn. See slice.Avg for the
// overflow and NaN caveats
func (g *group[K, V]) Avg[T Number](fn func(V) T) *dictionary[K, float64] {
	result := make(map[K]float64, len(g.m))

	for k, v := range g.m {
		result[k] = newSlice(v).Avg(fn)
	}

	return newDictionary(result)
}
