package linq2go

/*
All returns true if all of the elements in the slice satisfy the fn condition.
If the slice has no elements, All returns true.
*/
func (s *slice[T]) All(fn func(T) bool) bool {
	for _, v := range s.s {
		if !fn(v) {
			return false
		}
	}

	return true
}
