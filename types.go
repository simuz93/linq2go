package linq2go

// slice wraps a Go slice: it is the type every slice query flows through
//
//	FromSlice(src).Where(fn).Select(fn).ToSlice()
type slice[T any] struct {
	values []T
}

// newSlice wraps s without copying it, passing ownership to the wrapper; only the public From constructors copy
//
//	newSlice(v).Sum(fn) // how the group aggregations delegate to the slice ones
func newSlice[T any](s []T) *slice[T] {
	return &slice[T]{values: s}
}

// dictionary wraps a Go map: it is the type every map query flows through
//
//	FromMap(src).ChangeKey(fn).ToMap()
type dictionary[K comparable, V any] struct {
	values map[K]V
}

// newDictionary wraps m without copying it, like newSlice
//
//	newDictionary(result)
func newDictionary[K comparable, V any](m map[K]V) *dictionary[K, V] {
	return &dictionary[K, V]{values: m}
}

// group wraps a map of buckets, as produced by Group, and exits only through an aggregation or ToMap
//
//	FromSlice(src).Group(fn).Sum(fn).ToMap()
type group[K comparable, V any] struct {
	values map[K][]V
}

// newGroup wraps m without copying it, like newSlice
//
//	newGroup(groups)
func newGroup[K comparable, V any](m map[K][]V) *group[K, V] {
	return &group[K, V]{values: m}
}
