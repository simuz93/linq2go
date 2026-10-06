# Contributing

Contributions are welcome. This file covers the mechanics of getting a change in; the coding
guidelines a change is held to are in the [Contributing section of the README](README.md#contributing).

## Before writing code

Open an issue first for anything beyond a small fix: a new operator, a behaviour change, a
performance claim. The library is deliberately small and a few design questions are already
settled, so a short discussion up front saves a pull request that cannot be merged.

## Getting a change in

1. Fork the repository and create a branch from `main`.
2. Make the change, with its tests, and check locally that everything passes:

   ```sh
   go build ./...
   go vet ./...
   go test -race -cover ./...
   go fmt ./...
   ```

   Drive every tool through the `go` command: it resolves the toolchain pinned in `go.mod`
   (Go 1.27+, required by the generic methods). A standalone `gofmt` earlier on your `PATH` may be
   older and report `method must have no type parameters` on every generic method.
3. Open a pull request against `main`. Describe what changes and why; the template asks for the
   points the review will look at.
4. CI runs build, vet, gofmt and the test suite with the race detector. A pull request cannot be
   merged until it is green and has been approved by the maintainer, who is requested as reviewer
   automatically.

Pull requests are squash-merged or rebased; keep the branch focused on one change and the title
descriptive, since it becomes the commit message on `main`.

## What a pull request must have

- Tests for every change, with coverage staying at 100%.
- A doc comment on every new or changed exported identifier, with an example whose shown result
  has actually been executed.
- A benchmark run before and after for any change justified by performance.
- README's operator tables and "Behaviour worth knowing" updated when behaviour changes.

## Reporting bugs

Open an issue with the smallest `From*(...)...To*()` chain that reproduces it, the result you got
and the one you expected, and the output of `go version`.
