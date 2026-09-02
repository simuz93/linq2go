package linq2go

// Sum returns the sum of the values selected by fn, accumulated in V: if V is a narrow
// integer type the sum overflows silently, and a single NaN makes the whole sum NaN, as
// it would in plain Go arithmetic
//
//	FromSlice([]int{1, 2, 3}).Sum(func(v int) int { return v }) // 6
func (s *slice[T]) Sum[V Number](fn func(T) V) V {
	var sum V = 0

	for _, v := range s.values {
		sum += fn(v)
	}

	return sum
}

// Sum returns, for each group, the sum of the values selected by fn, with the overflow and
// NaN caveats of slice.Sum
//
//	FromSlice([]int{1, 2, 3, 4}).Group(func(v int) bool { return v%2 == 0 }).Sum(func(v int) int { return v }).ToMap() // map[false:4 true:6]
func (g *group[K, V]) Sum[NewV Number](fn func(V) NewV) *dictionary[K, NewV] {
	result := make(map[K]NewV, len(g.values))

	for k, v := range g.values {
		result[k] = newSlice(v).Sum(fn)
	}

	return newDictionary(result)
}
