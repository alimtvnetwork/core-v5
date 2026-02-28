# Typo in Package Name: `convertinteranl`

## Issue Summary

The internal package `internal/convertinteranl/` has a typo — it should be `convertinternal`. Since Go module paths are part of the public API (even for internal packages within the module), this creates a permanent misspelling in import paths.

## Root Cause Analysis

Likely a typo during initial creation that was never caught.

## Fix Description

1. Create a new package `internal/convertinternal/` with correct spelling.
2. Move all code from `convertinteranl/` to `convertinternal/`.
3. Update all import paths across the module.
4. Delete the old `convertinteranl/` directory.
5. This is a module-internal change so it won't break external consumers.

## Prevention and Non-Regression

- Add a linting rule or code review checklist for package naming.
- Review all package names for typos before any public release.

## TODO and Follow-Ups

- [ ] Create `internal/convertinternal/`
- [ ] Migrate code
- [ ] Update all imports
- [ ] Delete old directory
- [ ] Verify tests pass

## Done Checklist

- [ ] Fix applied
- [ ] Tests pass
- [ ] No remaining references to old name
