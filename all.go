package linq2go

// All reports whether every element satisfies fn. An empty slice returns true
func (s *slice[T]) All(fn func(T) bool) bool {
	for _, v := range s.values {
		if !fn(v) {
			return false
		}
	}

	return true
}
