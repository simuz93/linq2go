package linq2go

// TrimEnd trims the the given value from the end of the slice
func (s *Slice[T]) Skip(n int) *Slice[T] {
	return FromSlice(s.s[n:])
}
