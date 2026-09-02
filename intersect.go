package linq2go

// Intersect returns a new slice with the elements that match a value in slice according to fn, called as fn(element, value); duplicates are kept
//
//	FromSlice([]int{1, 2, 2, 3}).Intersect([]int{2, 3}, func(a, b int) bool { return a == b }).ToSlice() // [2 2 3]
func (s *slice[T]) Intersect(slice []T, fn func(T, T) bool) *slice[T] {
	result := []T{}

	values := newSlice(slice)

	for _, elem := range s.values {
		if values.Contains(elem, fn) {
			result = append(result, elem)
		}
	}

	return newSlice(result)
}
