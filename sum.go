package linq2go

// Sum returns the sum of the values selected by fn, accumulated in K: if K is a narrow
// integer type the sum overflows silently, and a single NaN makes the whole sum NaN, as
// it would in plain Go arithmetic
func (s *slice[T]) Sum[K Number](fn func(T) K) K {
	var sum K = 0

	for _, v := range s.s {
		sum += fn(v)
	}

	return sum
}

// Sum returns, for each group, the sum of the values selected by fn. See slice.Sum for the
// overflow and NaN caveats
func (g *group[K, V]) Sum[T Number](fn func(V) T) *dictionary[K, T] {
	result := make(map[K]T, len(g.m))

	for k, v := range g.m {
		result[k] = newSlice(v).Sum(fn)
	}

	return newDictionary(result)
}
