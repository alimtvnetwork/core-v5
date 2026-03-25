# Issue: args.Map ExpectedInput Falls Through to PrettyJSON in Assertion Output

## Summary

When `ShouldBeEqualMap` is used, the **expected** side of the assertion renders as a
struct-wrapped `PrettyJSON` format instead of flat `"key : value"` lines, causing a
length mismatch with the **actual** side.

## Symptom

```
Actual Received (4 lines):
  "isInvalid : false",
  "isValid : true",
  "isZero : false",
  "value : 5",

Expected Input (7 lines):
  "Map {",
  "    isInvalid : false,",
  "    isValid : true,",
  "    isZero : false,",
  "    value : 5,",
  "}",
  "",
```

Length mismatch: Expect ["4"] != ["7"] Actual — assertion always fails.

## Root Cause

In `CaseV1MapAssertions.go`, `ShouldBeEqualMap` does:

```go
actualLines := actual.CompileToStrings()    // → flat "key : value" lines ✓
it.ExpectedInput = it.ExpectedAsMap()        // → still args.Map (named type)
it.ShouldBe(...)                             // → calls ExpectedLines()
```

`ExpectedLines()` calls `convertinternal.AnyTo.Strings(it.ExpectedInput)`.

In the `AnyTo.Strings()` type switch, `args.Map` is a **named type** (`type Map map[string]any`).
Go's type switch does NOT match named types against their underlying type, so `args.Map`
does **not** match the `case map[string]any:` branch. It falls through to `default` →
`PrettyJsonLines()`, producing the wrapped struct format.

## Fix

Convert `ExpectedInput` to compiled strings (same as actual) before passing to `ShouldBe`:

```go
it.ExpectedInput = it.ExpectedAsMap().CompileToStrings()
```

This ensures both sides use identical `"key : value"` flat-line format.

## Affected Files

- `coretests/coretestcases/CaseV1MapAssertions.go` — primary fix location
- `internal/convertinternal/anyTo.go` — secondary: could also add `args.Map` case (but creates import cycle)

## Impact

All tests using `ShouldBeEqualMap` / `ShouldBeEqualMapFirst` (137+ files, 5400+ call sites).
