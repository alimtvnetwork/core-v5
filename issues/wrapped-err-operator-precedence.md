# Issue: WrappedErr.HasErrorOrException Operator Precedence Bug (FIXED)

## Summary
`HasErrorOrException()` panicked on nil receiver due to missing parentheses in boolean expression.

## Root Cause
```go
// BEFORE (broken):
func (it *WrappedErr) HasErrorOrException() bool {
    return it != nil &&
        it.HasError ||
        it.HasThrown  // ← evaluated even when it == nil
}
```

Go operator precedence: `&&` binds tighter than `||`, so this evaluates as:
`(it != nil && it.HasError) || it.HasThrown`

When `it == nil`, `it.HasThrown` still executes → nil pointer dereference.

## Fix
```go
func (it *WrappedErr) HasErrorOrException() bool {
    return it != nil &&
        (it.HasError ||
            it.HasThrown)
}
```

## Affected Files
- `internal/trydo/WrappedErr.go` — line 45-49

## Learning
- Always parenthesize mixed `&&`/`||` expressions in nil-guard patterns
- Never assume operator precedence is obvious — be explicit

## What Not To Repeat
- Writing `a != nil && b || c` without parentheses — this is a classic Go gotcha
