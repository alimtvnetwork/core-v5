# Go Version Outdated (1.17.8)

## Issue Summary

The project is pinned to Go 1.17.8, which is 5+ years old. This prevents using generics, `any` keyword, `errors.Join`, improved stdlib, and modern tooling.

## Root Cause Analysis

Project started on Go 1.17 and was never upgraded.

## Fix Description

See [Go Modernization Plan](/spec/01-app/11-go-modernization.md) for full details.

## Prevention and Non-Regression

- Set up Dependabot or Renovate for Go version tracking.
- Add CI checks for minimum Go version.

## TODO and Follow-Ups

- [ ] Update `go.mod` to Go 1.22+
- [ ] Update `makefile` GoVersion
- [ ] Update README prerequisites
- [ ] Run `go mod tidy`
- [ ] Fix any compilation issues
- [ ] Verify all tests pass

## Done Checklist

- [ ] Go version updated
- [ ] All tests pass on new version
