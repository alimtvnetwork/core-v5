# coretests/args

The `args` package provides **typed argument holders** for structuring test case inputs, expected outputs, and dynamic function invocation in the testing framework.

## Overview

This package solves the problem of passing heterogeneous arguments to test cases in a type-safe, introspectable way. It provides:

- **Positional arg holders** (`One`, `Two`, `Three`, `Four`, `Five`, `Six`) — hold 1–6 typed arguments
- **Func arg holders** (`OneFunc`, `TwoFunc`, ..., `SixFunc`) — same as above, plus a `WorkFunc` for dynamic invocation
- **Holder** — a flexible 6-slot holder with a typed `WorkFunc` and a fallback `Hashmap`
- **FuncWrap** — a reflection-based function wrapper for dynamic invocation, validation, and introspection
- **Map / Dynamic / DynamicFunc** — key-value based argument holders for fully dynamic test cases

## Generic Architecture

All positional types are **generic** with type parameters for each argument slot:

```go
// Typed usage — compile-time safety on field access
tc := args.Three[string, int, bool]{
    First:  "hello",
    Second: 42,
    Third:  true,
}

// Untyped usage — backward-compatible, uses any for all slots
tc := args.ThreeAny{
    First:  "hello",
    Second: 42,
    Third:  true,
}
```

### Type Aliases (backward compatibility)

Every generic type has a corresponding `*Any` alias defined in `aliases.go`:

| Generic Type | Any Alias |
|---|---|
| `FuncWrap[T]` | `FuncWrapAny` |
| `One[T1]` | `OneAny` |
| `Two[T1, T2]` | `TwoAny` |
| `Three[T1, T2, T3]` | `ThreeAny` |
| `Four[T1, T2, T3, T4]` | `FourAny` |
| `Five[T1, T2, T3, T4, T5]` | `FiveAny` |
| `Six[T1, T2, T3, T4, T5, T6]` | `SixAny` |
| `OneFunc[T1]` | `OneFuncAny` |
| `TwoFunc[T1, T2]` | `TwoFuncAny` |
| `ThreeFunc[T1, T2, T3]` | `ThreeFuncAny` |
| `FourFunc[T1, T2, T3, T4]` | `FourFuncAny` |
| `FiveFunc[T1, T2, T3, T4, T5]` | `FiveFuncAny` |
| `SixFunc[T1, T2, T3, T4, T5, T6]` | `SixFuncAny` |
| `Holder[T]` | `HolderAny` |

## Positional Types (One–Six)

Hold 1–6 arguments plus an optional `Expect` field. Each positional field
is parameterized with its own type parameter:

```go
// Three holds 3 typed positional arguments.
type Three[T1, T2, T3 any] struct {
    First  T1
    Second T2
    Third  T3
    Expect any  // expected output (always any)
}
```

### Common Methods

All positional types implement `ArgBaseContractsBinder`:

- `FirstItem() any` — returns First as any (interface-compatible)
- `HasFirst() bool` — checks if First is defined (non-nil)
- `Expected() any` — returns the expected value
- `ValidArgs() []any` — collects all defined arguments
- `Args(upTo int) []any` — collects arguments up to position N
- `Slice() []any` — all fields as a cached slice (no pointer-to-slice)
- `GetByIndex(index int) any` — safe indexed access via helper
- `String() string` — formatted string representation via helper
- `ArgsCount() int` — number of positional slots (not counting Expect)

### Downcast Methods

Convert to smaller arg types while preserving type parameters:

```go
three := args.Three[string, int, bool]{First: "a", Second: 1, Third: true}
two := three.ArgTwo()  // Two[string, int]{First: "a", Second: 1}
```

## Func Types (OneFunc–SixFunc)

Same as positional types but include a `WorkFunc any` field for dynamic function invocation.
The positional arguments are typed, while WorkFunc remains `any` for reflection compatibility:

```go
tc := args.ThreeFunc[string, int, bool]{
    First:    "input1",
    Second:   42,
    Third:    true,
    WorkFunc: myFunction,  // always any
    Expect:   "expected",
}

// Invoke the function with valid args
results, err := tc.InvokeWithValidArgs()
```

### Invocation Methods

- `FuncWrap() *FuncWrapAny` — wraps WorkFunc for reflection
- `Invoke(args ...any) ([]any, error)` — invoke with explicit args
- `InvokeMust(args ...any) []any` — invoke, panic on error (via `invokeMustHelper`)
- `InvokeWithValidArgs() ([]any, error)` — invoke with all defined positional args
- `InvokeArgs(upTo int) ([]any, error)` — invoke with args up to position N

## FuncWrap[T]

A generic reflection-based function wrapper where T is the function type:

```go
// Typed construction
fw := args.NewTypedFuncWrap(func(s string) int { return len(s) })

// Untyped construction (via creator)
fw := args.NewFuncWrap.Default(myFunc)

// Introspection
fw.ArgsCount()           // number of input params
fw.ReturnLength()        // number of return values
fw.GetInArgsTypesNames() // ["string"]
fw.GetFuncName()         // "myFunc"

// Invocation
results, err := fw.Invoke("hello")
```

## Holder[T]

A flexible 6-slot holder where `T` types the `WorkFunc` field.
Positional fields (First through Sixth) remain `any` for maximum flexibility.
Includes a `Hashmap` for overflow parameters:

```go
// Typed WorkFunc
h := args.Holder[func(string) error]{
    First:    "input",
    WorkFunc: myProcessor,
}

// Untyped (backward compat)
h := args.HolderAny{
    First:    "input",
    WorkFunc: myProcessor,
    Hashmap:  args.Map{"extra": "value"},
}
```

## Map / Dynamic / DynamicFunc

Key-value based argument holders for fully dynamic test scenarios:

```go
tc := args.Dynamic{
    Params: args.Map{
        "first":  "hello",
        "second": 42,
        "func":   myFunc,
    },
    Expect: "expected",
}

results, err := tc.InvokeWithValidArgs()
```

## Shared Helpers

The package uses internal helper functions (in `argsHelper.go`) to reduce code duplication:

- `getByIndex(slice, index)` — safe indexed access
- `buildToString(typeName, slice, cache)` — cached string formatting
- `appendIfDefined(args, value)` — conditional append for defined values
- `invokeMustHelper(fw, args...)` — invoke with panic on error

## Design Decisions

### Pointer-to-Slice Removal

All types use `[]any` + `bool` flag for slice caching instead of `*[]any`.
This follows the project's pointer optimization standards for simpler API
and better Go memory efficiency.

### WorkFunc Typing

In Func variants (OneFunc–SixFunc), the `WorkFunc` field remains `any`
because it requires reflection-based invocation via `FuncWrapAny`.
Only `Holder[T]` parameterizes WorkFunc with type `T` since it's the
primary typed-function-holder pattern.

In `FuncWrap[T]`, the `Func` field is typed as `T`, enabling both
typed (`NewTypedFuncWrap`) and untyped (`NewFuncWrap.Default`) construction.

## File Organization

| File | Purpose |
|---|---|
| `One.go`–`Six.go` | Generic positional arg holders |
| `OneFunc.go`–`SixFunc.go` | Generic func arg holders |
| `Holder.go` | Generic flexible 6-slot + typed WorkFunc holder |
| `FuncWrap.go` | Core generic function wrapper struct |
| `FuncWrapArgs.go` | Argument introspection methods |
| `FuncWrapInvoke.go` | Dynamic invocation methods |
| `FuncWrapValidation.go` | Validation and error methods |
| `FuncMap.go` | Named map of function wrappers |
| `Map.go` | Key-value argument map |
| `Dynamic.go` / `DynamicFunc.go` | Map-based dynamic holders |
| `LeftRight.go` | Two-item holder with Left/Right semantics |
| `aliases.go` | All `*Any` type aliases for backward compatibility |
| `argsHelper.go` | Shared unexported utility functions |
| `all-interfaces.go` | Interface definitions |
| `consts.go` / `vars.go` | Package constants and variables |
