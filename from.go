package linq2go

import (
	"iter"
	"maps"
	"slices"
)

// FromSlice wraps a copy of the given slice: later changes to s never affect the query
func FromSlice[T any](s []T) *slice[T] {
	return newSlice(slices.Clone(s))
}

// FromMap wraps a copy of the given map: later changes to m never affect the query
func FromMap[K comparable, V any](m map[K]V) *dictionary[K, V] {
	return newDictionary(maps.Clone(m))
}

// FromSeq drains the given sequence into a new slice and wraps it. A nil sequence yields an empty one
func FromSeq[T any](seq iter.Seq[T]) *slice[T] {
	result := []T{}

	if seq == nil {
		return newSlice(result)
	}

	return newSlice(slices.AppendSeq(result, seq))
}
