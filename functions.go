package linq2go

// Self returns its argument unchanged: it is the selector to pass wherever an operator asks which
// key to derive from an element and the element itself is the key
//
//	FromSlice([]int{1, 2, 2, 3}).Distinct(Self).ToSlice() // [1 2 3]
func Self[V any](v V) V {
	return v
}

// Equal reports whether two values are equal: it is the comparison to pass wherever an operator
// asks how to match two elements and plain equality is enough
//
//	FromSlice([]int{1, 2, 3}).Contains(2, Equal) // true
func Equal[V comparable](a, b V) bool {
	return a == b
}
