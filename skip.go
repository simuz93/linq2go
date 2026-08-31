package linq2go

// Skip returns a new slice without the first n elements. n is clamped to the slice bounds
func (s *slice[T]) Skip(n int) *slice[T] {
	n = max(n, 0)
	n = min(n, len(s.s))
	return newSlice(s.s[n:])
}
