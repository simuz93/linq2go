# linq2go

LINQ-style fluent queries for Go slices and maps.

Go's standard library gives you `slices` and `maps`, but chaining them means naming an intermediate variable at every step. `linq2go` wraps a slice or a map once, lets you chain operators on it, and hands you a plain Go value back at the end.

```go
report := linq2go.FromSlice(orders).
	Where(func(o order) bool { return o.Total > 0 }).
	Group(func(o order) string { return o.Customer }).
	Sum(func(o order) float64 { return o.Total }).
	ToMap()
// map[string]float64{"alice": 17, "bob": 5}
```

It is a small, dependency-light library: one package, no code generation, no reflection, and no runtime beyond the standard library plus a single constraint from `golang.org/x/exp`.

## Requirements

**Go 1.27 or newer.** The fluent design rests on *generic methods* — methods that declare their own type parameters:

```go
func (s *slice[T]) Select[NewT any](fn func(T) NewT) *slice[NewT]
```

Without them, `Select` could not change the element type while staying a method, and the whole chain would collapse into nested function calls. Older toolchains reject this with `method must have no type parameters`. `go.mod` pins `go 1.27.0`, which the `go` command downloads on its own when needed.

```sh
go get github.com/simuz93/linq2go
```

## Usage

A query is always **enter → chain → exit**.

```go
import "github.com/simuz93/linq2go"

evens := linq2go.FromSlice([]int{1, 2, 3, 4, 5, 6}).
	Where(func(v int) bool { return v%2 == 0 }).
	Select(func(v int) string { return strconv.Itoa(v) }).
	ToSlice()
// []string{"2", "4", "6"}
```

**Enter** — `FromSlice`, `FromMap`, `FromSeq` (from an `iter.Seq`).
**Exit** — `ToSlice`, `ToMap`, `Values`, or a terminal operator: `Sum`, `Min`, `Max`, `Avg`, `WhereMin`, `WhereMax`, `Any`, `All`, `Contains`, `FirstOrNil`, `FirstOrDefault`.

### Grouping and aggregating

`Group` buckets elements by a key; the aggregations then reduce each bucket and return a map keyed by group.

```go
type sale struct {
	Region string
	Amount float64
}

sales := []sale{
	{"north", 100}, {"south", 250}, {"north", 50}, {"south", 75},
}

byRegion := linq2go.FromSlice(sales).Group(func(s sale) string { return s.Region })

byRegion.Sum(func(s sale) float64 { return s.Amount }).ToMap() // map[north:150 south:325]
byRegion.Avg(func(s sale) float64 { return s.Amount }).ToMap() // map[north:75 south:162.5]
byRegion.Max(func(s sale) float64 { return s.Amount }).ToMap() // map[north:100 south:250]
```

`WhereMin` and `WhereMax` are the pair that returns the *element* rather than the value — the cheapest way to answer "which one is the largest?":

```go
biggest := linq2go.FromSlice(sales).WhereMax(func(s sale) float64 { return s.Amount })
// sale{Region: "south", Amount: 250}
```

### Working with maps

```go
prices := map[string]float64{"apple": 1.20, "pear": 2.00}

withVat := linq2go.FromMap(prices).
	Select(func(name string, price float64) float64 { return price * 1.22 }).
	ToMap()
// map[apple:1.464 pear:2.44]

upper := linq2go.FromMap(prices).
	ChangeKey(func(name string, price float64) string { return strings.ToUpper(name) }).
	ToMap()
// map[APPLE:1.2 PEAR:2]
```

`Transform` changes key and value together, `Select` only the value, `ChangeKey` only the key, and `Values` returns the values as a slice.

### Comparing without `comparable`

Set-like operators take an explicit comparison function instead of constraining `T` to `comparable`, so they work on structs that contain slices, maps or functions:

```go
type user struct {
	ID    int
	Roles []string // makes user non-comparable
}

active := linq2go.FromSlice(all).
	Except(banned, func(a, b user) bool { return a.ID == b.ID }).
	ToSlice()
```

The same signature backs `Contains` and `Intersect`. `Contains` calls it as `fn(value, element)` — the value you are searching for comes first.

### From an iterator

```go
squares := func(yield func(int) bool) {
	for i := 1; i <= 5; i++ {
		if !yield(i * i) {
			return
		}
	}
}

linq2go.FromSeq(squares).Sum(func(v int) int { return v }) // 55
```

## Operators

**On a slice**

| Operator | Result |
|---|---|
| `Where(fn)` | keeps the elements matching `fn` |
| `Select(fn)` / `SelectMany(fn)` | maps each element / maps and flattens |
| `Distinct(fn)` | drops duplicates by the key `fn` selects, keeping the first |
| `Take(n)` / `Skip(n)` | first `n` / all but the first `n`, both clamped to the bounds |
| `Concat(values)` | the receiver followed by `values` |
| `Except(values, fn)` / `Intersect(values, fn)` | excludes / keeps what `fn` matches |
| `Contains(value, fn)` | whether any element matches, as `fn(value, element)` |
| `Group(fn)` | buckets sharing the key `fn` selects |
| `Sum` `Min` `Max` `Avg` | numeric reductions over the selected value |
| `WhereMin` `WhereMax` | the *element* holding the smallest / largest value |
| `Any(fn)` / `All(fn)` | whether some / every element matches |
| `FirstOrNil(fn)` / `FirstOrDefault(fn)` | first match, as a pointer to a copy / as a value |
| `ToSlice()` / `ToMap(fn)` | exit |

**On a map**

| Operator | Result |
|---|---|
| `Select(fn)` | a new value for each key |
| `Transform(fn)` | a new key **and** value for each entry |
| `ChangeKey(fn)` | re-keys, values untouched |
| `Group(fn)` | buckets the values by a new key |
| `Values()` / `ToSlice(fn)` / `ToMap()` | exit |

**On a group** — `Sum`, `Min`, `Max`, `Avg`, each returning a map keyed by group, plus `ToMap`.

## Design notes

The parts of the design that are deliberate, and the reasoning behind them.

### The wrapper types are unexported

`FromSlice` returns a `*slice[T]` you cannot name, only chain on. That is on purpose: it means no caller can construct a wrapper around a value the library has not copied, which is what makes the copy contract below airtight. The cost is that you cannot store a half-built query in a struct field or pass it across a package boundary — a query is meant to be written and consumed in one expression.

### Copies happen once, on entry

`FromSlice` and `FromMap` clone their input. Nothing downstream clones anything: operators build fresh results, `Take` and `Skip` return capacity-capped sub-slices of the receiver's array (so appending to their result reallocates rather than writing into the shared array), and the `To*` exits return the internal storage as-is.

```go
src := []int{1, 2, 3}
q := linq2go.FromSlice(src)
src[0] = 99
q.ToSlice() // [1 2 3] — your mutation cannot reach the query
```

An earlier version copied at every stage, which is the obvious way to guarantee that and also the wasteful one: with the single entry copy, a `From → Where → Skip → Take → ToSlice` chain over 10.000 elements went from 62,8 to 44,2 µs and from 605 to 439 KB, and aggregating over a group dropped from 82 KB to 200 B per call, because each bucket was being cloned only to be summed.

The trade-off is on the way out: since `ToSlice` and `ToMap` hand back the internal storage rather than a defensive copy, writing into the result writes into the wrapper it came from, and two wrappers derived from the same `From*` share one array. Your own data is never at risk — only the library's private copy is — and it is documented on each exit.

### Nothing is ever nil

A `nil` input, a filter that discards everything, a `Concat` of two empty slices: all yield an empty, non-nil result. A query result can always be marshalled as `[]`, never as `null`, and never needs a nil check before ranging.

### The standard library does the work where it can

`Any` is `slices.ContainsFunc`, `WhereMin`/`WhereMax` are `slices.MinFunc`/`MaxFunc`, `Concat` is `slices.Concat`, the entry copies are `slices.Clone`/`maps.Clone`. Reimplementing them would mean maintaining a second set of edge cases and diverging from what a Go reader already expects. It also means the library inherits their semantics — see `NaN` below.

### Sizes are pre-allocated only when they are exactly known

`Select`, `Transform`, `dictionary.ToSlice` and the group aggregations build their result with `make` at the final size, which is known from the input. `Where`, `Distinct` and `slice.ToMap` deliberately do not: there the input length is only an upper bound, and pre-allocating it wastes memory that the caller then keeps, since exits return the internal storage. Measured on 10.000 elements with 100 distinct values, pre-sizing `Distinct` would grow it from 6,4 KB to 377 KB.

### What it is not

No laziness — every operator runs eagerly and materialises its result; there is no streaming or short-circuiting across a chain. No error handling — operators take pure selectors and cannot fail. No parallelism. If you need any of those, this is the wrong shape of library.

## Behaviour worth knowing

Documented, deliberate, and easy to mistake for bugs:

- **Empty input gives the zero value.** `Min`/`Max`/`Avg` return `0` and `FirstOrDefault` returns the zero value — indistinguishable from a legitimate result. `FirstOrNil` is the exception, returning a pointer to a copy, or `nil`.
- **`NaN` follows `cmp.Compare`**, which sorts it below every number: a single `NaN` wins a minimum and never wins a maximum, while `Sum` and `Avg` propagate it like plain Go arithmetic. Note this differs from `slices.Min`/`slices.Max`, which propagate `NaN` in both directions.
- **`Sum` accumulates in the type `fn` returns**, so a narrow integer overflows silently, and `Avg` inherits the wrong total.
- **`Transform` and `ChangeKey` are not deterministic on key collisions** — the survivor depends on Go's randomized map iteration order, and is not stable between runs. `slice.ToMap` instead keeps the first occurrence.
- **`Except` and `Intersect` keep duplicates.** They include or exclude values rather than building a set; chain `Distinct` if you want one. This diverges from LINQ.
- **`dictionary.ToSlice` and `Values` return an unordered result**, being map iterations.

## Contributing

Issues and pull requests are welcome.

```sh
go build ./...
go vet ./...
go test -race -cover ./...
go fmt ./...            # not bare `gofmt`: see below
```

Always drive the tools through the `go` command, which resolves the toolchain pinned in `go.mod`. A `gofmt` binary earlier on your `PATH` may come from an older Go and will report `method must have no type parameters` on every generic method — that is a stale tool, not a real error.

### Guidelines

- **One operator per file**, named after the operator: `where.go`, `distinct.go`, `min.go` and so on, with no exceptions. A file holds every receiver of that operator — `min.go` carries `Min` and `WhereMin` for slices and `Min` for groups. Add a new operator as a new file rather than appending to an existing one.
- **The selector or predicate goes last**, and is named `fn`. Comparison-based operators take an explicit `fn func(T, T) bool` instead of constraining `T` to `comparable`.
- **Never call `FromSlice` or `FromMap` from inside the package.** Use the internal `newSlice`/`newDictionary`/`newGroup` constructors: a `From*` call in an operator clones data that is already private, which is exactly what the design exists to avoid.
- **Never return nil.** Initialise results as `result := []T{}` or `make([]T, 0, n)`, and guard the paths where the standard library can hand back a nil.
- **Doc comments are a concise description plus one simple usage example**, in the style of `where.go`. The description opens with the identifier's name and stays short, even when it spans lines; a behaviour that cannot be read off the signature goes in a subordinate clause or after a semicolon. The example sits after a blank comment line, tab-indented so godoc renders it as a code block, with the actual result in a trailing comment — run it before writing it down.
- **Type parameters follow their role**: `T` for an element or a generic result, `K` and `V` for a key and a value (alone too — `Distinct[K]` selects a key, `Sum[V]` selects a value), and a `New` prefix when the receiver already binds the name (`Select[NewT]`, `Transform[NewK, NewV]`).
- **Back performance claims with a benchmark.** Several plausible optimisations in this library measured *worse* once tried; a `-benchmem` run before and after settles it.
- Prefer wrapping a standard library function over reimplementing it.

## License

Not yet chosen — add a `LICENSE` file before publishing.
