# Bump klauspost/compress to v1.18.7 (CVE-2026-63209)

- **Date:** 2026-11-04
- **Author:** Nitish Bhat
- **Related PR(s):** TBD
- **Related issue(s) / JIRA:** NO-JIRA (security remediation)

## Context

A trivy image scan (`trivy image --ignore-unfixed --severity HIGH,CRITICAL`)
of the images built from `main` flagged `github.com/klauspost/compress`
v1.17.9, CVE-2026-63209 (HIGH), fixed in v1.18.7, in
`usr/local/bin/manager` and `usr/local/bin/webhook-server`. It is an indirect
dependency that is linked into those shipped binaries.

`vendor/` on `main` was also stale relative to `go.mod`: since #89 `go.mod`
requires `x/net` v0.60.0, `x/text` v0.42.0, `x/sys` v0.48.0, `x/sync` v0.23.0,
`x/term` v0.46.0 and `x/tools` v0.49.0, but `vendor/modules.txt` still held the
versions from #88. Re-vendoring for any bump therefore also pulls in that
backlog.

## Approach

Two commits so each diff is reviewable on its own:

1. `chore: re-sync vendor/ and go.sum with go.mod` — `go mod tidy && go mod
   vendor` against the existing `go.mod`, no version changes. `go.mod` is
   untouched; `go.sum` only loses stale lines for the old `x/*` versions, and
   `vendor/` catches up to what `go.mod` already requires.
2. `fix(security): bump klauspost/compress to v1.18.7` — `go get
   github.com/klauspost/compress@v1.18.7`, `go mod tidy`, `go mod vendor`. The
   diff is one `go.mod` line, two `go.sum` lines, the vendored library files
   and its `vendor/modules.txt` entry; no other module moves.

Pinned to the minimum patched release, not latest. Both steps were run inside
a `golang:1.26.10` container.

### Alternatives considered

- Editing only `go.mod` / `go.sum` and leaving `vendor/` stale — rejected; the
  repo vendors its dependencies and the earlier bot PRs updated `vendor/`.

## Scope

- **In scope:** `klauspost/compress`, plus the vendor re-sync needed to keep
  the bump diff minimal.
- **Out of scope:** the Go toolchain bump (separate PR) and any other module.

## Validation

- `go build ./...` succeeds after each step with the regenerated `vendor/`.
- Rebuilt all four images with `make docker-build-all`; `trivy image` on the
  operator and webhook-server images shows the `klauspost/compress` finding
  gone (before: 1 in each of `manager` and `webhook-server`; after: 0). The
  only finding left in those images is the Go `stdlib` one, fixed separately
  by the Go 1.26.10 bump PR (this branch is based on `main`, which is still on
  Go 1.26.9).
- Unit tests were not run for this change.

## Risks and rollback

- Known risks: a minor bump of a compression library used by the Kubernetes
  client stack, plus the vendor catch-up for the `x/*` modules that `go.mod`
  already required; low risk since `go build` passes.
- Rollback plan: `git revert` the two commits.
