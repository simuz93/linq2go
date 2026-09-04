package linq2go

// Skip returns a new slice without the first n elements, clamping n to the slice bounds
//
//	FromSlice([]int{1, 2, 3}).Skip(1).ToSlice() // [2 3]
func (s *slice[T]) Skip(n int) *slice[T] {
	n = max(n, 0)
	l := len(s.values)
	n = min(n, l)

	// the capacity is capped so that appending to the result reallocates instead of writing into the receiver's array
	return newSlice(s.values[n:l:l])
}

// SkipWhile returns a new slice without the leading elements satisfying fn, keeping everything from the first that does not
//
//	FromSlice([]int{1, 2, 3, 1}).SkipWhile(func(v int) bool { return v < 3 }).ToSlice() // [3 1]
func (s *slice[T]) SkipWhile(fn func(T) bool) *slice[T] {
	n := 0
	for n < len(s.values) && fn(s.values[n]) {
		n++
	}

	return s.Skip(n)
}
