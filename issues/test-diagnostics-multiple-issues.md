# Issue: Test Diagnostics — Multiple Failures and Output Quality (FIXED)

## Issue 1: Stack Trace Printed Twice

### Symptom
```
Stack-Trace: 
  - Stack-Trace :
  - D:/wp-work/.../defaulterr.go:42
  - C:/Program Files/Go/src/runtime/proc.go:7670
```

### Root Cause
In `errcore/stackTraceEnhance.go`, `trace()` (line 88) calls:
```go
reflectinternal.CodeStack.StacksStringCount(2+skip, 4)
```

`StacksStringCount` (in `internal/reflectinternal/codeStack.go:192`) **already prefixes** its output with `"Stack-Trace :\n  - ..."`.

Then `stackEnhanceFormat` (in `errcore/consts.go:40`) is:
```go
"%s - \n%s\n %s: \n  - %s"
//                ^constants.StackTrace   ^toTrace (already has "Stack-Trace :" header)
```

So the final output becomes:
```
methodName - 
msg
 Stack-Trace: 
  - Stack-Trace :
  - line1
  - line2
```

**Double header.** The format includes `constants.StackTrace` as a label, AND the trace string itself starts with `"Stack-Trace :"`.

### Fix Direction
Either:
- A) Change `trace()` to call `StacksStringsCount` (returns raw `[]string`) and join manually — so `stackEnhanceFormat` provides the only header.
- B) Change `StacksStringCount` to NOT include the header, making it a pure join function.

Option A is safer (fewer callers affected).

---

## Issue 2: System Library Paths in Stack Traces

### Symptom
```
Stack-Trace :
  - D:/wp-work/.../defaulterr.go:42
  - C:/Program Files/Go/src/runtime/proc.go:7670
  - C:/Program Files/Go/src/runtime/proc.go:7637
  - C:/Program Files/Go/src/runtime/proc.go:256
```

### Root Cause
`codeStack.NewFileWithLines()` uses `runtime.Caller(i)` to collect frames but does **no filtering**. Go runtime frames (`runtime/`, `testing/`) are included as-is.

### Fix Direction
In `NewFileWithLines` (or in `StacksStringsCount`), skip frames whose `file` path contains known Go standard library prefixes:
- `runtime/`
- `testing/`
- Or more broadly: anything under `GOROOT`

Use `runtime.GOROOT()` to detect and filter.

---

## Issue 3: Single-Value Comparison Shows Line-Diff Format

### Symptom
```
ExpectedInput: ""   (single empty string)
Actual result: ""   (single empty string)

Output shows:
  ActualLines, ExpectedLines Length is not equal. - Expect ["1"] != ["0"] Actual
  ============================>
  0 ) Actual Received:
      Nil info SafeName returns empty
  ============================>
  "",
  ============================>
```

A simple `string == string` comparison is rendered as a multi-line diff block, which is confusing and unnecessary.

### Root Cause
`ShouldBeEqual` always converts to `[]string` via `ExpectedLines()` and runs through the `SliceValidator` pipeline. For `ExpectedInput: ""`, `convertinternal.AnyTo.Strings("")` returns `[]string{""}` (1 element). But there may be an empty-vs-nil distinction causing length mismatch.

More importantly, even when lengths match, the output format is overly verbose for single-value comparisons. A single string comparison should show:
```
Expected: ""
Actual:   ""
```
Not a full line-by-line diff block.

### Fix Direction
In the assertion output formatting (likely `GherkinsString` or the validator's error message builder), detect when both actual and expected are single-element slices and use a compact format instead of the full diff block.

---

## Issue 4: Gibberish Characters in Test Titles (ΓÇö)

### Symptom
```
BytesOnce 'hello' ΓÇö Value, String, IsEmpty false, Length 5
```

`ΓÇö` should be `—` (em dash, U+2014).

### Root Cause
The test title strings contain UTF-8 em dashes (`—`), but the terminal/console output is interpreting them using a non-UTF-8 code page (likely Windows CP1252 or similar). This is a **terminal encoding issue**, not a code bug.

The em dash `—` (bytes: `0xE2 0x80 0x94` in UTF-8) is displayed as `ΓÇö` when decoded as CP1252.

### Fix Direction
Two options:
- A) Replace em dashes in test titles with ASCII alternatives (`-`, `--`, or `|`)
- B) Set `chcp 65001` in the terminal (Windows UTF-8 code page)

Option A is more portable and doesn't depend on user terminal configuration.

---

## Issue 5: BytesOnce Panic on Nil InitializerFunc

### Symptom
```
panic: runtime error: invalid memory address or nil pointer dereference
  BytesOnce.Value() — BytesOnce.go:40
```

### Root Cause
In `BytesOnce_test.go`, when `UseNilInit` is true, the test creates `&coreonce.BytesOnce{}` with no `initializerFunc`. Then `Value()` calls `it.initializerFunc()` at line 40, which panics because `initializerFunc` is nil.

### Fix Direction
Add a nil guard in `BytesOnce.Value()`:
```go
if it.initializerFunc == nil {
    it.isInitialized = true
    return it.innerData  // zero value
}
```

---

## Issue 6: Actual/Expected Alignment in Diff Output

### Symptom
```
  Line   2 [MISMATCH]:
           actual : `isDefined : true`
         expected : `isDefined : false`
```

The `actual` and `expected` labels have inconsistent left-padding, making visual comparison harder.

### Root Cause
In `errcore/LineDiff.go`, the format string uses:
```go
"  Line %3d [MISMATCH]:\n"+
"           actual : `%s`\n"+
"         expected : `%s`\n"
```

`actual` has 11 spaces prefix, `expected` has 9 — misaligned.

### Fix Direction
Align both labels to the same column:
```go
"  Line %3d [MISMATCH]:\n"+
"         actual   : `%s`\n"+
"         expected : `%s`\n"
```

---

## Issue 7: Expected Map Values Not Copy-Pasteable

### Symptom  
Diff shows `isDefined : true` instead of the Go literal format `"isDefined": true` that can be copy-pasted into test case definitions.

### Root Cause
`args.Map.CompileToStrings()` uses `"key : value"` format (with spaces around `:`). The expected output section should show the raw Go map literal format for easy copy-paste into `_testcases.go` files.

### Fix Direction
Add a separate "Expected Values" block in the failure output that shows the map in Go literal format:
```
Expected:
    "isDefined": false,
    "isEmpty":   true,
```
