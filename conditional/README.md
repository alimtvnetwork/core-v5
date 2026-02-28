# conditional — Generic Ternary & Nil-Safe Helpers

## Overview

The `conditional` package provides generic ternary expressions, nil-safe defaults, conditional function execution, and batch function runners for Go. It replaces verbose `if/else` blocks with concise, type-safe one-liners.

## Typed Convenience Wrappers (`typed_*.go`)

For common primitive types, typed wrappers eliminate the need to specify type parameters:

```go
result := conditional.IfInt(isTrue, 2, 7)                          // no type param needed
name   := conditional.IfFuncString(ok, trueFunc, falseFunc)        // lazy evaluation
val    := conditional.IfTrueFuncInt(ok, func() int { return 42 })  // evaluate only on true
items  := conditional.IfSliceString(ok, listA, listB)              // slice ternary
ptr    := conditional.IfPtrInt(ok, &a, &b)                         // pointer ternary
defVal := conditional.NilDefFloat64(ptr, 3.14)                     // nil-safe default
defPtr := conditional.NilDefPtrString(ptr, "fallback")             // nil-safe pointer default
```

### Available Typed Wrappers

Each type has the following functions (using `Int` as example):

| Function | Description |
|----------|-------------|
| `IfInt(cond, t, f)` | Ternary for `int` values |
| `IfFuncInt(cond, tF, fF)` | Lazy-evaluated ternary |
| `IfTrueFuncInt(cond, tF)` | Evaluate only when true |
| `IfSliceInt(cond, t, f)` | Slice ternary |
| `IfSlicePtrInt(cond, t, f)` | Pointer-to-slice ternary |
| `IfSlicePtrFuncInt(cond, tF, fF)` | Lazy pointer-to-slice ternary |
| `IfPtrInt(cond, t, f)` | Pointer ternary |
| `NilDefInt*` | Nil-safe default (where available) |
| `NilDefPtrInt(ptr, def)` | Nil-safe pointer default |

### Supported Types

| File | Types | NilDef Available |
|------|-------|------------------|
| `typed_bool.go` | `bool` | `NilDefPtrBool` only (NilDef conflicts with deprecated) |
| `typed_int.go` | `int` | `NilDefPtrInt` only (NilDef conflicts with deprecated) |
| `typed_int8.go` | `int8` | Both `NilDefInt8` and `NilDefPtrInt8` |
| `typed_int16.go` | `int16` | Both `NilDefInt16` and `NilDefPtrInt16` |
| `typed_int32.go` | `int32` | Both `NilDefInt32` and `NilDefPtrInt32` |
| `typed_int64.go` | `int64` | Both `NilDefInt64` and `NilDefPtrInt64` |
| `typed_float32.go` | `float32` | Both `NilDefFloat32` and `NilDefPtrFloat32` |
| `typed_float64.go` | `float64` | Both `NilDefFloat64` and `NilDefPtrFloat64` |
| `typed_string.go` | `string` | Both `NilDefString` and `NilDefPtrString` |
| `typed_byte.go` | `byte` | `NilDefPtrByte` only (NilDef conflicts with deprecated) |

> **Note**: For `bool`, `int`, and `byte`, `NilDef<Type>` is omitted because
> deprecated functions with the same name but different signatures already exist.
> Use `NilDef[bool](ptr, defVal)` directly for those types.

## Core Generic Functions (`generic.go`)

### Ternary Helpers

```go
result := conditional.If[int](isTrue, 2, 7)                    // generic ternary
name   := conditional.IfFunc[string](ok, trueFunc, falseFunc)   // lazy evaluation
val    := conditional.IfTrueFunc[int](ok, func() int { ... })   // evaluate only on true
items  := conditional.IfSlice[string](ok, listA, listB)         // slice ternary
ptr    := conditional.IfPtr[int](ok, &a, &b)                    // pointer ternary
```

### Nil-Safe Defaults

```go
val := conditional.NilDef[int](ptr, 42)         // dereference or default
p   := conditional.NilDefPtr[string](ptr, "x")  // return pointer or pointer-to-default
res := conditional.NilCheck(maybeNil, onNil, onNonNil)  // any-typed nil branch
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

## Legacy Per-Type Functions (Deprecated)

Retained for backward compatibility — use generic equivalents instead:

| Deprecated | Replacement |
|-----------|-------------|
| `Bool(cond, t, f)` | `If[bool](cond, t, f)` |
| `Int(cond, t, f)` | `If[int](cond, t, f)` |
| `String(cond, t, f)` | `If[string](cond, t, f)` |
| `Byte(cond, t, f)` | `If[byte](cond, t, f)` |
| `Interface(cond, t, f)` | `If[any](cond, t, f)` |
| `Integers(cond, t, f)` | `IfSlice[int](cond, t, f)` |
| `Strings(cond, t, f)` | `IfSlice[string](cond, t, f)` |
| `BoolFunc(cond, tF, fF)` | `IfFunc[bool](cond, tF, fF)` |
| `StringFunc(cond, tF, fF)` | `IfFunc[string](cond, tF, fF)` |
| `StringTrueFunc(cond, tF)` | `IfTrueFunc[string](cond, tF)` |
| `BooleanTrueFunc(cond, tF)` | `IfTrueFunc[bool](cond, tF)` |
| `BytesTrueFunc(cond, tF)` | `IfTrueFunc[[]byte](cond, tF)` |
| `NilDefStr(ptr, def)` | `NilDef[string](ptr, def)` |
| `NilDefInt(ptr, def)` | `NilDef[int](ptr, def)` |
| `NilDefBool(ptr, def)` | `NilDef[bool](ptr, def)` |
| `NilDefByte(ptr, def)` | `NilDef[byte](ptr, def)` |
| `InterfaceFunc(cond, tF, fF)` | `IfFunc[any](cond, tF, fF)` |

### Deprecated Pointer/Slice Variants

| Deprecated | Replacement |
|-----------|-------------|
| `StringPtr(cond, t, f)` | `IfPtr[string](cond, t, f)` |
| `IntegersPtr(cond, t, f)` | `IfSlice[int](cond, t, f)` |
| `StringsPtr(cond, t, f)` | `IfSlice[string](cond, t, f)` |
| `BytesPtr(cond, t, f)` | `IfSlice[byte](cond, t, f)` |
| `BooleansPtr(cond, t, f)` | `IfSlice[bool](cond, t, f)` |
| `InterfacesPtr(cond, t, f)` | `IfSlice[any](cond, t, f)` |

## File Organization

| File | Responsibility |
|------|---------------|
| `generic.go` | All generic functions (`If`, `IfFunc`, `NilDef`, etc.) |
| `typed_bool.go` ... `typed_byte.go` | Typed convenience wrappers for 10 primitive types |
| `funcs.go` | Internal helper functions |
| `Bool.go`, `String.go`, `Int.go`, `Byte.go` | Deprecated per-type ternaries |
| `Booleans.go`, `Strings.go`, `Integers.go`, `Bytes.go` | Deprecated slice ternaries |
| `*Ptr.go` | Deprecated pointer variants |
| `*TrueFunc.go` | Deprecated true-only function variants |
| `*Func.go` | Deprecated function-based ternaries |
| `NilDef*.go`, `NilCheck.go`, `DefOnNil.go` | Nil-safe default helpers |
| `VoidFunctions.go` | Void batch execution |
| `Functions.go`, `FunctionsExecuteResults.go` | Result batch execution |
| `ErrorFunc.go`, `ErrorFunctionsExecuteResults.go` | Error batch execution |
| `TypedErrorFunctionsExecuteResults.go` | Typed error batch execution |
| `AnyFunctions.go`, `AnyFunctionsExecuteResults.go` | Any-typed batch execution |
| `Setter.go`, `SetterDefault.go` | Conditional setters |
| `BoolByOrder.go`, `BoolFunctionsByOrder.go` | Order-based boolean helpers |
| `StringsIndexVal.go`, `StringDefault.go` | String utility helpers |
| `ErrorFunctionResult.go` | Error function result type |
| `executeErrorFunctions.go`, `executeVoidFunctions.go` | Internal execution logic |

## Key Patterns

- **`isTake` / `isBreak`**: Control flags for batch execution — `isTake` determines whether a result is collected, `isBreak` halts execution.
- **Error aggregation**: All errors from batch execution are merged via `errcore.SliceToError` with index metadata appended for debugging.
- **Generic-first**: New code should use the generic functions (`If[T]`, `NilDef[T]`). Per-type wrappers exist only for backward compatibility.

## How to Extend Safely

- **New generic helper**: Add to `generic.go`.
- **New batch execution variant**: Create a dedicated file following the `*FunctionsExecuteResults.go` naming convention.
- **New type-specific function**: **Don't** — use the generic equivalent instead.

## Related Docs

- [Repo Overview](../spec/01-app/00-repo-overview.md)
