package linq2go

import (
	"iter"
	"maps"
	"slices"
)

// FromSlice wraps a copy of the given slice: later changes to s never affect the query.
// A nil slice yields an empty one, never nil
func FromSlice[T any](s []T) *slice[T] {
	if s == nil {
		return newSlice([]T{})
	}

	return newSlice(slices.Clone(s))
}

// FromMap wraps a copy of the given map: later changes to m never affect the query.
// A nil map yields an empty one, never nil
func FromMap[K comparable, V any](m map[K]V) *dictionary[K, V] {
	if m == nil {
		return newDictionary(map[K]V{})
	}

	return newDictionary(maps.Clone(m))
}

// FromSeq drains the given sequence into a new slice and wraps it. A nil sequence yields an empty one, never nil
func FromSeq[T any](seq iter.Seq[T]) *slice[T] {
	result := []T{}

	if seq == nil {
		return newSlice(result)
	}

	return newSlice(slices.AppendSeq(result, seq))
}
