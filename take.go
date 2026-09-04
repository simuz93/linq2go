package linq2go

// Take returns a new slice with the first n elements, clamping n to the slice bounds
//
//	FromSlice([]int{1, 2, 3}).Take(2).ToSlice() // [1 2]
func (s *slice[T]) Take(n int) *slice[T] {
	n = max(n, 0)
	n = min(n, len(s.values))

	// the capacity is capped so that appending to the result reallocates instead of writing into the receiver's array
	return newSlice(s.values[:n:n])
}

// TakeWhile returns a new slice with the leading elements satisfying fn, stopping at the first that does not
//
//	FromSlice([]int{1, 2, 3, 1}).TakeWhile(func(v int) bool { return v < 3 }).ToSlice() // [1 2]
func (s *slice[T]) TakeWhile(fn func(T) bool) *slice[T] {
	n := 0
	for n < len(s.values) && fn(s.values[n]) {
		n++
	}

	return s.Take(n)
}
