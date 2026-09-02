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
