package linq2go

// All reports whether every element satisfies fn; an empty slice returns true
//
//	FromSlice([]int{2, 4, 6}).All(func(v int) bool { return v%2 == 0 }) // true
func (s *slice[T]) All(fn func(T) bool) bool {
	for _, v := range s.values {
		if !fn(v) {
			return false
		}
	}

	return true
}
