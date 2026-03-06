# Typo in Package Name: `convertinteranl`

## Status: ✅ RESOLVED

## Issue Summary

The internal package `internal/convertinteranl/` had a typo — it should be `convertinternal`. Since Go module paths are part of the public API (even for internal packages within the module), this created a permanent misspelling in import paths.

## Root Cause Analysis

Likely a typo during initial creation that was never caught.

## Fix Description

1. Created `internal/convertinternal/` with correct spelling.
2. Moved all code from `convertinteranl/` to `convertinternal/`.
3. Updated all import paths across the module.
4. Deleted the old `convertinteranl/` directory.
5. Module-internal change — no external consumers affected.

## Prevention and Non-Regression

- Add a linting rule or code review checklist for package naming.
- Review all package names for typos before any public release.

## Done Checklist

- [x] Fix applied
- [x] Tests pass
- [x] No remaining references to old name
