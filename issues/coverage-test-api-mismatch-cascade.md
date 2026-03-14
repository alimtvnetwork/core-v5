# Issue: Coverage Test API Mismatch Cascade

## Summary
12 test packages blocked from compiling due to Coverage test files written with incorrect API signatures. This is a regression — the blocked count went from 10 → 12, meaning new broken files were added to previously-working packages (coreoncetests, errcoretests, simplewraptests).

## Root Cause
Coverage test files (Coverage2-6_test.go) were generated with **assumed** API signatures instead of verified ones. The AI wrote tests referencing methods, parameters, and return types that don't exist in the source packages. This violates the 3-step methodology: the root cause was never identified (actual API signatures) before the solution was applied (writing test code).

## Categories of Mismatch

### 1. Wrong Parameter Count / Types
- `dc.AddAny("hello")` → needs `dc.AddAny("hello", true)` (missing `isValid bool`)
- `result.Clone()` → needs `result.Clone(false)` (missing `isDeepClone bool`)
- `pw.Clone()` → needs `pw.Clone(false)`, returns `(PayloadWrapper, error)` not `*PayloadWrapper`
- `ConditionalWrapWith(true, "[", "x", "]")` → `ConditionalWrapWith('[', "x", ']')` (3 params: byte, string, byte)
- `CurlyWrapOption("hello")` → `CurlyWrapOption(false, "hello")` (2 params: bool, any)
- `MustBeEmpty("test", "")` → `MustBeEmpty(err)` (1 param: error)
- `MergeErrorsToString(nil, nil)` → `MergeErrorsToString(joiner, errors...)` (string + variadic error)

### 2. Non-Existent Methods / Fields
- `fw.Name()` → `fw.Name` (field, not method)
- `fw.String()` → FuncWrapAny has no `String` method
- `bo.IsTrue()` / `bo.IsFalse()` → BoolOnce has no such methods
- `pw.PayloadIdentifier()` → no such method exists
- `hm.AllKeysSorted()` → exists on HashmapDiff, not Hashmap
- `hm.IsEquals()` → method is `IsEqual` (singular, takes value not pointer)
- `corejson.Deserialize.FromBytesTo` → correct path is `Deserialize.BytesTo`
- `errcore.Ref()` / `errcore.RefToError()` → don't exist as standalone functions
- `codestack.NewStacks` → not exported; use `codestack.New.StackTrace`

### 3. Value vs Pointer Type Mismatches
- `pc.Add(pw)` where pw is `*PayloadWrapper` → Add expects `PayloadWrapper` value
- `result != nil` for `corejson.Result` → Result is a value type, can't compare to nil
- `hm.Get("k1")` → returns `(string, bool)`, not single string
- `Collection.Take/Skip` returns `*Collection`, not `[]string` → can't use `len()`
- `Collection.Filter` expects `IsStringFilter` type (3-return func), not `func(string) bool`
- `Collection.ConcatNew` expects `(int, ...string)`, not `*Collection`

### 4. Build Artifact Conflict
- `isanytests`: stale compiled binary conflicts with fresh build — "build output already exists and is not an object file"

## Impact
- 12 packages blocked from compilation → excluded from coverage run
- Packages that HAD coverage (codestack, coredynamic, corejson, etc.) now show 0%
- Build Failure Cascade: blocked packages prevent coverage measurement for ALL cross-referenced source packages

## Solution
Fix each test file to use verified API signatures. Every API call must be checked against the source before writing test code.

## Learning / What Not To Repeat
1. **NEVER write test code with assumed signatures** — always grep/read the source first
2. **NEVER submit coverage test files in bulk** without verifying each one compiles
3. **Check return types before asserting** — value types can't be nil-checked, multi-return funcs need multi-value assignment
4. **Test files that compile but fail are better than files that don't compile** — a non-compiling test blocks the entire package
