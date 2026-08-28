package linq2go

// Where creates a new slice by filtering the given slice according to the given function
func (s *slice[T]) Where(fn func(T) bool) *slice[T] {
	result := []T{}

	for _, v := range s.s {
		if fn(v) {
			result = append(result, v)
		}
	}

	return FromSlice(result)
}
