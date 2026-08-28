package linq2go

type Slice[T any] struct {
	s []T
}

type Map[K comparable, V any] struct {
	m map[K]V
}

type Group[K comparable, V any] struct {
	m map[K][]V
}
