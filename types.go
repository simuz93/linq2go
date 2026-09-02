package linq2go

// slice wraps a Go slice: it is the type every slice query flows through
type slice[T any] struct {
	values []T
}

// newSlice wraps s without copying it: ownership passes to the wrapper. Only the public From constructors copy
func newSlice[T any](s []T) *slice[T] {
	return &slice[T]{values: s}
}

// dictionary wraps a Go map: it is the type every map query flows through
type dictionary[K comparable, V any] struct {
	values map[K]V
}

// newDictionary wraps m without copying it. See newSlice
func newDictionary[K comparable, V any](m map[K]V) *dictionary[K, V] {
	return &dictionary[K, V]{values: m}
}

// group wraps a map of buckets, as produced by Group, and exits only through an aggregation or ToMap
type group[K comparable, V any] struct {
	values map[K][]V
}

// newGroup wraps m without copying it. See newSlice
func newGroup[K comparable, V any](m map[K][]V) *group[K, V] {
	return &group[K, V]{values: m}
}
