# CaseNilSafe Design Document

## Purpose

`CaseNilSafe` provides a **compile-time-safe, table-driven** pattern for systematically testing nil-receiver safety across pointer receiver methods. It replaces scattered, inconsistent nil tests with a single declarative structure.

## Problem Statement

The codebase has **~51 files** containing nil-receiver tests in at least **4 different styles**:

| Style | Example | Count | Issues |
|-------|---------|-------|--------|
| Inline `if` + `t.Error` | `corestrtests/Hashset_test.go` | ~15 files | No AAA, no diagnostics, raw `t.Error` |
| CaseV1 table-driven | `coreinstructiontests/StringCompare` | ~12 files | Verbose, per-method boilerplate |
| Custom wrapper struct | `coreoncetests/BytesErrorOnce` | ~5 files | Non-standard, hard to discover |
| GenericGherkins | `corepayloadtests/TypedCollection` | ~8 files | Overkill for simple nil checks |

**CaseNilSafe** unifies all four into one consistent pattern.

## Architecture

### Package Layout

```
coretests/
├── results/                      ← NEW package
│   ├── Result.go                 ← Result[T] base struct
│   ├── ResultHelpers.go          ← ToMap, ReturnCount helpers
│   ├── Results.go                ← Results[T1, T2] for multi-return
│   ├── Invoke.go                 ← InvokeWithPanicRecovery engine
│   ├── MethodName.go             ← Reflection-based name extraction
│   └── aliases.go                ← ResultBool, ResultString, etc.
├── coretestcases/
│   ├── CaseNilSafe.go            ← Test case struct + Invoke
│   ├── CaseNilSafeAssertions.go  ← ShouldBeEqualMap, ShouldBeSafe
│   └── CaseNilSafeAssertHelper.go← Delegation to errcore
```

### Core Types

```go
// Result[T] — single-return invocation result
type Result[T any] struct {
    Value        T       // primary return value
    Error        error   // error from function or nil
    Panicked     bool    // true if recovered from panic
    PanicValue   any     // raw recover() value
    AllResults   []any   // ALL return values (for multi-return)
    ReturnCount  int     // number of return values
}

// Results[T1, T2] — two-return invocation result
type Results[T1, T2 any] struct {
    Result[T1]
    Result2 T2
}

// CaseNilSafe — the test case
type CaseNilSafe struct {
    Title    string     // scenario name
    Func     any        // direct method ref: (*Type).Method
    Args     []any      // optional input arguments
    Expected args.Map   // expected outcome map
}
```

## Edge Cases

### 1. Pointer vs Value Receivers

```go
type MyStruct struct { Name string }

// Pointer receiver — nil-safe if guarded
func (it *MyStruct) IsValid() bool {
    if it == nil { return false }
    return it.Name != ""
}

// Value receiver — ALWAYS panics on nil (Go dereferences)
func (it MyStruct) String() string {
    return it.Name
}
```

**Behavior:**
- `(*MyStruct).IsValid` with nil → `Result{Value: false, Panicked: false}` ✓
- `(*MyStruct).String` with nil → `Result{Panicked: true}` (Go auto-dereferences)

**CaseNilSafe handles both** — the test case simply declares `"panicked": true/false` in Expected.

### 2. Multi-Return Methods

```go
func (it *MyStruct) Parse(s string) (int, error) { ... }
```

**Solution:** `extractResult` populates `AllResults []any` with all return values. `Value` gets the first, `Error` gets the last if it implements `error`.

```go
// Test case
CaseNilSafe{
    Func: (*MyStruct).Parse,
    Args: []any{"123"},
    Expected: args.Map{
        "value":       0,
        "hasError":    true,
        "panicked":    false,
        "returnCount": 2,
    },
}
```

### 3. Void Methods (No Returns)

```go
func (it *MyStruct) Reset() { it.data = nil }
```

**Solution:** `ReturnCount == 0`, `Value` stays zero-value.

```go
CaseNilSafe{
    Func: (*MyStruct).Reset,
    Expected: args.Map{
        "panicked":    false,
        "returnCount": 0,
    },
}
```

### 4. Methods Returning Non-Error Second Value

```go
func (it *MyStruct) Lookup(key string) (string, bool) { ... }
```

**Solution:** `AllResults[1]` is `bool`, not `error`. `Error` stays nil. Callers use `AllResults` for full access.

```go
CaseNilSafe{
    Func: (*MyStruct).Lookup,
    Args: []any{"key"},
    Expected: args.Map{
        "value":       "",
        "panicked":    false,
        "returnCount": 2,
    },
}
```

### 5. Methods with Complex Arguments

```go
func (it *MyStruct) Process(ctx context.Context, opts Options) error { ... }
```

**Solution:** `Args` field accepts `[]any`:

```go
CaseNilSafe{
    Func: (*MyStruct).Process,
    Args: []any{context.Background(), Options{Verbose: true}},
    Expected: args.Map{
        "panicked": false,
        "hasError": true,
    },
}
```

### 6. Interface Return Types

```go
func (it *MyStruct) Clone() Cloneable { ... }
```

**Pitfall:** `reflect.Value.Interface()` on a nil interface returns `nil` typed as `any`, not as `Cloneable`. `extractResult` handles this by checking `IsNil()` before extracting.

### 7. Generic Type Methods (Function Literal Wrapper)

Go does not support method expressions on generic types:

```go
// ❌ COMPILE ERROR — Go cannot form method expressions on generic types
(*coregeneric.LinkedList[int]).IsEmpty
```

**Recommended pattern:** Use a **function literal wrapper** that calls the method directly. This preserves compile-time safety — renaming the method still causes a build error at the call site.

```go
CaseNilSafe{
    Title: "Length on nil returns 0",
    Func: func(c *corepayload.TypedPayloadCollection[testUser]) int {
        return c.Length()
    },
    Expected: args.Map{
        "value":    "0",
        "panicked": false,
    },
}
```

**Why this works:**
- The literal's signature tells `buildCallArgs` the first parameter type, so it creates a typed nil pointer of the correct generic instantiation.
- `InvokeWithPanicRecovery` calls the literal via reflection identically to a method expression.
- Renaming `Length()` → build error inside the literal body.

**When to use:** Any method on a generic struct (`LinkedList[T]`, `TypedPayloadCollection[T]`, `TypedSimpleGenericRequest[T]`, etc.). Non-generic types should still prefer direct method expressions (`(*MyStruct).IsValid`).

### 8. Nil Func Reference

If `Func` is nil (programmer error), `InvokeWithPanicRecovery` returns `Result{Panicked: true, PanicValue: "funcRef is nil"}`. This fails the test loudly rather than silently passing.

## Helper Methods on Result

| Method | Returns | Purpose |
|--------|---------|---------|
| `IsSafe()` | `bool` | `!Panicked && Error == nil` |
| `HasError()` | `bool` | `Error != nil` |
| `HasPanicked()` | `bool` | `Panicked` |
| `IsResult(expected)` | `bool` | `%v` equality check |
| `IsResultTypeOf(expected)` | `bool` | `reflect.Type` assignability |
| `IsError(msg)` | `bool` | Error message match |
| `ValueString()` | `string` | `fmt.Sprintf("%v", Value)` |
| `ToMap()` | `args.Map` | Standard map for assertion |

### ToMap Output

```go
result.ToMap() → args.Map{
    "value":       result.Value,
    "panicked":    result.Panicked,
    "isSafe":      result.IsSafe(),
    "hasError":    result.HasError(),
    "returnCount": result.ReturnCount,
}
```

This integrates directly with `ShouldBeEqualMap`.

## Usage Pattern

### Test Case File (`_testcases.go`)

```go
var myStructNilSafeTestCases = []coretestcases.CaseNilSafe{
    {
        Title: "IsValid on nil returns false",
        Func:  (*MyStruct).IsValid,
        Expected: args.Map{
            "value":    false,
            "panicked": false,
            "isSafe":   true,
        },
    },
    {
        Title: "String on nil panics (value receiver)",
        Func:  (*MyStruct).String,
        Expected: args.Map{
            "panicked": true,
        },
    },
    {
        Title: "Parse on nil with args",
        Func:  (*MyStruct).Parse,
        Args:  []any{"hello"},
        Expected: args.Map{
            "value":       0,
            "panicked":    false,
            "hasError":    true,
            "returnCount": 2,
        },
    },
}
```

### Test Logic File (`_test.go`)

```go
func Test_MyStruct_NilSafety(t *testing.T) {
    for caseIndex, tc := range myStructNilSafeTestCases {
        // Arrange (implicit — nil receiver)

        // Act
        result := tc.InvokeNil()

        // Assert
        tc.ShouldBeEqualMap(t, caseIndex, result.ToMap())
    }
}
```

**The entire test loop is 3 lines of logic** — no per-method branching.

## Migration Plan

### Priority 1 — Inline `if` + `t.Error` (highest value)

These tests (~15 files) violate the "no raw `t.Error`" standard:
- `corestrtests/Hashset_test.go` (6 nil tests)
- `regexnewtests/LazyRegex_EdgeCases_test.go` (10 nil tests)
- `reflectmodeltests/MethodProcessor_test.go`
- `corevalidatortests/SliceValidators_test.go`
- `corevalidatortests/TextValidator_test.go`

### Priority 2 — CaseV1 nil sections

These are well-structured but verbose:
- `coreinstructiontests/StringCompare` (5 nil cases)
- `coreinstructiontests/Identifier`
- `coregenerictests/LinkedList`
- `namevaluetests/Collection`

### Priority 3 — Custom wrappers

These use bespoke nil-handling logic:
- `coreoncetests/BytesErrorOnce` (custom `IsNilReceiver` field)
- `corepayloadtests/TypedCollection` (nil receiver section)
- `coreapitests/TypedConversions`

## Compile-Time Safety Guarantee

The core value proposition: **renaming a method breaks the build**.

```go
// Before rename: compiles
CaseNilSafe{Func: (*MyStruct).IsValid}

// After rename to IsOk: build error
// (*MyStruct).IsValid undefined (type *MyStruct has no field or method IsValid)
```

String-based approaches (`"IsValid"`) would silently pass or require runtime discovery.

## Related Docs

- [Testing Patterns](../01-app/13-testing-patterns.md)
- [coretests folder spec](../01-app/folders/07-coretests.md)
- [Coding Guidelines](../01-app/17-coding-guidelines.md)
