// Package linq2go provides LINQ-style fluent queries over Go slices and maps.
//
// A query is always enter → chain → exit: a From function wraps the input, every
// operator returns a new wrapper without mutating its receiver, and a To method —
// or a terminal operator such as Sum or FirstOrNil — hands back a plain Go value.
//
//	FromSlice([]int{1, 2, 3, 4}).Where(func(v int) bool { return v%2 == 0 }).ToSlice() // [2 4]
//
// The wrapper types are unexported, so a caller names them neither entering nor
// leaving a query. FromSlice and FromMap are the only functions that copy, which
// puts the caller's own data beyond the reach of the pipeline, and a result is
// never nil, so it marshals as [] and is safe to range over without a check.
//
// Requires Go 1.27 or newer, for generic methods.
package linq2go
