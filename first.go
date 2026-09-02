package linq2go

// FirstOrNil returns a pointer to a copy of the first element matching fn, or nil if none matches
func (s *slice[T]) FirstOrNil(fn func(T) bool) *T {
	for _, v := range s.values {
		if fn(v) {
			return &v
		}
	}

	return nil
}

// FirstOrDefault returns the first element matching fn, or the zero value if none matches
func (s *slice[T]) FirstOrDefault(fn func(T) bool) T {
	for _, v := range s.values {
		if fn(v) {
			return v
		}
	}

	return *new(T)
}
