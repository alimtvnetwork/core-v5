# corevalidator — Improvement Plan

## Status: 🔧 PARTIAL — Item #1 FIXED, Items #2-7 Open

## Date: 2026-03-03

## Summary

This document outlines improvement opportunities for the `corevalidator` package, covering bug fixes, new capabilities (particularly wildcard/relative path matching), refactoring, and API enhancements.

---

## 1. Bug Fix — `LineValidator.AllVerifyError` Early Return

**Priority:** HIGH
**File:** `corevalidator/LineValidator.go:131-160`

### Problem

`AllVerifyError` declares `var sliceErr []string` but never appends to it. Instead, it returns immediately on the first error (`return err` at line 154), making it behave identically to `VerifyFirstError`.

### Current (broken)

```go
func (it *LineValidator) AllVerifyError(...) error {
    var sliceErr []string

    for _, textWithLine := range contentsWithLine {
        err := it.VerifyError(...)
        if err != nil {
            return err  // BUG: returns immediately, ignores sliceErr
        }
    }

    return errcore.SliceToError(sliceErr) // always empty
}
```

### Fix

```go
func (it *LineValidator) AllVerifyError(...) error {
    var sliceErr []string

    for _, textWithLine := range contentsWithLine {
        err := it.VerifyError(...)
        if err != nil {
            sliceErr = append(sliceErr, err.Error())  // collect, don't return
        }
    }

    return errcore.SliceToError(sliceErr)
}
```

### Impact

Any caller expecting `AllVerifyError` to report ALL failures (not just the first) currently gets incomplete error output. The `VerifyMany(isContinueOnError=true, ...)` path is affected.

---

## 2. Wildcard / Relative Path Matching

**Priority:** HIGH — User-requested feature
**Effort:** MEDIUM

### Problem Statement

Currently, validators compare fixed, exact strings. When validating file paths, CLI output, or generated code, the output often contains **dynamic segments** (timestamps, hashes, absolute path prefixes, UUIDs) that change between runs. Users need to match the **meaningful/relative part** while ignoring the variable prefix/suffix.

### Proposed Solutions

#### Option A: Glob Pattern Matching (Recommended)

Add a new `stringcompareas.Variant` called `Glob` that uses Go's `filepath.Match` or `path.Match` semantics:

```
/home/*/projects/myapp/src/main.go   → matches any username
build/output-*/result.json           → matches any timestamped dir
**/src/main.go                       → matches any prefix path
```

**Implementation steps:**

1. Add `Glob` and `NonGlob` variants to `enums/stringcompareas/Variant.go`
2. Create `isGlobFunc.go` in `enums/stringcompareas/` using `filepath.Match` or a richer glob library
3. Register in the function map
4. No changes needed in `corevalidator` — it already delegates to `stringcompareas.IsCompareSuccess`

**Pros:** Simple, familiar to developers, built-in Go support
**Cons:** `filepath.Match` doesn't support `**` (double-star recursive). May need a third-party glob library for full power.

#### Option B: Segment-Based Path Validator

A new `PathValidator` type that splits paths by separator and validates individual segments:

```go
type PathSegmentValidator struct {
    Segments []TextValidator  // one validator per path segment
    Separator string          // default "/"
}
```

Example usage:

```go
pv := PathSegmentValidator{
    Separator: "/",
    Segments: []TextValidator{
        {Search: "*", SearchAs: stringcompareas.Anywhere},   // any first segment
        {Search: "projects", SearchAs: stringcompareas.Equal},
        {Search: "myapp", SearchAs: stringcompareas.Equal},
        {Search: "src", SearchAs: stringcompareas.Equal},
        {Search: ".go", SearchAs: stringcompareas.EndsWith},
    },
}
```

**Pros:** Maximum flexibility per-segment, reuses existing TextValidator
**Cons:** More verbose, new type to maintain

#### Option C: Template Placeholder Matching

Support `{{*}}` or `${*}` placeholders in search strings that match any content:

```go
v := TextValidator{
    Search:   "/home/{{*}}/projects/myapp/{{*}}.go",
    SearchAs: stringcompareas.Template,  // new variant
}
// Matches: /home/john/projects/myapp/main.go
// Matches: /home/ci-user/projects/myapp/utils.go
```

**Pros:** Intuitive, single-string definition
**Cons:** New parsing logic, potential escaping issues

### Recommendation

**Start with Option A (Glob)** — it's the lowest-effort, highest-value change. It integrates cleanly into the existing `stringcompareas` system and requires zero changes to `corevalidator` itself. If deeper per-segment control is needed later, Option B can layer on top.

---

## 3. SliceValidator Refactoring

**Priority:** MEDIUM
**File:** `corevalidator/SliceValidator.go` (634 lines)

### Problem

`SliceValidator.go` is the largest file in the package at 634 lines. It mixes:

- Core comparison logic (`isValidLines`, `initialVerifyError`, `isLengthOkay`)
- Multiple verify strategies (`AllVerifyError`, `VerifyFirstError`, `AllVerifyErrorUptoLength`, `AllVerifyErrorExceptLast`)
- Error formatting (`UserInputsMergeWithError`, `ActualInputMessage`, `UserExpectingMessage`, `ActualInputWithExpectingMessage`)
- Test helpers (`AssertAllQuick`, `AllVerifyErrorTestCase`, `AllVerifyErrorQuick`)
- Constructors (`NewSliceValidatorUsingErr`, `NewSliceValidatorUsingAny`)

### Proposed Split

| New File | Contents | Lines (est.) |
|----------|----------|-------------|
| `SliceValidator.go` | Core struct, `IsValid`, `isValidLines`, `isLengthOkay`, `SetActual`, `Dispose` | ~150 |
| `SliceValidatorVerify.go` | All `Verify*` and `AllVerifyError*` methods | ~200 |
| `SliceValidatorMessages.go` | `UserInputsMergeWithError`, `ActualInputMessage`, `UserExpectingMessage`, `ActualInputWithExpectingMessage`, `ActualLinesString`, `ExpectingLinesString` | ~100 |
| `SliceValidatorConstructors.go` | `NewSliceValidatorUsingErr`, `NewSliceValidatorUsingAny` | ~60 |
| `SliceValidatorAssert.go` | `AssertAllQuick`, `AllVerifyErrorTestCase`, `AllVerifyErrorQuick` | ~80 |

### Benefit

Follows the project's "one file per function" design principle. Each file is focused and under 200 lines.

---

## 4. Diff-Style Error Output

**Priority:** MEDIUM
**Effort:** MEDIUM

### Problem

When slice validation fails, the current error output shows individual line mismatches but doesn't provide a unified diff view. For large outputs (20+ lines), it's hard to spot what changed.

### Proposal

Add an optional diff-style error formatter:

```go
type Parameter struct {
    // ... existing fields
    IsDiffOutput bool  // When true, produce unified diff format
}
```

Output example:

```diff
--- Expected
+++ Actual
@@ -3,4 +3,4 @@
 line 2
 line 3
-expected line 4
+actual line 4
 line 5
```

### Implementation

- Add a `diffLines(expected, actual []string) string` helper
- Integrate into `AllVerifyErrorUptoLength` when `params.IsDiffOutput` is true
- Minimal, no external dependencies (simple LCS or line-by-line diff)

---

## 5. Partial Slice Matching (Contains-Style for Slices)

**Priority:** MEDIUM
**Effort:** LOW

### Problem

`SliceValidator` requires `len(actual) == len(expected)`. Sometimes you only care that certain lines **exist somewhere** in the output, regardless of order or extra lines.

### Proposal

Add a `SliceContainsValidator` or a flag on `SliceValidator`:

```go
type SliceValidator struct {
    // ... existing fields
    IsSubsetMatch bool  // When true, every expected line must appear in actual (order-independent)
}
```

This enables:

```go
sv := SliceValidator{
    ActualLines:   []string{"a", "b", "c", "d", "e"},
    ExpectedLines: []string{"b", "d"},       // just check these exist
    CompareAs:     stringcompareas.Equal,
    IsSubsetMatch: true,
}
sv.IsValid(true) // true — "b" and "d" both found
```

---

## 6. Validation Chain / Builder API

**Priority:** LOW
**Effort:** MEDIUM

### Problem

Setting up validators requires manually constructing structs with many fields. A builder/chain API would be more ergonomic.

### Proposal

```go
// Fluent builder
v := corevalidator.NewTextMatch("expected output").
    As(stringcompareas.Contains).
    WithTrim().
    CaseInsensitive()

// Slice builder
sv := corevalidator.NewSliceMatch(expectedLines).
    As(stringcompareas.Equal).
    WithTrim().
    VerifyAll(params, actualLines...)
```

### Benefit

Reduces boilerplate in test setup, self-documenting, harder to misconfigure.

---

## 7. Additional Enhancements

### 7a. LineValidator.AllVerifyError Test Coverage

After fixing the bug in §1, add tests that verify multiple errors are actually collected:

```go
func Test_LineValidator_AllVerifyError_CollectsMultiple(t *testing.T) {
    // Arrange: 3 items, all fail
    // Assert: error message contains all 3 failures
}
```

### 7b. Nil Guard Consistency Audit

Some methods check `it == nil`, others don't. Standardize:

- All pointer-receiver methods that could be called on nil should have nil guards
- Document which methods are nil-safe in the README

### 7c. SearchTextFinalized Thread Safety

`SearchTextFinalizedPtr()` lazily initializes `searchTextFinalized` without synchronization. In concurrent test execution, this could cause a race. Consider:

- Using `sync.Once` for the lazy init
- Or documenting that `TextValidator` is not safe for concurrent use

---

## Implementation Priority

| # | Item | Priority | Effort | Dependency |
|---|------|----------|--------|------------|
| 1 | Fix `LineValidator.AllVerifyError` bug | HIGH | LOW | None |
| 2 | Add Glob pattern matching (`stringcompareas.Glob`) | HIGH | MEDIUM | None |
| 3 | Split `SliceValidator.go` into focused files | MEDIUM | LOW | None |
| 4 | Diff-style error output | MEDIUM | MEDIUM | None |
| 5 | Partial slice matching (`IsSubsetMatch`) | MEDIUM | LOW | None |
| 6 | Builder API | LOW | MEDIUM | None |
| 7a | Bug fix test coverage | HIGH | LOW | #1 |
| 7b | Nil guard audit | LOW | LOW | None |
| 7c | Thread safety for SearchTextFinalized | LOW | LOW | None |

---

## Decision Log

| Date | Decision | Rationale |
|------|----------|-----------|
| 2026-03-03 | Created improvement plan | User requested path/wildcard validation and general improvement review |
| — | Glob recommended over Template/Segment for path matching | Lowest effort, integrates at `stringcompareas` level, zero changes to `corevalidator` |
| — | `LineValidator.AllVerifyError` bug confirmed | Line 154 returns early instead of appending to `sliceErr` |
