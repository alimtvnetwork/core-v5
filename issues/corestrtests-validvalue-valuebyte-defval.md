# Issue: ValidValue.ValueByte ignores defVal parameter

## Date: 2026-03-25

## Status: ✅ FIXED

## Failing Test

- `Test_C40_ValidValue_NumericConversions` (line 63)

## Root Cause

`ValidValue.ValueByte(defVal byte)` in `coredata/corestr/ValidValue.go` returned
`constants.Zero` on parse error instead of the caller-supplied `defVal`:

```go
if err != nil || toInt < 0 {
    return constants.Zero  // BUG: ignores defVal
}
```

The test called `bad.ValueByte(88)` where `bad` has value `"abc"` (unparseable).
Expected return: `88` (the default). Actual return: `0`.

## Solution

Return `defVal` instead of `constants.Zero` on error:

```go
if err != nil || toInt < 0 {
    return defVal
}
```

## Affected Files

- `coredata/corestr/ValidValue.go` (line ~151) — production fix
