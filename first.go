package linq2go

// FirstOrNil returns the first occurrence of an element that matches the given function or nil if no element matches it
func (s *Slice[T]) FirstOrNil(fn func(T) bool) *T {
	for _, v := range s.s {
		if fn(v) {
			return &v
		}
	}

	return nil
}

// FirstOrDefault returns the first occurrence of an element that matches the given function or its default value if no element matches it
func (s *Slice[T]) FirstOrDefault(fn func(T) bool) T {
	for _, v := range s.s {
		if fn(v) {
			return v
		}
	}

	return *new(T)
}
