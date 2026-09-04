package linq2go

// FirstOrNil returns a pointer to a copy of the first element matching fn, or nil if none matches
//
//	FromSlice([]int{1, 2, 3}).FirstOrNil(func(v int) bool { return v > 1 }) // pointer to 2
func (s *slice[T]) FirstOrNil(fn func(T) bool) *T {
	for _, v := range s.values {
		if fn(v) {
			// the copy is declared here and not in the loop header: a variable whose address
			// is returned is heap-allocated where it is declared, so returning &v directly
			// would allocate the loop variable once per element scanned instead of once per match
			match := v

			return &match
		}
	}

	return nil
}

// FirstOrDefault returns the first element matching fn, or the zero value if none matches
//
//	FromSlice([]int{1, 2, 3}).FirstOrDefault(func(v int) bool { return v > 5 }) // 0
func (s *slice[T]) FirstOrDefault(fn func(T) bool) T {
	for _, v := range s.values {
		if fn(v) {
			return v
		}
	}

	return *new(T)
}
