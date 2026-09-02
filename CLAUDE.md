# CLAUDE.md

Working notes for Claude Code on this repository: what an agent needs in order to edit this code
correctly. Anything addressed to someone *using* or *contributing to* the library — what it does,
examples, design rationale, contribution guidelines — belongs in `README.md`, not here.

## Project

`github.com/simuz93/linq2go` — LINQ-style fluent queries over Go slices and maps. A single flat
package `linq2go` at the repo root: no subpackages, no `cmd/`, no generated code, no build tooling
beyond the Go toolchain. One dependency, `golang.org/x/exp`, for a single constraint.

## Commands

```sh
go build ./...
go vet ./...
go test ./...                        # no test files exist yet
go test -run '^TestWhere$' ./...     # single test
go test -race -cover ./...
go test -bench . -benchmem ./...
go fmt ./...                         # NOT bare `gofmt` — see below
```

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
one, symlink the repo to a space-free path first and point `replace` at the symlink.

## Where things live

One operator per file, named after the operator, and a file holds **every receiver** of that
operator.

| file | contents |
|---|---|
| `types.go` | the three wrappers, plus `newSlice` / `newDictionary` / `newGroup` |
| `interfaces.go` | the `Number` constraint, nothing else |
| `from.go` | `FromSlice`, `FromMap`, `FromSeq` |
| `to.go` | `slice.ToSlice`, `slice.ToMap`, `dictionary.ToMap`, `dictionary.ToSlice`, `group.ToMap` |
| `values.go` | `dictionary.Values` |
| `min.go` / `max.go` | `slice.Min` + `slice.WhereMin` + `group.Min`, and the mirror for max |
| `sum.go` / `avg.go` | the `slice` and `group` version of each |
| `select.go` | `slice.Select`, `slice.SelectMany`, `dictionary.Select` |
| `group.go` | `slice.Group`, `dictionary.Group` |
| `transform.go` | `dictionary.Transform`, `dictionary.ChangeKey` |
| `first.go` | `FirstOrNil`, `FirstOrDefault` |
| the rest | one `slice` method each: `where`, `distinct`, `take`, `skip`, `concat`, `except`, `intersect`, `contains`, `any`, `all` |

Three **unexported** wrapper types, each a single field around a plain Go value:

- `slice[T]{values []T}`
- `dictionary[K comparable, V any]{values map[K]V}`
- `group[K comparable, V any]{values map[K][]V}` — reachable only through `slice.Group` /
  `dictionary.Group`; there is deliberately no `FromGroup`

Callers never name these types: they enter through a `From*` and get a plain Go value back through
an exit. The `new*` constructors wrap **without copying**.

The flow is a closed loop: **enter** (`from.go`) → **chain** (every intermediate operator returns a
freshly allocated wrapper; receivers are never mutated) → **exit** (`to.go`, `values.go`, or a
terminal operator in `sum`/`min`/`max`/`avg`/`first`/`any`/`all`/`contains`).

A `group` exits only through `ToMap` and the four aggregations, which return a `*dictionary[K, NewV]`
keyed by group. They re-wrap each bucket with `newSlice(v)` and delegate to the `slice` version —
keep that delegation when adding a group operator instead of duplicating the loop body. It costs
nothing (measured: escape analysis keeps the wrapper on the stack) and it is what stops `group.Min`
and `slice.Min` from drifting apart.

## Invariants an edit must not break

**The copy contract** — the invariant most likely to fall to a well-meaning edit:

- `FromSlice` and `FromMap` are the **only** places that copy. They clone the input, so nothing the
  caller later does to the source can reach the pipeline.
- Everything downstream copies nothing. Operators build fresh results, `Take`/`Skip` return
  sub-slices of the receiver's array — with **capacity capped** (`[:n:n]`, `[n:l:l]`) so that an
  `append` on an exited result reallocates instead of writing into the shared array; do not
  "simplify" those slice expressions — and the `To*` exits hand back the internal storage as-is.
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

## Conventions for new code

- **A new operator is a new file**, never an append to an existing one.
- The selector or predicate is the **last** parameter and is named `fn`. Comparison-based operators
  (`Contains`, `Except`, `Intersect`) take an explicit `fn func(T, T) bool` rather than constraining
  `T` to `comparable`. `Contains` calls it as `fn(value, element)` — searched value first, and
  `Except`/`Intersect` inherit that order as `fn(element, value)`, receiver element first.
- **Type parameters follow their role**: `T` is an element or a generic result
  (`dictionary.ToSlice[T]`), `K` and `V` are a key and a value — alone too: `Distinct[K]` and
  `Group[K]` select a key, the aggregations `[V Number]` select a value — and the `New` prefix marks
  a role whose name the receiver already binds (`Select[NewT]`, `Transform[NewK, NewV]`, the group
  aggregations `[NewV Number]`).
- Wrap the standard library where it already does the job (`Any` → `slices.ContainsFunc`,
  `WhereMin`/`WhereMax` → `slices.MinFunc`/`MaxFunc`, `Concat` → `slices.Concat`, the entry copies →
  `slices.Clone`/`maps.Clone`) instead of reimplementing it.
- **Doc comments: a concise description plus one simple usage example**, in the style of
  `where.go`. The description opens with the identifier's name and stays short — it may span lines,
  but keep it tight, not prolix; where behaviour cannot be read off the signature, say it in a
  subordinate clause or after a semicolon — see `Transform` (key collisions), `Except` (duplicates
  kept), `ToSlice` (internal storage returned). The example follows a blank `//` line, indented with
  a tab so godoc renders it as a code block: one chained call with the result in a trailing comment.
  **The result shown must be real** — every example in the package has been executed; verify yours
  the same way.

## Behaviours that look like bugs and are not

Check this list before "fixing" one of them:

- **Empty input yields the zero value.** `Min`/`Max`/`Avg` return `0`, `WhereMin`/`WhereMax` return
  `*new(T)`, `FirstOrDefault` returns the zero value. The library has no way to say "there is
  nothing"; `FirstOrNil` is the one exception, and it returns a pointer to a *copy*.
- **`NaN` follows `cmp.Compare`, not `slices.Min`/`Max`.** `WhereMin`/`WhereMax` compare through
  `slices.MinFunc`/`MaxFunc`, and `cmp.Compare` sorts `NaN` below every number: a single `NaN` wins
  a minimum and never wins a maximum. `Sum` and `Avg` do plain arithmetic, so one `NaN` poisons both.
- **`Sum` accumulates in `V`**, the type `fn` returns, so a narrow integer overflows silently and
  `Avg` inherits the wrong total.
- **`Transform` and `ChangeKey` are not deterministic on key collisions.** The survivor depends on
  Go's randomized map iteration order and is not stable between runs.
- **`Except` and `Intersect` keep duplicates.** They include or exclude values, they do not build a
  set; chain `Distinct` for that. This diverges from LINQ.
- **`slice.ToMap` keeps the first element** on a key collision, unlike the map operators above.
- **`dictionary.ToSlice` and `Values` return an unordered result**, being map iterations.

## Benchmarking

Micro-benchmarks in this package lie unless the optimizer is held off. Two traps produced wrong
conclusions here before being caught:

- pass the selector as a **package-level `var sel = func(...)`**, otherwise the compiler inlines the
  operator and devirtualizes `fn`, and the numbers describe code that will not exist at a call site;
- assign the result to a **package-level `sink`**, otherwise escape analysis skips the allocation
  entirely and `-benchmem` reports a win that is not there.

Already measured and **rejected** — do not retry without a reason:

- pre-sizing `Distinct`, `Where` and `slice.ToMap`: the output size is unknown there and the upper
  bound can be far off (`Distinct` on 10.000 elements with 100 distinct values: 6.4 KB → 377 KB);
- inlining `Contains` into `Except`/`Intersect` (no gain);
- `slices.Collect` instead of `slices.AppendSeq` in `FromSeq` (identical);
- dropping the per-bucket `newSlice` in the group aggregations (no gain — it never escapes).

Applied and kept: `make` at the exact final size in `Select`, `dictionary.ToSlice`, `Transform` and
the four group aggregations, all cases where the size is known from the input.

Two directions were explored and then dropped **by the maintainer's decision** — do not reopen them
as improvements: SQL-style `NaN` skipping (more complexity than it was worth) and rewriting `Number`
to drop `golang.org/x/exp`.

## Testing

No test files exist; coverage is 0%. The first ones should pin the copy contract, the never-nil
paths, and the "look like bugs" list above — those are what a future edit is most likely to break by
accident, and none of them can be read off a signature.

## Documentation layout

- `README.md` — for users and contributors: description, examples, operator tables, design
  rationale, contributing guidelines. English.
- `CLAUDE.md` — this file: only what is needed to work on the code here. English.

A change in library behaviour usually touches three places: the doc comment, README's "Behaviour
worth knowing", and the list above. Keep them in sync.
