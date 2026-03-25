# Issue: Collection.AddFuncResult panics on nil function pointer

## Date: 2026-03-25

## Status: ✅ FIXED

## Failing Test

- `Test_Cov42_Collection_AddFuncResult_Nil` (nil pointer dereference panic)

## Root Cause

`Collection.AddFuncResult()` in `coredata/corestr/Collection.go` checked
`getterFunctions == nil` (the variadic slice), but did NOT check individual
function pointers for nil:

```go
func (it *Collection) AddFuncResult(getterFunctions ...func() string) *Collection {
    if getterFunctions == nil { return it }  // only guards nil slice
    for _, getterFunc := range getterFunctions {
        item := getterFunc()  // PANIC if getterFunc is nil
    }
}
```

Calling `col.AddFuncResult(nil)` passes `[]func() string{nil}` — the slice
is non-nil but contains a nil function pointer.

## Solution

Add nil guard for individual function pointers:

```go
for _, getterFunc := range getterFunctions {
    if getterFunc == nil { continue }
    item := getterFunc()
    items = append(items, item)
}
```

## Affected Files

- `coredata/corestr/Collection.go` (line ~1804) — production fix
