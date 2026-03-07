# conditional — Generic Ternary & Nil-Safe Helpers

## Overview

The `conditional` package provides generic ternary expressions, nil-safe defaults, conditional function execution, and batch function runners for Go. It replaces verbose `if/else` blocks with concise, type-safe one-liners.

## Architecture

```
conditional/
├── generic.go                          # Generic base: If[T], IfFunc[T], NilDef[T], etc.
├── typed_bool.go                       # bool typed wrappers (11 functions)
├── typed_byte.go                       # byte typed wrappers (11 functions)
├── typed_string.go                     # string typed wrappers (11 functions)
├── typed_int.go                        # int typed wrappers (11 functions)
├── typed_int8.go                       # int8 typed wrappers (11 functions)
├── typed_int16.go                      # int16 typed wrappers (11 functions)
├── typed_int32.go                      # int32 typed wrappers (11 functions)
├── typed_int64.go                      # int64 typed wrappers (11 functions)
├── typed_uint.go                       # uint typed wrappers (11 functions)
├── typed_uint8.go                      # uint8 typed wrappers (11 functions)
├── typed_uint16.go                     # uint16 typed wrappers (11 functions)
├── typed_uint32.go                     # uint32 typed wrappers (11 functions)
├── typed_uint64.go                     # uint64 typed wrappers (11 functions)
├── typed_float32.go                    # float32 typed wrappers (11 functions)
├── typed_float64.go                    # float64 typed wrappers (11 functions)
├── typed_wrappers.go                   # Composite/any typed wrappers
├── NilOrEmpty.go                       # String-specific nil/empty checks
├── NilCheck.go                         # Any-typed nil check (legacy, unique)
├── DefOnNil.go                         # Any-typed default-on-nil (unique)
├── funcs.go                            # Internal helper functions
├── VoidFunctions.go                    # Void batch execution
├── Functions.go                        # Result batch execution
├── FunctionsExecuteResults.go          # Result batch with isTake/isBreak
├── ErrorFunc.go                        # Error batch execution
├── ErrorFunctionsExecuteResults.go     # Error batch with aggregation
├── TypedErrorFunctionsExecuteResults.go # Typed (T, error) batch execution
├── AnyFunctions.go                     # any-typed batch execution
├── AnyFunctionsExecuteResults.go       # any-typed batch with control
├── Setter.go                           # Conditional setters
├── SetterDefault.go                    # Conditional setters with default
├── BoolByOrder.go                      # Order-based boolean helpers
├── BoolFunctionsByOrder.go             # Order-based boolean function helpers
├── StringsIndexVal.go                  # Index-based string selection
├── StringDefault.go                    # String with empty default
├── ErrorFunctionResult.go              # Error function result type
├── Func.go                             # Function selection (returns func)
├── executeErrorFunctions.go            # Internal error execution
└── executeVoidFunctions.go             # Internal void execution
```

## Generic Base Functions (`generic.go`)

All typed wrappers delegate to these generic base functions. Use these directly when
working with custom types, or use the [typed wrappers](#typed-convenience-wrappers) for primitives.

### Signature Table

| Function | Signature | Description |
|----------|-----------|-------------|
| `If[T]` | `(cond bool, t, f T) T` | Ternary — returns `t` if true, `f` if false |
| `IfFunc[T]` | `(cond bool, tF, fF func() T) T` | Lazy ternary — evaluates only the chosen branch |
| `IfTrueFunc[T]` | `(cond bool, tF func() T) T` | True-only — evaluates `tF` on true, zero value on false |
| `IfSlice[T]` | `(cond bool, t, f []T) []T` | Slice ternary |
| `IfPtr[T]` | `(cond bool, t, f *T) *T` | Pointer ternary |
| `NilDef[T]` | `(ptr *T, def T) T` | Dereference or default |
| `NilDefPtr[T]` | `(ptr *T, def T) *T` | Return pointer or pointer-to-default |
| `NilVal[T]` | `(ptr *T, onNil, onNonNil T) T` | Choose between two values based on nil check (no deref) |
| `NilValPtr[T]` | `(ptr *T, onNil, onNonNil T) *T` | Like `NilVal` but returns pointer to chosen value |
| `ValueOrZero[T]` | `(ptr *T) T` | Dereference or zero value |
| `PtrOrZero[T]` | `(ptr *T) *T` | Return pointer or pointer-to-zero (guaranteed non-nil) |

### Usage Examples

```go
// Ternary
result := conditional.If[int](isTrue, 2, 7)                    // returns 2 or 7
name   := conditional.If[string](len(s) > 0, s, "default")

// Lazy evaluation — only the chosen branch is called
val := conditional.IfFunc[string](ok,
    func() string { return expensiveCall() },
    func() string { return "fallback" },
)

// True-only — zero value on false
val := conditional.IfTrueFunc[int](ok, func() int { return compute() })

// Nil-safe defaults
val := conditional.NilDef[int](ptr, 42)         // dereference or 42
p   := conditional.NilDefPtr[string](ptr, "x")  // pointer or &"x"

// Nil branching without dereference
label := conditional.NilVal[string](namePtr, "(unknown)", "has name")

// Zero-value deref
active := conditional.ValueOrZero[bool](flagPtr)   // false if nil
safePtr := conditional.PtrOrZero[int](intPtr)       // guaranteed non-nil
```

## Typed Convenience Wrappers

For all 15 primitive types, typed wrappers eliminate the need to specify type parameters.
Each type gets the same 11 functions, named with the type suffix (e.g., `IfInt`, `IfFuncString`).

```go
result := conditional.IfInt(isTrue, 2, 7)                          // no type param needed
name   := conditional.IfFuncString(ok, trueFunc, falseFunc)        // lazy evaluation
val    := conditional.IfTrueFuncInt(ok, func() int { return 42 })  // evaluate only on true
items  := conditional.IfSliceString(ok, listA, listB)              // slice ternary
ptr    := conditional.IfPtrInt(ok, &a, &b)                         // pointer ternary
defVal := conditional.NilDefFloat64(ptr, 3.14)                     // nil-safe default
defPtr := conditional.NilDefPtrString(ptr, "fallback")             // nil-safe pointer default
label  := conditional.NilValInt(ptr, -1, 100)                      // nil branch without deref
zero   := conditional.ValueOrZeroBool(flagPtr)                     // false if nil
safe   := conditional.PtrOrZeroInt64(numPtr)                       // guaranteed non-nil
```

### Function Set Per Type (11 per type)

| Function | Signature | Description |
|----------|-----------|-------------|
| `If{T}` | `(cond, t, f) T` | Ternary |
| `IfFunc{T}` | `(cond, tF, fF) T` | Lazy ternary |
| `IfTrueFunc{T}` | `(cond, tF) T` | True-only lazy |
| `IfSlice{T}` | `(cond, t, f) []T` | Slice ternary |
| `IfPtr{T}` | `(cond, t, f) *T` | Pointer ternary |
| `NilDef{T}` | `(ptr, def) T` | Deref or default |
| `NilDefPtr{T}` | `(ptr, def) *T` | Pointer or &default |
| `NilVal{T}` | `(ptr, onNil, onNonNil) T` | Nil branch (no deref) |
| `NilValPtr{T}` | `(ptr, onNil, onNonNil) *T` | Nil branch → pointer |
| `ValueOrZero{T}` | `(ptr) T` | Deref or zero |
| `PtrOrZero{T}` | `(ptr) *T` | Pointer or &zero |

### All 15 Supported Types

| File | Type | Full 11 Functions |
|------|------|:-----------------:|
| `typed_bool.go` | `bool` | ✓ |
| `typed_byte.go` | `byte` | ✓ |
| `typed_string.go` | `string` | ✓ |
| `typed_int.go` | `int` | ✓ |
| `typed_int8.go` | `int8` | ✓ |
| `typed_int16.go` | `int16` | ✓ |
| `typed_int32.go` | `int32` | ✓ |
| `typed_int64.go` | `int64` | ✓ |
| `typed_uint.go` | `uint` | ✓ |
| `typed_uint8.go` | `uint8` | ✓ |
| `typed_uint16.go` | `uint16` | ✓ |
| `typed_uint32.go` | `uint32` | ✓ |
| `typed_uint64.go` | `uint64` | ✓ |
| `typed_float32.go` | `float32` | ✓ |
| `typed_float64.go` | `float64` | ✓ |

> For the complete per-type function reference, see [typed-wrappers.md](./typed-wrappers.md).

## String-Specific Helpers (`NilOrEmpty.go`)

```go
// Check for both nil and empty string
val := conditional.NilOrEmptyStr(strPtr, "fallback", "has value")
ptr := conditional.NilOrEmptyStrPtr(strPtr, "fallback", "has value")
```

## Batch Function Execution

### Void Functions (`VoidFunctions.go`)

Execute a sequence of void functions. Uses `isTake` / `isBreak` semantics to control collection and short-circuiting.

```go
conditional.VoidFunctions(fn1, fn2, fn3)
```

### Result Functions (`Functions.go`, `FunctionsExecuteResults.go`)

Execute functions and collect results:

```go
results := conditional.Functions(fn1, fn2, fn3)             // collect []T results
results := conditional.FunctionsExecuteResults(fn1, fn2)    // with isTake/isBreak control
```

### Error Functions (`ErrorFunc.go`, `ErrorFunctionsExecuteResults.go`)

Execute error-returning functions with aggregation:

```go
err := conditional.ErrorFunc(fn1, fn2, fn3)                           // aggregate errors
results, err := conditional.ErrorFunctionsExecuteResults(fn1, fn2)    // results + error
```

Errors are aggregated via `errcore.SliceToError` with index metadata for debugging.

### Typed Error Functions (`TypedErrorFunctionsExecuteResults.go`)

Execute functions returning `(T, error)` with aggregation:

```go
results, err := conditional.TypedErrorFunctionsExecuteResults(fn1, fn2)
```

### Any Functions (`AnyFunctions.go`, `AnyFunctionsExecuteResults.go`)

Execute functions returning `any`:

```go
results := conditional.AnyFunctions(fn1, fn2)
```

## Conditional Setters (`Setter.go`, `SetterDefault.go`)

```go
conditional.Setter(isApply, &target, value)              // set if condition true
conditional.SetterDefault(isApply, &target, value, def)  // set value or default
```

## Decision Guide

| Scenario | Use |
|----------|-----|
| Known primitive type | Typed wrapper: `IfInt(...)`, `NilDefString(...)` |
| Custom struct/interface | Generic: `If[MyType](...)`, `NilDef[MyType](...)` |
| Slice of primitives | Typed: `IfSliceInt(...)` |
| Slice of custom type | Generic: `IfSlice[MyType](...)` |
| Map type | Generic: `If[map[K]V](...)` |
| Func type | Generic: `If[func() error](...)` |

## Key Patterns

- **`isTake` / `isBreak`**: Control flags for batch execution — `isTake` determines whether a result is collected, `isBreak` halts execution.
- **Error aggregation**: All errors from batch execution are merged via `errcore.SliceToError` with index metadata appended for debugging.
- **Generic-first**: New code should use generic functions (`If[T]`, `NilDef[T]`) or typed wrappers.

## How to Extend Safely

- **New generic helper**: Add to `generic.go`.
- **New batch execution variant**: Create a dedicated file following the `*FunctionsExecuteResults.go` naming convention.
- **New type-specific function**: **Don't** — use the generic equivalent instead.

## Related Docs

- [Typed Wrappers Reference](./typed-wrappers.md)
- [Repo Overview](../spec/01-app/00-repo-overview.md)
