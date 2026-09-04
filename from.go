package linq2go

import (
	"iter"
	"maps"
	"slices"
)

// FromSlice wraps a copy of the given slice, so later changes to s never affect the query; a nil slice yields an empty one, never nil
//
//	FromSlice([]int{1, 2, 3}).Where(func(v int) bool { return v > 1 }).ToSlice() // [2 3]
func FromSlice[T any](s []T) *slice[T] {
	if s == nil {
		return newSlice([]T{})
	}

	return newSlice(slices.Clone(s))
}

// FromMap wraps a copy of the given map, so later changes to m never affect the query; a nil map yields an empty one, never nil
//
//	FromMap(map[string]int{"a": 1, "b": 2}).Values() // [1 2], in any order
func FromMap[K comparable, V any](m map[K]V) *dictionary[K, V] {
	if m == nil {
		return newDictionary(map[K]V{})
	}

	return newDictionary(maps.Clone(m))
}

// FromSeq drains the given sequence into a new slice and wraps it; a nil sequence yields an empty one, never nil
//
//	FromSeq(slices.Values([]int{1, 2, 3})).Sum(func(v int) int { return v }) // 6
func FromSeq[T any](seq iter.Seq[T]) *slice[T] {
	result := []T{}

	if seq == nil {
		return newSlice(result)
	}

	return newSlice(slices.AppendSeq(result, seq))
}

// FromSeq2 drains the given key-value sequence into a new map and wraps it; on a key
// collision the last pair wins, and a nil sequence yields an empty map, never nil
//
//	FromSeq2(maps.All(map[string]int{"a": 1, "b": 2})).ToMap() // map[a:1 b:2]
func FromSeq2[K comparable, V any](seq iter.Seq2[K, V]) *dictionary[K, V] {
	if seq == nil {
		return newDictionary(map[K]V{})
	}

	return newDictionary(maps.Collect(seq))
}
