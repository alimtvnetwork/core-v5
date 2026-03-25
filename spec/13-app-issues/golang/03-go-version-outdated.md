# Go Version Outdated (1.17.8)

## Status: ✅ RESOLVED

## Issue Summary

The project was pinned to Go 1.17.8, which is 5+ years old. This prevented using generics, `any` keyword, `errors.Join`, improved stdlib, and modern tooling.

## Root Cause Analysis

Project started on Go 1.17 and was never upgraded.

## Fix Description

Updated to Go 1.24.0 — see [Go Modernization Plan](/spec/01-app/11-go-modernization.md).

## Prevention and Non-Regression

- Set up Dependabot or Renovate for Go version tracking.
- Add CI checks for minimum Go version.

## Done Checklist

- [x] `go.mod` updated to Go 1.24.0
- [x] `makefile` GoVersion updated
- [x] README prerequisites updated
- [x] `go mod tidy` run
- [x] All tests pass on new version
