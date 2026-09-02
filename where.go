package linq2go

// Where returns a new slice with the elements satisfying fn
//
//	FromSlice([]int{1, 2, 3}).Where(func(v int) bool { return v > 1 }).ToSlice() // [2 3]
func (s *slice[T]) Where(fn func(T) bool) *slice[T] {
	result := []T{}

	for _, v := range s.values {
		if fn(v) {
			result = append(result, v)
		}
	}

	return newSlice(result)
}
