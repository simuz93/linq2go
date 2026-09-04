package linq2go

import "cmp"

// Min returns the index and the smallest value selected by fn, calling it once per
// element, or -1 and 0 if the slice is empty; values compare with cmp.Compare, which
// sorts NaN before every number so a single NaN wins
//
//	FromSlice([]int{9, 4, 6}).Min(func(v int) int { return v }) // 1, 4
func (s *slice[T]) Min[V Number](fn func(T) V) (int, V) {
	bestIdx := -1
	var best V

	for i, e := range s.values {
		// the strict comparison keeps the first of equal minima
		if v := fn(e); bestIdx < 0 || cmp.Compare(v, best) < 0 {
			bestIdx, best = i, v
		}
	}

	return bestIdx, best
}

// WhereMin returns the index and the element whose value selected by fn is the smallest,
// or -1 and the zero value if the slice is empty
//
//	FromSlice([]string{"go", "linq"}).WhereMin(func(s string) int { return len(s) }) // 0, "go"
func (s *slice[T]) WhereMin[V Number](fn func(T) V) (int, T) {
	idx, _ := s.Min(fn)
	if idx < 0 {
		return -1, *new(T)
	}

	return idx, s.values[idx]
}

// Min returns, for each group, the smallest value selected by fn, comparing NaN as slice.Min does
//
//	FromSlice([]int{1, 2, 3, 4}).Group(func(v int) bool { return v%2 == 0 }).Min(func(v int) int { return v }).ToMap() // map[false:1 true:2]
func (g *group[K, V]) Min[NewV Number](fn func(V) NewV) *dictionary[K, NewV] {
	result := make(map[K]NewV, len(g.values))

	for k, v := range g.values {
		_, result[k] = newSlice(v).Min(fn)
	}

	return newDictionary(result)
}
