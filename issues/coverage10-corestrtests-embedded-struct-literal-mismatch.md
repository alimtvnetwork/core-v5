# Issue: Coverage10 corestrtests — Embedded Struct Literal & Missing Method Mismatch

## Date: 2026-03-16

## Summary

`corestrtests/Coverage10_test.go` failed to compile with 11 errors, blocking the entire package from coverage.

## Root Cause

Three categories of API mismatch:

### 1. Embedded Struct Field Access in Struct Literals

`LeftMiddleRight` embeds `coregeneric.Triple[string,string,string]` and `LeftRight` embeds `coregeneric.Pair[string,string]`.
Go does **not** allow promoted field names in struct literals — you must either:
- Use the embedding field name: `corestr.LeftMiddleRight{Triple: coregeneric.Triple[...]{Left: "a", ...}}`
- Use a constructor: `corestr.NewLeftMiddleRight("a", "b", "c")`

**Wrong:** `corestr.LeftMiddleRight{Left: "a", Middle: "b", Right: "c"}`
**Right:** `corestr.NewLeftMiddleRight("a", "b", "c")`

### 2. Wrong Fields on ValueStatus

`ValueStatus` has fields `ValueValid *ValidValue` and `Index int` — **not** `Value string` and `IsValid bool`.
The test assumed flat fields that actually live inside the nested `ValidValue` struct.

**Wrong:** `corestr.ValueStatus{Value: "hello", IsValid: true}`
**Right:** `corestr.ValueStatus{ValueValid: &corestr.ValidValue{Value: "hello"}, Index: 0}`

### 3. Non-Existent Method on Hashmap

`SortedKeys()` exists on `HashmapDiff`, not `Hashmap`.
`Hashmap` has `Keys()`, `AllKeys()`, `KeysLock()`, `KeysCollection()`, etc.

**Wrong:** `h.SortedKeys()`
**Right:** `h.Keys()`

## Fix Applied

- Line 449: `LeftMiddleRight` → use `NewLeftMiddleRight()` constructor
- Line 462: `LeftRight` → use `NewLeftRight()` constructor
- Lines 472-477: `ValueStatus` → use correct field structure with `ValidValue`
- Line 558: `Hashmap.SortedKeys()` → `Hashmap.Keys()`

## Learning

1. **Embedded struct types cannot use promoted field names in struct literals** — always check if a constructor exists or use the embedding field name.
2. **Nested struct fields ≠ flat fields** — always read the actual struct definition, not just the type name.
3. **Method availability differs between related types** — `HashmapDiff.AllKeysSorted()` ≠ `Hashmap.SortedKeys()`.
4. These are all instances of the general rule: **read the source before writing tests**.
