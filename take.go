package linq2go

// TrimEnd trims the the given value from the end of the slice
func (s *slice[T]) Take(n int) *slice[T] {
	n = max(n, 0)
	n = min(n, len(s.s))
	return newSlice(s.s[:n:n])
}
