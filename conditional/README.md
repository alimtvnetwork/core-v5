# conditional — Generic Ternary & Nil-Safe Helpers

## Overview

The `conditional` package provides generic ternary expressions, nil-safe defaults, conditional function execution, and batch function runners for Go. It replaces verbose `if/else` blocks with concise, type-safe one-liners.

## Core Generic Functions

### Ternary Helpers

```go
result := conditional.If[int](isTrue, 2, 7)              // generic ternary
name   := conditional.IfFunc[string](ok, trueFunc, falseFunc)  // lazy evaluation
val    := conditional.IfTrueFunc[int](ok, func() int { ... })  // evaluate only on true
items  := conditional.IfSlice[string](ok, listA, listB)        // slice ternary
ptr    := conditional.IfPtr[int](ok, &a, &b)                   // pointer ternary
```

### Nil-Safe Defaults

```go
val := conditional.NilDef[int](ptr, 42)       // dereference or default
p   := conditional.NilDefPtr[string](ptr, "x") // return pointer or pointer-to-default
res := conditional.NilCheck(maybeNil, onNil, onNonNil)  // any-typed nil branch
```

## Legacy Per-Type Functions (Deprecated)

The following are retained for backward compatibility but should be replaced with generics:

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
| `StringTrueFunc(cond, tF)` | `IfTrueFunc[string](cond, tF)` |
| `NilDefStr(ptr, def)` | `NilDef[string](ptr, def)` |
| `NilDefInt(ptr, def)` | `NilDef[int](ptr, def)` |
| `NilDefBool(ptr, def)` | `NilDef[bool](ptr, def)` |
| `NilDefByte(ptr, def)` | `NilDef[byte](ptr, def)` |

## Batch Function Execution

| Function | Description |
|----------|-------------|
| `VoidFunctions(fns...)` | Execute all void functions sequentially |
| `Functions(fns...)` | Execute all, collect results |
| `AnyFunctions(fns...)` | Execute all, collect `any` results |
| `ErrorFunc(fns...)` | Execute all, collect errors |

## Setter Utilities

| Function | Description |
|----------|-------------|
| `Setter(cond, target, value)` | Set target to value if condition is true |
| `SetterDefault(cond, target, value, def)` | Set value or default based on condition |

## File Organization

| File | Responsibility |
|------|---------------|
| `generic.go` | All generic functions (`If`, `IfFunc`, `NilDef`, etc.) |
| `Bool.go`, `String.go`, `Int.go`, etc. | Deprecated per-type ternaries |
| `NilDef*.go`, `NilCheck.go` | Nil-safe default helpers |
| `*Functions*.go` | Batch function execution |
| `Setter*.go` | Conditional setters |

## Contributors

## Issues for Future Reference
