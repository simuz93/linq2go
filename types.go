package linq2go

type slice[T any] struct {
	s []T
}

type dictionary[K comparable, V any] struct {
	m map[K]V
}

type group[K comparable, V any] struct {
	m map[K][]V
}
