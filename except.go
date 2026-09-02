package linq2go

// Except returns a new slice with the elements that match no value in slice according to fn, called as fn(element, value); duplicates are kept
//
//	FromSlice([]int{1, 2, 2, 3}).Except([]int{2}, func(a, b int) bool { return a == b }).ToSlice() // [1 3]
func (s *slice[T]) Except(slice []T, fn func(T, T) bool) *slice[T] {
	result := []T{}

	values := newSlice(slice)

	for _, elem := range s.values {
		if !values.Contains(elem, fn) {
			result = append(result, elem)
		}
	}

	return newSlice(result)
}
