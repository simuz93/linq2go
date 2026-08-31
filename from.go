package linq2go

import (
	"iter"
	"maps"
	"slices"
)

func FromSlice[T any](s []T) *slice[T] {
	return newSlice(slices.Clone(s))
}

func FromMap[K comparable, V any](m map[K]V) *dictionary[K, V] {
	return newDictionary(maps.Clone(m))
}

// FromSeq drains the given sequence into a slice and wraps it. A nil sequence yields an empty Slice
func FromSeq[T any](seq iter.Seq[T]) *slice[T] {
	result := []T{}

	if seq == nil {
		return newSlice(result)
	}

	return newSlice(slices.AppendSeq(result, seq))
}
