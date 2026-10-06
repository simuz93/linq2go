<!-- Thanks for contributing. CONTRIBUTING.md has the mechanics, README.md the guidelines. -->

## What and why

<!-- What changes, and the reason. Link the issue if there is one: "Closes #12". -->

## Checklist

- [ ] `go build ./...`, `go vet ./...`, `go fmt ./...` and `go test -race -cover ./...` pass locally
- [ ] Every change has tests, with `Empty` and `Nil` cases, and coverage stays at 100%
- [ ] New or changed exported identifiers have a doc comment with an example whose shown result was actually executed
- [ ] A new operator is in its own file, a variant joins its family's file, and the test file is alongside
- [ ] No `FromSlice`/`FromMap` call inside the package, and no path can return `nil`
- [ ] Performance claims come with a before/after `-benchmem` run (numbers in the description)
- [ ] README's operator tables and "Behaviour worth knowing" are updated if behaviour changes
