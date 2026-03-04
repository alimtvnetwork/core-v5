# coretests/args

The `args` package provides **typed argument holders** for structuring test case inputs, expected outputs, and dynamic function invocation in the testing framework.

## Architecture

```
args/
├── Positional Types
│   ├── One.go – Six.go            ← Generic 1–6 slot holders
│   ├── OneFunc.go – SixFunc.go    ← Same + WorkFunc for invocation
│   └── Holder.go                  ← 6-slot + typed WorkFunc + Hashmap
├── Function Wrapping
│   ├── FuncWrap.go                ← Core generic wrapper struct
│   ├── FuncWrapArgs.go            ← Argument introspection methods
│   ├── FuncWrapInvoke.go          ← Dynamic invocation methods
│   ├── FuncWrapValidation.go      ← Validation and error methods
│   ├── FuncWrapTypedHelpers.go    ← Signature checkers + typed invoke helpers
│   ├── newFuncWrapCreator.go      ← Factory methods (Default, Map, Many, etc.)
│   ├── FuncMap.go                 ← Named map of function wrappers
│   └── funcDetector.go            ← Function detection utilities
├── Map-Based Types
│   ├── Map.go                     ← Key-value argument map
│   ├── Dynamic.go                 ← Generic map-based dynamic holder (T = Expect type)
│   └── DynamicFunc.go             ← Generic map-based dynamic holder (T = WorkFunc type)
├── Support
│   ├── aliases.go                 ← *Any type aliases for backward compat
│   ├── all-interfaces.go          ← Interface definitions
│   ├── argsHelper.go              ← Shared unexported utilities
│   ├── LeftRight.go               ← Generic two-item holder with Left/Right semantics
│   ├── String.go                  ← String type helpers
│   ├── emptyCreator.go            ← Empty value factories
│   ├── toString.go                ← String conversion helper
│   ├── consts.go / vars.go        ← Package constants and variables
│   └── README.md                  ← This file
└── FuncWrap-README.md             ← Dedicated FuncWrap documentation
```

## When to Use What

### Decision Guide

| Scenario | Type to Use | Why |
|----------|-------------|-----|
| Test case with 1–3 typed inputs | `One[T]`, `Two[T1,T2]`, `Three[T1,T2,T3]` | Compile-time type safety on fields |
| Test case with function to invoke | `OneFunc[T]`, `TwoFunc[T1,T2]`, etc. | Holds both args and the function |
| Legacy test code or mixed types | `OneAny`, `TwoAny`, `ThreeAny` | Backward-compatible, all fields `any` |
| Need to invoke a function dynamically | `FuncWrap[T]` or `FuncWrapAny` | Reflection-based invocation + introspection |
| Key-value based test parameters | `Map` | Flexible named access with typed getters |
| Dynamic test with typed Expect | `Dynamic[T]` or `DynamicAny` | Map-based args with typed expected value |
| Dynamic test with typed WorkFunc | `DynamicFunc[T]` or `DynamicFuncAny` | Map-based args with typed function |
| Typed function holder with overflow params | `Holder[T]` | 6 slots + typed WorkFunc + Hashmap |
| Comparing two values (expected vs actual) | `LeftRight[TLeft, TRight]` or `LeftRightAny` | Typed Left/Right naming |

### Typed vs Untyped

**Use typed generics** when you know the argument types at compile time:

```go
// ✅ Typed — compiler catches type errors
tc := args.Two[string, int]{
    First:  "hello",
    Second: 42,
    Expect: "expected result",
}
```

**Use `*Any` aliases** for legacy code or when fields hold heterogeneous types:

```go
// ✅ Untyped — flexible, backward-compatible
tc := args.TwoAny{
    First:  someInterface,
    Second: anotherInterface,
    Expect: "expected result",
}
```

> **Rule**: Since `*Any` aliases use `=` (type aliases, not new types), Go's `%T`
> reflection output remains the base type name (e.g., `args.Two`, not `args.TwoAny`).
> This is important for `ExpectedInput` strings that include type assertions.

## Positional Types (One–Six)

Hold 1–6 arguments plus an optional `Expect` field. Each positional field
is parameterized with its own type parameter.

### Struct Definition

```go
type Three[T1, T2, T3 any] struct {
    First  T1
    Second T2
    Third  T3
    Expect any  // expected output (always any)
}
```

### Common Methods

All positional types implement `ArgBaseContractsBinder`:

| Method | Returns | Description |
|--------|---------|-------------|
| `FirstItem()` | `any` | Returns First as any (interface-compatible) |
| `HasFirst()` | `bool` | Checks if First is defined (non-nil, non-zero) |
| `Expected()` | `any` | Returns the Expect field |
| `HasExpect()` | `bool` | Checks if Expect is defined |
| `ValidArgs()` | `[]any` | Collects all defined arguments (skips nil/zero) |
| `Args(upTo)` | `[]any` | Collects arguments up to position N |
| `Slice()` | `[]any` | All fields as a cached slice (includes Expect if defined) |
| `GetByIndex(i)` | `any` | Safe indexed access, returns nil if out of bounds |
| `String()` | `string` | Formatted: `"Three { val1, val2, val3 }"` |
| `ArgsCount()` | `int` | Number of positional slots (not counting Expect) |

### Downcast Methods

Convert to smaller arg types while preserving type parameters:

```go
three := args.Three[string, int, bool]{First: "a", Second: 1, Third: true}
two := three.ArgTwo()    // Two[string, int]{First: "a", Second: 1}
one := three.ArgOne()    // One[string]{First: "a"}
```

### Caching Behavior

All positional types use an `isSliceCached` flag with lazy evaluation for the
internal `[]any` slice representation. The `Slice()` method builds the slice
once on first call and returns the cached version thereafter.

### Usage in Test Cases

```go
var testCases = []coretestcases.CaseV1{
    {
        Title: "addition of two positive integers",
        ArrangeInput: args.TwoAny{
            First:  5,
            Second: 3,
        },
        ExpectedInput: "8",
    },
    {
        Title: "addition with zero",
        ArrangeInput: args.TwoAny{
            First:  0,
            Second: 7,
        },
        ExpectedInput: "7",
    },
}
```

```go
func Test_Add(t *testing.T) {
    for caseIndex, tc := range testCases {
        // Arrange
        input := tc.ArrangeInput.(args.TwoAny)

        // Act
        result := fmt.Sprintf("%d", Add(input.First.(int), input.Second.(int)))

        // Assert
        tc.ShouldBeEqual(t, caseIndex, result)
    }
}
```

## Func Types (OneFunc–SixFunc)

Same as positional types but include a `WorkFunc any` field for dynamic
function invocation. The positional arguments are typed, while `WorkFunc`
remains `any` because it requires reflection-based invocation via `FuncWrapAny`.

### When to Use Func Types

Use these when:
- Your test cases need to invoke different functions with the same argument structure
- You want to parameterize both the function and its arguments in the test data
- You need `FuncWrap` introspection (arg types, return types) per test case

```go
tc := args.TwoFunc[string, int]{
    First:    "input1",
    Second:   42,
    WorkFunc: myFunction,  // always any — for reflection
    Expect:   "expected",
}
```

### Invocation Methods

| Method | Returns | Description |
|--------|---------|-------------|
| `FuncWrap()` | `*FuncWrapAny` | Wraps WorkFunc for reflection |
| `Invoke(args...)` | `([]any, error)` | Invoke with explicit args |
| `InvokeMust(args...)` | `[]any` | Invoke, panic on error |
| `InvokeWithValidArgs()` | `([]any, error)` | Invoke with all defined positional args |
| `InvokeArgs(upTo)` | `([]any, error)` | Invoke with args up to position N |
| `GetWorkFunc()` | `any` | Returns the raw WorkFunc |
| `GetFuncName()` | `string` | Returns the function name via reflection |

### Usage Pattern

```go
var testCases = []coretestcases.CaseV1{
    {
        Title: "isany.Null returns true for nil",
        ArrangeInput: args.OneFuncAny{
            First:    nil,
            WorkFunc: isany.Null,
        },
        ExpectedInput: "true",
    },
}

func Test_IsAny(t *testing.T) {
    for caseIndex, tc := range testCases {
        input := tc.ArrangeInput.(args.OneFuncAny)
        checkerFunc := input.WorkFunc.(func(any) bool)

        result := fmt.Sprintf("%v", checkerFunc(input.First))

        tc.ShouldBeEqual(t, caseIndex, result)
    }
}
```

## FuncWrap[T]

A generic reflection-based function wrapper. See **[FuncWrap-README.md](FuncWrap-README.md)**
for comprehensive documentation including all methods, creation patterns, and usage examples.

Quick reference:

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

A flexible 6-slot holder where `T` types the `WorkFunc` field. Positional
fields (First through Sixth) remain `any` for maximum flexibility. Includes
a `Hashmap` for overflow parameters.

### When to Use Holder

Use `Holder[T]` when:
- You need more than 6 arguments (use Hashmap for overflow)
- You want a typed `WorkFunc` (unlike Func types where WorkFunc is `any`)
- You need a single flexible container for complex test setups

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
    Hashmap:  args.Map{"timeout": 30, "retries": 3},
}
```

## Map

A `map[string]any` type with typed getter methods for extracting values
by key name. Used as `ArrangeInput` when tests need named parameters
rather than positional ones.

### Key Methods

| Method | Returns | Description |
|--------|---------|-------------|
| `Get(key)` | `(any, bool)` | Raw map access |
| `GetAsInt(key)` | `(int, error)` | Type-safe int extraction |
| `GetAsString(key)` | `(string, error)` | Type-safe string extraction |
| `GetAsBool(key)` | `(bool, error)` | Type-safe bool extraction |
| `WorkFunc()` | `any` | Returns the value at key "func" |
| `HasFunc()` | `bool` | Checks if "func" key exists |
| `HasExpect()` | `bool` | Checks if "expected" key exists |
| `FirstItem()` | `any` | Returns value at "first" key |
| `ArgsCount()` | `int` | Count of keys excluding "expected" and "func" |

### Usage Pattern

```go
var testCases = []coretestcases.CaseV1{
    {
        Title: "MaxInt returns larger value",
        ArrangeInput: args.Map{
            "a": 5,
            "b": 3,
        },
        ExpectedInput: "5",
    },
}

func Test_MaxInt(t *testing.T) {
    for caseIndex, tc := range testCases {
        input := tc.ArrangeInput.(args.Map)
        a, _ := input.GetAsInt("a")
        b, _ := input.GetAsInt("b")

        result := coremath.MaxInt(a, b)

        tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
    }
}
```

## Dynamic[T] / DynamicFunc[T]

Generic map-based argument holders for fully dynamic test scenarios.

- **`Dynamic[T]`** — `T` parameterizes the `Expect` field. Use `DynamicAny` (`= Dynamic[any]`) for untyped usage.
- **`DynamicFunc[T]`** — `T` parameterizes the `WorkFunc` field (following the `Holder[T]` pattern). Use `DynamicFuncAny` (`= DynamicFunc[any]`) for untyped usage.

### When to Use Dynamic Types

Use these when:
- Arguments are purely key-value based (no positional structure)
- The function to invoke is part of the test data
- You need maximum flexibility at the cost of type safety

```go
// Typed Expect
tc := args.Dynamic[string]{
    Params: args.Map{
        "first":  "hello",
        "second": 42,
    },
    Expect: "expected",
}

// Untyped (backward-compatible)
tc := args.DynamicAny{
    Params: args.Map{
        "first":  "hello",
        "second": 42,
    },
    Expect: 123,
}

// Typed WorkFunc
tc := args.DynamicFunc[func(string) error]{
    Params: args.Map{
        "input": "test",
    },
    WorkFunc: myProcessor,
    Expect:   nil,
}

// Untyped (backward-compatible)
tc := args.DynamicFuncAny{
    Params: args.Map{
        "input": "test",
    },
    WorkFunc: myFunc,
    Expect:   "result",
}

results, err := tc.InvokeWithValidArgs()
```

### Dynamic Methods

Both `Dynamic[T]` and `DynamicFunc[T]` delegate to their `Params` map and share
the same rich method set including `Get`, `GetAsInt`, `GetAsString`, `HasDefined`,
`HasDefinedAll`, `IsKeyInvalid`, `IsKeyMissing`, `ValidArgs`, `Invoke`, and more.

`DynamicFunc[T]` additionally exposes `GetWorkFunc()`, `HasFunc()`, `GetFuncName()`,
and `FuncWrap()` for typed function access.

## LeftRight[TLeft, TRight]

A generic two-item holder with Left/Right semantics, providing a semantic
alternative to `Two` for cases where the directionality of arguments matters.

### Type Parameters

- `TLeft` — type of the Left field
- `TRight` — type of the Right field
- Use `LeftRightAny` (`= LeftRight[any, any]`) for untyped usage

```go
// Typed
lr := args.LeftRight[string, int]{
    Left:   "expected",
    Right:  42,
    Expect: true,
}

// Untyped (backward-compatible)
lr := args.LeftRightAny{
    Left:   someValue,
    Right:  anotherValue,
    Expect: "match",
}
```

### LeftRight Methods

| Method | Returns | Description |
|--------|---------|-------------|
| `FirstItem()` / `SecondItem()` | `any` | Returns Left / Right as any |
| `HasFirst()` / `HasSecond()` | `bool` | Checks if Left / Right is defined |
| `HasLeft()` / `HasRight()` | `bool` | Semantic aliases for HasFirst / HasSecond |
| `ArgTwo()` | `TwoFuncAny` | Converts to TwoFuncAny |
| `Clone()` | `LeftRight[TLeft, TRight]` | Returns an independent typed copy |
| `ValidArgs()` | `[]any` | All defined positional args |
| `Slice()` | `[]any` | All fields as a cached slice |

## Generic Type Aliases

Every generic type has a corresponding `*Any` alias in `aliases.go`:

| Generic Type | Any Alias | Use Case |
|---|---|---|
| `FuncWrap[T]` | `FuncWrapAny` | Dynamic function wrapping |
| `One[T1]` | `OneAny` | Single-arg test cases |
| `Two[T1, T2]` | `TwoAny` | Two-arg test cases |
| `Three[T1, T2, T3]` | `ThreeAny` | Three-arg test cases |
| `Four[T1, T2, T3, T4]` | `FourAny` | Four-arg test cases |
| `Five[T1, T2, T3, T4, T5]` | `FiveAny` | Five-arg test cases |
| `Six[T1, T2, T3, T4, T5, T6]` | `SixAny` | Six-arg test cases |
| `OneFunc[T1]` | `OneFuncAny` | Single-arg with function |
| `TwoFunc[T1, T2]` | `TwoFuncAny` | Two-arg with function |
| `ThreeFunc[T1, T2, T3]` | `ThreeFuncAny` | Three-arg with function |
| `FourFunc[T1, T2, T3, T4]` | `FourFuncAny` | Four-arg with function |
| `FiveFunc[T1, T2, T3, T4, T5]` | `FiveFuncAny` | Five-arg with function |
| `SixFunc[T1, T2, T3, T4, T5, T6]` | `SixFuncAny` | Six-arg with function |
| `Holder[T]` | `HolderAny` | Flexible holder |
| `LeftRight[TLeft, TRight]` | `LeftRightAny` | Two-item Left/Right holder |
| `Dynamic[T]` | `DynamicAny` | Map-based dynamic holder |
| `DynamicFunc[T]` | `DynamicFuncAny` | Map-based dynamic holder with func |

> **Important**: These are type aliases (`=`), NOT new types. Go's `%T` reflection
> output shows the base name (e.g., `args.Two`), not the alias name (`args.TwoAny`).

## Shared Helpers

Internal helper functions in `argsHelper.go` reduce code duplication:

| Helper | Purpose |
|--------|---------|
| `getByIndex(slice, index)` | Safe indexed access, returns nil if out of bounds |
| `buildToString(typeName, slice, cache)` | Cached `"TypeName { val1, val2 }"` string formatting |
| `appendIfDefined(args, value)` | Conditional append — only adds non-nil/non-zero values |
| `invokeMustHelper(fw, args...)` | Invoke with panic on error (eliminates duplicate InvokeMust patterns) |

## Interface Hierarchy

```
ArgBaseContractsBinder          ← Core: item access, validation, slicing, String()
├── OneParameter                ← Single arg + AsArgBaseContractsBinder
│   ├── TwoParameter            ← + SecondItem()
│   │   ├── ThreeParameter      ← + ThirdItem()
│   │   │   ├── FourParameter   ← + FourthItem()
│   │   │   │   ├── FifthParameter  ← + FifthItem()
│   │   │   │   │   └── SixthParameter  ← + SixthItem()

ArgFuncContractsBinder          ← Base + FuncNumber
├── OneFuncParameter            ← OneParameter + FuncNumber
│   ├── TwoFuncParameter        ← TwoParameter + FuncNumber
│   │   ├── ThreeFuncParameter  ← etc.
│   │   │   └── ... up to SixthFuncParameter

ArgsMapper                      ← Map-based: ArgBase + FuncNamer + named getters
FuncWrapper                     ← Full FuncWrap contract
```

## Design Decisions

### Pointer-to-Slice Removal

All types use `[]any` + `bool` flag for slice caching instead of `*[]any`.
This follows the project's pointer optimization standards for simpler API
and better Go memory efficiency.

### WorkFunc Typing

In Func variants (OneFunc–SixFunc), the `WorkFunc` field remains `any`
because it requires reflection-based invocation via `FuncWrapAny`.
`Holder[T]` and `DynamicFunc[T]` parameterize WorkFunc with type `T`
for typed function holder patterns.

In `FuncWrap[T]`, the `Func` field is typed as `T`, enabling both
typed (`NewTypedFuncWrap`) and untyped (`NewFuncWrap.Default`) construction.

`Dynamic[T]` parameterizes the `Expect` field with `T`, allowing
typed expected values while keeping `Params` as a flexible `Map`.

`LeftRight[TLeft, TRight]` parameterizes both the `Left` and `Right`
fields independently, enabling typed comparisons between heterogeneous values.

## Related Docs

- [FuncWrap-README.md](FuncWrap-README.md) — Detailed FuncWrap documentation
- [spec/01-app/16-testing-guidelines.md](/spec/01-app/16-testing-guidelines.md) — Testing guidelines
- [coretests/](/coretests/) — Parent testing framework
