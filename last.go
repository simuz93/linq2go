package linq2go

// LastOrNil returns a pointer to a copy of the last element matching fn, or nil if none matches
//
//	FromSlice([]int{1, 2, 3}).LastOrNil(func(v int) bool { return v < 3 }) // pointer to 2
func (s *slice[T]) LastOrNil(fn func(T) bool) *T {
	for i := len(s.values) - 1; i >= 0; i-- {
		if fn(s.values[i]) {
			// the copy is declared here and not in the loop: a variable whose address is
			// returned is heap-allocated where it is declared, so declaring it in the loop
			// would allocate once per element scanned instead of once per match
			match := s.values[i]

			return &match
		}
	}

	return nil
}

// LastOrDefault returns the last element matching fn, or the zero value if none matches
//
//	FromSlice([]int{1, 2, 3}).LastOrDefault(func(v int) bool { return v > 5 }) // 0
func (s *slice[T]) LastOrDefault(fn func(T) bool) T {
	for i := len(s.values) - 1; i >= 0; i-- {
		if fn(s.values[i]) {
			return s.values[i]
		}
	}

	return *new(T)
}
