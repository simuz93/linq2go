package linq2go

import "slices"

// Concat concatenates the given slices. Shortcut for slices.Concat, left here for retrocompatibility
func (s *Slice[T]) Concat(values ...[]T) *Slice[T] {
	return FromSlice(slices.Concat(values...))
}
