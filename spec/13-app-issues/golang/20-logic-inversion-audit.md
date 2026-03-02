# Codebase-Wide Bug Audit: Recursion, Logic Inversion, and Semantic Errors

## Date: 2026-03-02

## Scan Scope
- 402 Go source files (excluding tests, vendor, node_modules)
- Patterns checked: self-calling methods, mutual recursion, deprecated wrappers calling themselves, inverted logic (`Has*` using `Is.Null`), `Count`/`Length` mismatches, `IsSafe*` returning unsafe values

## Bugs Found and Fixed

### Bug 1: `LeftRight.HasLeft()` / `HasRight()` — Inverted Logic
**File:** `coredata/coredynamic/LeftRight.go` lines 22-30
**Severity:** HIGH — returns `true` when field is `nil`, `false` when present
**Root cause:** Missing `!` negation — used `Is.Null(it.Left)` instead of `!Is.Null(it.Left)`
**Impact:** Any caller checking `HasLeft()` / `HasRight()` gets opposite results. `IsLeftEmpty()` and `IsRightEmpty()` were correct (used `Is.Null` without negation in a `return it == nil ||` pattern).
**Fix:** Added `!` before `reflectinternal.Is.Null(...)` in both methods.

### Bug 2: `LinesValidators.Count()` — Off-By-One
**File:** `corevalidator/LinesValidators.go` line 41-43
**Severity:** MEDIUM — `Count()` returns `Length() - 1` (LastIndex) instead of `Length()`
**Root cause:** Copy-paste error — `Count()` delegates to `LastIndex()` instead of `Length()`
**Impact:** Any caller using `Count()` gets one fewer than actual. Every other type in the codebase correctly has `Count() { return it.Length() }`.
**Fix:** Changed to `return it.Length()`.

### Bug 3: `Attributes.IsSafeValid()` — Inverted Logic
**File:** `coredata/corepayload/AttributesGetters.go` line 74-76
**Severity:** HIGH — returns `true` when attributes have issues, `false` when safe
**Root cause:** Missing `!` negation — `IsSafeValid` should be the opposite of `HasIssuesOrEmpty`
**Impact:** Any caller checking `IsSafeValid()` gets inverted safety check.
**Fix:** Changed to `return !it.HasIssuesOrEmpty()`.

## Patterns Verified Clean (No Issues Found)
- **Self-calling methods**: No method directly calls itself (AllValidVersionValues was already fixed)
- **Deprecated wrappers**: All correctly delegate to their replacement functions
- **Count/Length aliases**: All other types correctly use `Count() { return it.Length() }`
- **String() methods**: No infinite recursion via String() → JsonString() → String() cycles
- **Error/CompiledError**: No circular delegation
- **Serialize/Deserialize**: No self-referencing patterns
- **AllValidVersionValues**: Confirmed fixed (calls `AllVersionValues()` correctly)

## Previously Fixed (from earlier audits)
- `AllValidVersionValues` infinite recursion (spec/13-app-issues/golang/18-recursion-audit.md)
- `AuthInfo.Clone()` missing `Identifier` field (spec/13-app-issues/golang/19-value-receiver-corepayload.md)
