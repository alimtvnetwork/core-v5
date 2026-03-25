# Issue: corestrtests — Test_Seg4_SS_TakeDynamic type assertion panic

## Date: 2026-03-25

## Status: ✅ FIXED

## Failing Test

`Test_Seg4_SS_TakeDynamic` — panics with:
```
panic: interface conversion: interface {} is []string, not corestr.SimpleSlice
```

## Root Cause

The test at `Coverage38_Seg4_SimpleSlice_test.go:287` performs:
```go
s := corestr.SimpleSlice{"a", "b"}
actual := args.Map{"len": len(s.TakeDynamic(10).(corestr.SimpleSlice))}
```

But `SimpleSlice.TakeDynamic` (SimpleSlice.go:333-343) returns `[]string`, not `SimpleSlice`:
```go
func (it *SimpleSlice) TakeDynamic(takeDynamicItems int) any {
    if takeDynamicItems >= it.Length() {
        return []string(*it)   // ← returns []string
    }
    return []string((*it)[:takeDynamicItems])  // ← also []string
}
```

Both branches explicitly cast to `[]string`. The type assertion `.(corestr.SimpleSlice)` panics
because `[]string` is not the same named type as `corestr.SimpleSlice`.

## Fix

Change the type assertion from `.(corestr.SimpleSlice)` to `.([]string)`.

## Affected Files

- `tests/integratedtests/corestrtests/Coverage38_Seg4_SimpleSlice_test.go:287` — **fix location**
- `coredata/corestr/SimpleSlice.go:333-343` — production code (correct, no change needed)

## Learning

`TakeDynamic` returns `any` with concrete type `[]string`. Even though `SimpleSlice` is
defined as `type SimpleSlice []string`, Go treats them as distinct types for assertions.
