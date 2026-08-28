package linq2go

func FromSlice[T any](s []T) *Slice[T] {
	return &Slice[T]{s: s}
}

func FromMap[K comparable, V any](m map[K]V) *Map[K, V] {
	return &Map[K, V]{m: m}
}
