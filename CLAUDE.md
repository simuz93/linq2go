# CLAUDE.md

Working notes for Claude Code on this repository: what an agent needs in order to edit this code
correctly, plus the findings that cost a measurement to get and should not be re-derived. Anything
addressed to someone *using* or *contributing to* the library — what it does, examples, design
rationale, contribution guidelines — belongs in `README.md`, not here.

## Project

`github.com/simuz93/linq2go` — LINQ-style fluent queries over Go slices and maps. A single flat
package `linq2go` at the repo root: no subpackages, no `cmd/`, no generated code, no build tooling
beyond the Go toolchain. One dependency, `golang.org/x/exp`, for a single constraint.

The library is **complete and green**: 50 operators, 100% statement coverage, `go vet` and
`go fmt` clean. There is no backlog. Treat a request as an addition to a finished thing, not as
work resumed halfway.

## Commands

```sh
go build ./...
go vet ./...
go test ./...
go test -run '^Test_Where_Slice$' ./...   # single test
go test -race -cover ./...                # the one to run before declaring done
go fmt ./...                              # NOT bare `gofmt` — see below
```

There are no benchmarks in the tree: they are written as throwaway files, measured, then deleted
(see **Performance**).

## Toolchain traps

**Go 1.27+ is mandatory.** The design rests on *generic methods* — methods declaring their own type
parameters:

```go
func (s *slice[T]) Select[NewT any](fn func(T) NewT) *slice[NewT]
```

That is invalid before Go 1.27 and rejected with `method must have no type parameters`. `go.mod`
pins `go 1.27.0` and the `go` command downloads it on its own (`go env GOTOOLCHAIN` reads
`go1.25.0+auto` here, resolving to `~/sdk/go1.27.0`). Never "fix" a generic method by hoisting it
back into a free function — check which toolchain actually ran first.

**Drive every tool through the `go` command**, which resolves the pinned toolchain. A standalone
binary earlier on `PATH` — Homebrew's `/opt/homebrew/bin/gofmt` in particular — may come from an
older Go and will report `method must have no type parameters` on *every* generic method in the
package. That is a stale tool, not a real error. Use `go fmt ./...` or `$(go env GOROOT)/bin/gofmt`.

**The repo path contains a space** (`.../Simone Serra/Linq2Go`). A `replace` directive pointing at
it fails with `malformed module path: empty path element`. To compile a scratch module against this
one, symlink the repo to a space-free path first and point `replace` at the symlink. Simpler still:
prototype *inside* the package in a temporary `zz_*.go` file and delete it after.

## Architecture

Three **unexported** wrapper types, each a single field around a plain Go value:

- `slice[T any]{values []T}` — 32 methods, the bulk of the library
- `dictionary[K comparable, V any]{values map[K]V}` — 11 methods
- `group[K comparable, V any]{values map[K][]V}` — 7 methods; reachable only through
  `slice.Group` / `dictionary.Group`, and there is deliberately no `FromGroup`

Callers never name these types: they enter through a `From*` and get a plain Go value back through
an exit. The `new*` constructors wrap **without copying**.

The flow is a closed loop: **enter** (`from.go`) → **chain** (every intermediate operator returns a
freshly allocated wrapper; receivers are never mutated) → **exit** (`to.go`, `values.go`/`keys.go`,
or a terminal operator in `sum`/`min`/`max`/`avg`/`first`/`last`/`any`/`all`/`contains`/`count`).

A `group` chains only through `Where` (a filter over whole buckets) and exits through `ToMap`,
`Count` and the four aggregations, which return a `*dictionary[K, NewV]` keyed by group.

**Delegate to the primitive instead of duplicating its loop.** This is the pattern that keeps
mirrored code from drifting, and it is used throughout:

| this | delegates to |
|---|---|
| the four `group` aggregations | `newSlice(bucket).Sum/Min/Max/Avg(fn)` |
| `slice.TakeWhile` / `SkipWhile` | `Take` / `Skip` (which own the capacity capping) |
| `slice.WhereMin` / `WhereMax` | `Min` / `Max` (reusing the returned index) |
| `dictionary.Select`, `ChangeKey` | `Transform` |
| `dictionary.Values`, `Keys` | `dictionary.ToSlice` |

The per-bucket `newSlice` costs nothing (measured: escape analysis keeps the wrapper on the stack).
Keep every one of these delegations when editing.

### Where things live

One file per operator **family**, named after the operator, holding **every receiver and every
variant** of it. A genuinely new operator gets a new file; a variant of an existing one joins its
family's file — the test is whether a reader hunting for it would look under that name.

| file | contents |
|---|---|
| `types.go` | the three wrappers, plus `newSlice` / `newDictionary` / `newGroup` |
| `interfaces.go` | the `Number` constraint, nothing else |
| `functions.go` | `Self` and `Equal`, the two ready-made `fn` arguments |
| `from.go` | `FromSlice`, `FromMap`, `FromSeq`, `FromSeq2` |
| `to.go` | every exit: `slice.ToSlice`/`ToMap`/`ToSeq`, `dictionary.ToMap`/`ToSlice`/`ToSeq`, `group.ToMap` |
| `values.go` / `keys.go` | `dictionary.Values`, and the mirror for keys |
| `min.go` / `max.go` | `slice.Min` + `slice.WhereMin` + `group.Min`, and the mirror for max |
| `sum.go` / `avg.go` | the `slice` and `group` version of each |
| `select.go` | `slice.Select`, `slice.SelectMany`, `dictionary.Select` |
| `group.go` | `slice.Group`, `dictionary.Group` |
| `transform.go` | `dictionary.Transform`, `dictionary.ChangeKey` |
| `first.go` / `last.go` | `FirstOrNil` + `FirstOrDefault`, and the mirror for last |
| `orderby.go` / `reverse.go` | `slice.OrderBy` + `slice.OrderByDescending`, and `slice.Reverse` |
| `take.go` / `skip.go` | `slice.Take` + `slice.TakeWhile`, and the mirror for skip |
| `where.go` | `slice.Where`, `dictionary.Where`, `group.Where` |
| `count.go` | `slice.Count`, `dictionary.Count`, `group.Count` |
| the rest | one `slice` method each: `distinct`, `concat`, `except`, `intersect`, `contains`, `any`, `all` |

Every file has a matching `_test.go`, except `types.go` and `interfaces.go`, which hold no
behaviour of their own.

## Invariants an edit must not break

**The copy contract** — the invariant most likely to fall to a well-meaning edit:

- `FromSlice` and `FromMap` are the **only** places that copy. They clone the input, so nothing the
  caller later does to the source can reach the pipeline.
- Everything downstream copies nothing. Operators build fresh results, `Take`/`Skip` return
  sub-slices of the receiver's array — with **capacity capped** (`[:n:n]`, `[n:l:l]`) so that an
  `append` on an exited result reallocates instead of writing into the shared array; do not
  "simplify" those slice expressions — and the `To*` exits hand back the internal storage as-is.
- The exception is the reordering operators, `OrderBy`/`OrderByDescending`/`Reverse`: the stdlib
  sorts and reverses **in place**, so they clone first and reorder the clone. That is not an extra
  defensive copy but the result allocation itself — `Select`'s `make` in another form — and it is
  what stops an in-place reorder from reaching the receiver and every wrapper sharing its array.
- **Never call `FromSlice`/`FromMap` from inside the package.** Use `newSlice`/`newDictionary`/
  `newGroup`: a `From*` call inside an operator clones data that is already private, which is
  exactly what this design exists to avoid. Group aggregation went from 82 KB to 200 B per call by
  removing one such call.
- Accepted consequence, already documented on the `To*` methods: two wrappers derived from the same
  `From*` share one array, so `base.Skip(1).ToSlice()[0] = 42` is visible through `base`. The
  caller's own data is never at risk — only the library's private copy is.

**Never return nil.** Intermediate results start as `result := []T{}` or `make([]T, 0, n)`, `From*`
turns a nil input into an empty wrapper, and `Concat` guards the case where `slices.Concat` returns
nil for two empty inputs. Any result must marshal as `[]`, never `null`, and must be safe to range
over without a nil check. When adding an operator, check every early-return path against this.
Watch for stdlib functions that preserve nilness: `slices.Clone(nil)` is `nil`, which is safe here
only because `From*` guarantees the wrapped slice is never nil.

## Conventions for new code

- **Place the file by the family rule above**, and ship its test in the same change.
- Receivers are named after the wrapper: `s` for `slice`, `d` for `dictionary`, `g` for `group`.
- The selector or predicate is the **last** parameter and is named `fn`.
- **There are exactly two `fn` shapes**, and `functions.go` holds a ready-made argument for each:
  `Self` for a selector `func(T) K`, `Equal` for a comparison `func(T, T) bool`. A new operator
  takes one of those two shapes unless there is a reason not to, so that both keep working.
- Comparison-based operators (`Contains`, `Except`, `Intersect`) take an explicit
  `fn func(T, T) bool` rather than constraining `T` to `comparable`. `Contains` calls it as
  `fn(value, element)` — searched value first — and `Except`/`Intersect` inherit that order as
  `fn(element, value)`, receiver element first. Both directions are pinned by a test.
- **Type parameters follow their role**: `T` is an element or a generic result
  (`dictionary.ToSlice[T]`), `K` and `V` are a key and a value — alone too: `Distinct[K]` and
  `Group[K]` select a key, the aggregations `[V Number]` select a value — and the `New` prefix marks
  a role whose name the receiver already binds (`Select[NewT]`, `Transform[NewK, NewV]`, the group
  aggregations `[NewV Number]`).
- Wrap the standard library where it already does the job (`Any` → `slices.ContainsFunc`,
  `Concat` → `slices.Concat`, the entry copies → `slices.Clone`/`maps.Clone`) instead of
  reimplementing it. There are two measured exceptions, both under **Performance**.
- **Doc comments: a concise description plus one simple usage example**, in the style of
  `where.go`. The description opens with the identifier's name and stays short — it may span lines,
  but keep it tight, not prolix; where behaviour cannot be read off the signature, say it in a
  subordinate clause or after a semicolon — see `Transform` (key collisions), `Except` (duplicates
  kept), `ToSlice` (internal storage returned). The example follows a blank `//` line, indented with
  a tab so godoc renders it as a code block: one chained call with the result in a trailing comment.
  **The result shown must be real** — every example in the package has been executed; verify yours
  the same way (see **Testing**).

## Behaviours that look like bugs and are not

Check this list before "fixing" one of them:

- **Empty input signals itself, except in `FirstOrDefault`/`LastOrDefault`.** `Min`/`Max` return
  `-1` and `0`, `WhereMin`/`WhereMax` return `-1` and `*new(T)`, `Avg` returns `NaN` (the
  arithmetic 0/0), `FirstOrNil`/`LastOrNil` return nil (otherwise a pointer to a *copy*). Only
  `FirstOrDefault` and `LastOrDefault` stay ambiguous by choice, returning a zero value
  indistinguishable from a legitimate result.
- **`NaN` follows `cmp.Compare`, not `slices.Min`/`Max`.** `Min` and `Max` compare with
  `cmp.Compare`, which sorts `NaN` below every number: a single `NaN` wins a minimum and never wins
  a maximum, `OrderBy` puts it first and `OrderByDescending` last. `Sum` and `Avg` do plain
  arithmetic, so one `NaN` poisons both.
- **`Sum` accumulates in `V`**, the type `fn` returns, so a narrow integer overflows silently and
  `Avg` inherits the wrong total.
- **`Transform` and `ChangeKey` are not deterministic on key collisions.** The survivor depends on
  Go's randomized map iteration order and is not stable between runs.
- **`Except` and `Intersect` keep duplicates.** They include or exclude values, they do not build a
  set; chain `Distinct` for that. This diverges from LINQ. They are also **O(n·m)** by construction:
  equality comes from `fn`, so there is no set to hash into. That is the design, not an oversight.
- **`slice.ToMap` keeps the first element** on a key collision, unlike the map operators above;
  `FromSeq2` keeps the **last** pair, as `maps.Collect` does.
- **`dictionary.ToSlice`, `Values`, `Keys` and `dictionary.ToSeq` return an unordered result**,
  being map iterations.

## Performance

Micro-benchmarks in this package lie unless the optimizer is held off. Two traps produced wrong
conclusions here before being caught:

- pass the selector as a **package-level `var sel = func(...)`**, otherwise the compiler inlines the
  operator and devirtualizes `fn`, and the numbers describe code that will not exist at a call site;
- assign the result to a **package-level `sink`**, otherwise escape analysis skips the allocation
  entirely and `-benchmem` reports a win that is not there.

Benchmarks are not kept in the tree. Write one as `zz_*_test.go`, run
`go test -run '^$' -bench . -benchmem -count=5 ./...`, record the number **here**, delete the file.

### The two measured exceptions to "wrap the stdlib"

**`Min`/`Max`/`WhereMin`/`WhereMax` do not use `slices.MinFunc`/`MaxFunc`**, which evaluate `fn`
twice per comparison — 2,3× slower with a costly selector (1337 → 579 µs on 10.000 elements) — while
`slices.Min`/`Max` would need an allocated projection and would change the `NaN` semantics. They are
single-pass loops calling `fn` exactly once per element and returning index and value. Keep the two
loops mirrored.

**`LastOrNil`/`LastOrDefault` use an indexed backwards loop, not `slices.Backward`**: the
range-over-func costs `LastOrNil` 9,6 → 16,0 µs (1,67×) and `LastOrDefault` 9,7 → 22,3 µs (2,3×) on
10.000 elements, at equal allocations. `gopls` will keep suggesting the modernization; this
paragraph is the answer.

### Never return the address of a loop variable

`FirstOrNil`/`LastOrNil` return a pointer to a *copy*, and the obvious `return &v` heap-allocates
one copy **per element scanned**, not per match. Escape analysis marks the *variable*, and a
variable that escapes is allocated where it is **declared** — so with `&v` the allocation sits in
the loop header and runs every iteration. `-gcflags=-m` shows it: `moved to heap: v` on the `for`
line, versus `moved to heap: match` on the line inside the `if` when the copy is made there.
Measured on 10.000 elements with no match: 80 KB and 10.000 allocations, 9,5 → 76 µs (**8×**). The
contract is identical either way, so keep the copy in the branch; the comment in both files says so.

### Known, unmeasured

`OrderBy`/`OrderByDescending` hand `cmp.Compare(fn(a), fn(b))` to `slices.SortStableFunc`, so `fn`
runs **twice per comparison** — the same cost that pushed `Min`/`Max` off `slices.MinFunc`, but here
never measured. With a costly selector a decorated sort (project once, sort the pairs) would fix it.
Measure before assuming it matters, and mind the pre-sizing rule if you try.

### Measured and rejected — do not retry without a reason

- pre-sizing `Distinct`, `Where` and `slice.ToMap`: the output size is unknown there and the upper
  bound can be far off (`Distinct` on 10.000 elements with 100 distinct values: 6.4 KB → 377 KB);
- inlining `Contains` into `Except`/`Intersect` (no gain);
- `slices.Collect` instead of `slices.AppendSeq` in `FromSeq` (identical);
- dropping the per-bucket `newSlice` in the group aggregations (no gain — it never escapes).

Applied and kept: `make` at the exact final size in `Select`, `dictionary.ToSlice`, `Transform` and
the four group aggregations, all cases where the size is known from the input.

## Design questions already settled

Do not reopen these without new information; each cost a real investigation.

**Go cannot add a constraint to the receiver's type parameter.** An operator that exists only when
`T` is comparable — a no-argument `Distinct()` — is not expressible: the method is rejected with
`T does not satisfy comparable`, and a constraint referring to `T` with `term cannot be a type
parameter`. Passing a function is the only way to carry `comparable` in, which is why `Self` and
`Equal` exist. The alternative is a second wrapper type binding the key once
(`keyedSlice[T, K comparable]` embedding `*slice[T]`): it works, but Go has no self type, so the
promoted methods return `*slice[T]` and the chain silently decays — **12 chaining operators would
have to be shadowed** to prevent it, and a `reflect`-based guard would catch only 9 of them, since
generic methods are invisible to `reflect` (19 of 33 methods are visible). It only pays from two or
three key-based operators in one chain; below that `.AsComparable(Self).Distinct()` is longer than
`.Distinct(Self)`.

**Dropped by the maintainer's decision**: SQL-style `NaN` skipping (more complexity than it was
worth) and rewriting `Number` to drop `golang.org/x/exp`. The `slice` parameter name in
`Except`/`Intersect` also stays as it is.

## Testing

Coverage is **100%** (`go test -race -cover`) and should stay there — an uncovered statement is
usually a sign that something was added but not wired up. One `_test.go` per operator file,
table-driven, named `Test_<Operator>[_<Scenario>]`, with `Empty` and `Nil` cases alongside the happy
path. Every table asserts the result is not nil before comparing.

Beyond the happy path, the suite pins what a future edit is most likely to break by accident and
what cannot be read off a signature:

- the copy contract: `Test_FromSlice_CopiesTheInput`, `Test_ToSlice_ReturnsTheInternalStorage`,
  `Test_Take_CapsTheCapacity` and its three siblings;
- the never-nil paths, on every operator;
- the "look like bugs" list above — `NaN` in three operators, `fn` direction in `Contains`,
  `Except` and `Intersect`, first-vs-last on key collisions, `-1` on empty input,
  `FirstOrNil`/`LastOrNil` pointing at a copy, sort stability, receivers not mutated.

Two workflows use temporary files that must not be committed:

- **verifying a doc example**: write `zz_doc_test.go` executing the exact expression from the
  comment, run it, delete it. This is how "the result shown must be real" is actually enforced;
- **benchmarking**: as described under **Performance**.

## Documentation layout

- `README.md` — for users and contributors: description, examples, operator tables, design
  rationale, contributing guidelines. English.
- `CLAUDE.md` — this file: only what is needed to work on the code here. English.

A change in library behaviour usually touches four places: the doc comment, its test, README's
"Behaviour worth knowing", and the list above. Keep them in sync — a stale line here is worse than
no line, because the next session will trust it.
