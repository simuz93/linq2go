package linq2go

import (
	"iter"
	"slices"
)

func FromSlice[T any](s []T) *slice[T] {
	return &slice[T]{s: s}
}

func FromMap[K comparable, V any](m map[K]V) *dictionary[K, V] {
	return &dictionary[K, V]{m: m}
}

// FromSeq drains the given sequence into a slice and wraps it. A nil sequence yields an empty Slice
func FromSeq[T any](seq iter.Seq[T]) *slice[T] {
	result := []T{}

	if seq == nil {
		return FromSlice(result)
	}

	return FromSlice(slices.AppendSeq(result, seq))
}
