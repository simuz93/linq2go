# Security Policy

## Reporting a vulnerability

Please do not open a public issue for a security problem. Report it privately through GitHub's
[private vulnerability reporting](https://github.com/simuz93/linq2go/security/advisories/new):
describe the problem, how to reproduce it and what impact you see.

Once the problem is confirmed, a fix is published as a new tagged release and the advisory is disclosed with it.

## Supported versions

Only the latest tagged release receives security fixes.

## Scope

linq2go is a pure in-memory library over Go slices and maps: it performs no I/O, parses no external
input and has a single dependency (`golang.org/x/exp`). Most reports are therefore expected to be
about that dependency or about Go itself, which are best reported upstream.
