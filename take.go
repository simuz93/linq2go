package linq2go

// Take returns a new slice with the first n elements. n is clamped to the slice bounds
func (s *slice[T]) Take(n int) *slice[T] {
	n = max(n, 0)
	n = min(n, len(s.s))
	return newSlice(s.s[:n:n])
}
