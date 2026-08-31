package linq2go

type slice[T any] struct {
	s []T
}

// newSlice wraps the given slice without copying it: ownership passes to the
// returned wrapper. Used internally, where the slice is freshly built or is only
// read; the public From constructors are the ones that copy.
func newSlice[T any](s []T) *slice[T] {
	return &slice[T]{s: s}
}

type dictionary[K comparable, V any] struct {
	m map[K]V
}

// newDictionary wraps the given map without copying it: ownership passes to the
// returned wrapper. See newSlice.
func newDictionary[K comparable, V any](m map[K]V) *dictionary[K, V] {
	return &dictionary[K, V]{m: m}
}

type group[K comparable, V any] struct {
	m map[K][]V
}

// newGroup wraps the given map without copying it: ownership passes to the
// returned wrapper. See newSlice.
func newGroup[K comparable, V any](m map[K][]V) *group[K, V] {
	return &group[K, V]{m: m}
}
