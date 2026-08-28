package linq2go

// TrimEnd trims the the given value from the end of the slice
func (s *Slice[T]) Take(n int) *Slice[T] {
	n = max(n, 0)
	n = min(n, len(s.s))
	return FromSlice(s.s[:n:n])
}
