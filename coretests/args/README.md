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

Every generic type has a corresponding `*Any` alias:

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

Hold 1–6 arguments plus an optional `Expect` field:

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
- `Slice() []any` — all fields as a cached slice
- `GetByIndex(index int) any` — safe indexed access
- `String() string` — formatted string representation
- `ArgsCount() int` — number of positional slots (not counting Expect)

### Downcast Methods

Convert to smaller arg types:

```go
three := args.ThreeAny{First: "a", Second: 1, Third: true}
two := three.ArgTwo()   // Two[any, any]{First: "a", Second: 1}
```

## Func Types (OneFunc–SixFunc)

Same as positional types but include a `WorkFunc any` field for dynamic function invocation:

```go
tc := args.ThreeFuncAny{
    First:    "input1",
    Second:   "input2",
    Third:    "input3",
    WorkFunc: myFunction,  // any callable
    Expect:   "expected",
}

// Invoke the function with valid args
results, err := tc.InvokeWithValidArgs()
```

### Invocation Methods

- `FuncWrap() *FuncWrapAny` — wraps WorkFunc for reflection
- `Invoke(args ...any) ([]any, error)` — invoke with explicit args
- `InvokeMust(args ...any) []any` — invoke, panic on error
- `InvokeWithValidArgs() ([]any, error)` — invoke with all defined positional args
- `InvokeArgs(upTo int) ([]any, error)` — invoke with args up to position N

## FuncWrap[T]

A generic reflection-based function wrapper:

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

A flexible 6-slot holder where `T` types the `WorkFunc` field:

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

The package uses internal helper functions to reduce code duplication:

- `getByIndex(slice, index)` — safe indexed access
- `buildToString(typeName, slice, cache)` — cached string formatting
- `appendIfDefined(args, value)` — conditional append for defined values
- `invokeMustHelper(fw, args...)` — invoke with panic on error

## File Organization

| File | Purpose |
|---|---|
| `One.go`–`Six.go` | Positional arg holders (generic) |
| `OneFunc.go`–`SixFunc.go` | Func arg holders (generic) |
| `Holder.go` | Flexible 6-slot + WorkFunc holder |
| `FuncWrap.go` | Core generic function wrapper struct |
| `FuncWrapArgs.go` | Argument introspection methods |
| `FuncWrapInvoke.go` | Dynamic invocation methods |
| `FuncWrapValidation.go` | Validation and error methods |
| `FuncMap.go` | Named map of function wrappers |
| `Map.go` | Key-value argument map |
| `Dynamic.go` / `DynamicFunc.go` | Map-based dynamic holders |
| `LeftRight.go` | Two-item holder with Left/Right semantics |
| `aliases.go` | All `*Any` type aliases |
| `argsHelper.go` | Shared unexported utility functions |
| `all-interfaces.go` | Interface definitions |
| `consts.go` / `vars.go` | Package constants and variables |
