package linq2go

import (
	"iter"
	"slices"
)

// SeqToSlice converts a sequence of sequences into a single slice
func SeqToSlice[T any](seqs ...iter.Seq[T]) []T {
	res := []T{}

	for _, seq := range seqs {
		if seq == nil {
			continue
		}
		res = slices.AppendSeq(res, seq)
	}

	return res
}
