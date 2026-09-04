package linq2go

import "cmp"

// Max returns the index and the largest value selected by fn, calling it once per
// element, or -1 and 0 if the slice is empty; values compare with cmp.Compare, which
// sorts NaN before every number so a NaN never wins unless every value is one
//
//	FromSlice([]int{1, 5, 3}).Max(func(v int) int { return v }) // 1, 5
func (s *slice[T]) Max[V Number](fn func(T) V) (int, V) {
	bestIdx := -1
	var best V

	for i, e := range s.values {
		// the strict comparison keeps the first of equal maxima
		if v := fn(e); bestIdx < 0 || cmp.Compare(v, best) > 0 {
			bestIdx, best = i, v
		}
	}

	return bestIdx, best
}

// WhereMax returns the index and the element whose value selected by fn is the largest,
// or -1 and the zero value if the slice is empty
//
//	FromSlice([]string{"go", "linq"}).WhereMax(func(s string) int { return len(s) }) // 1, "linq"
func (s *slice[T]) WhereMax[V Number](fn func(T) V) (int, T) {
	idx, _ := s.Max(fn)
	if idx < 0 {
		return -1, *new(T)
	}

	return idx, s.values[idx]
}

// Max returns, for each group, the largest value selected by fn, comparing NaN as slice.Max does
//
//	FromSlice([]int{1, 2, 3, 4}).Group(func(v int) bool { return v%2 == 0 }).Max(func(v int) int { return v }).ToMap() // map[false:3 true:4]
func (g *group[K, V]) Max[NewV Number](fn func(V) NewV) *dictionary[K, NewV] {
	result := make(map[K]NewV, len(g.values))

	for k, v := range g.values {
		_, result[k] = newSlice(v).Max(fn)
	}

	return newDictionary(result)
}
