# corefuncs — Function Type Definitions

Package `corefuncs` defines reusable function type signatures for callbacks, processors, and functional composition. It includes both **legacy** (`any`-based) and **generic** (`[T]`-based) types.

## Architecture

```
corefuncs/
├── genericFuncs.go                            # Generic function types: InOutFuncOf[T,U], etc.
├── funcs.go                                   # Legacy function types: ExecFunc, InOutFunc, etc.
├── GetFunc.go / GetFuncName.go                # Runtime function name extraction
├── GetFuncFullName.go                         # Full qualified function name extraction
├── ActionReturnsErrorFuncWrapper.go           # Wrapper: func() error with name
├── InActionReturnsErrFuncWrapperOf.go         # Generic: func(T) error with name
├── InOutErrFuncWrapper.go                     # Wrapper: func(any) (any, error) with name
├── InOutErrFuncWrapperOf.go                   # Generic: func(T) (U, error) with name
├── InOutFuncWrapperOf.go                      # Generic: func(T) U with name + SerializeOutputFuncWrapperOf
├── IsSuccessFuncWrapper.go                    # Wrapper: func() bool with name
├── NamedActionFuncWrapper.go                  # Wrapper: func() with name
├── ResultDelegatingFuncWrapper.go             # Wrapper: func(any) error with name
├── ResultDelegatingFuncWrapperOf.go           # Generic: func(T) error with name
└── newCreator.go                              # New Creator pattern + package-level generic constructors
```

## Function Types

### Generic (Type-Safe) — Recommended

| Type | Signature |
|------|-----------|
| `InOutFuncOf[TIn, TOut]` | `func(TIn) TOut` |
| `InOutErrFuncOf[TIn, TOut]` | `func(TIn) (TOut, error)` |
| `SerializeOutputFuncOf[TIn]` | `func(TIn) ([]byte, error)` |
| `InActionReturnsErrFuncOf[TIn]` | `func(TIn) error` |
| `ResultDelegatingFuncOf[T]` | `func(T) error` |

### Legacy (any-based)

| Type | Signature |
|------|-----------|
| `ExecFunc` | `func()` |
| `ActionFunc` | `func()` |
| `IsBooleanFunc` | `func() bool` |
| `IsApplyFunc` | `func() bool` |
| `IsSuccessFunc` | `func() bool` |
| `IsFailureFunc` | `func() bool` |
| `InOutFunc` | `func(any) any` |
| `InOutErrFunc` | `func(any) (any, error)` |
| `ActionReturnsErrorFunc` | `func() error` |
| `ResultDelegatingFunc` | `func(any) error` |
| `InActionReturnsErrFunc` | `func(any) error` |
| `SerializeOutputFunc` | `func(any) ([]byte, error)` |
| `SerializerVoidFunc` | `func() ([]byte, error)` |
| `PayloadProcessorFunc` | `func([]byte) error` |
| `StringerActionFunc` | `func() string` |
| `StringerWithErrorActionFunc` | `func() (string, error)` |
| `NamedActionFunc` | `func(string)` |
| `NextReturnErrWrapperFunc` | `func(ActionReturnsErrorFunc) error` |
| `NextVoidActionFunc` | `func(ExecFunc)` |

### Named Wrappers

Named wrappers pair a function with a `Name` field for logging, tracing, and debugging. All provide `Exec()`, `AsActionFunc()`, `AsActionReturnsErrorFunc()`, and `ToLegacy()` methods.

| Wrapper | Inner Type | Description |
|---------|-----------|-------------|
| `ActionReturnsErrorFuncWrapper` | `func() error` | Named error-returning action |
| `InOutErrFuncWrapperOf[T, U]` | `func(T) (U, error)` | Generic named transform with error |
| `InOutFuncWrapperOf[T, U]` | `func(T) U` | Generic named pure transform |
| `InActionReturnsErrFuncWrapperOf[T]` | `func(T) error` | Generic named input-to-error action |
| `SerializeOutputFuncWrapperOf[T]` | `func(T) ([]byte, error)` | Generic named serializer |
| `ResultDelegatingFuncWrapperOf[T]` | `func(T) error` | Generic named delegating processor |
| `IsSuccessFuncWrapper` | `func() bool` | Named boolean check |
| `NamedActionFuncWrapper` | `func()` | Named void action |
| `InOutErrFuncWrapper` | `func(any) (any, error)` | Legacy named transform |
| `ResultDelegatingFuncWrapper` | `func(any) error` | Legacy named processor |

## Usage

### Generic Wrapper Constructors (Preferred)

Package-level constructor functions provide clean type inference:

```go
import "gitlab.com/auk-go/core/corefuncs"

// InOutErr — transform with error
wrapper := corefuncs.NewInOutErrWrapper[string, int](
    "parseAge",
    func(s string) (int, error) {
        return strconv.Atoi(s)
    },
)
result, err := wrapper.Exec("25") // result=25, err=nil

// InOut — pure transform (no error)
wrapper := corefuncs.NewInOutWrapper[string, int](
    "strlen",
    func(s string) int { return len(s) },
)
result := wrapper.Exec("hello") // result=5

// InActionErr — input action returning error
wrapper := corefuncs.NewInActionErrWrapper[string](
    "validate-email",
    func(email string) error { return validateEmail(email) },
)
err := wrapper.Exec("test@example.com")

// ResultDelegating — delegating to a target pointer
wrapper := corefuncs.NewResultDelegatingWrapper[*MyStruct](
    "unmarshal-user",
    func(target *MyStruct) error { return json.Unmarshal(data, target) },
)
var user MyStruct
err := wrapper.Exec(&user)

// Serialize — typed serializer
wrapper := corefuncs.NewSerializeWrapper[MyStruct](
    "json-marshal",
    func(m MyStruct) ([]byte, error) { return json.Marshal(m) },
)
bytes, err := wrapper.Exec(myStruct)
```

### Wrapper Methods

All generic wrappers provide:

```go
wrapper.Exec(input)                       // Execute with typed input
wrapper.AsActionFunc(input)               // Returns ActionFunc (panics on error)
wrapper.AsActionReturnsErrorFunc(input)   // Returns ActionReturnsErrorFunc
wrapper.ToLegacy()                        // Convert to any-based wrapper
wrapper.Name                              // Name string for tracing/logging
```

### Generic Function Types (Without Wrappers)

```go
// Strongly typed transformation
var transform corefuncs.InOutFuncOf[string, int] = func(s string) int {
    return len(s)
}
result := transform("hello") // 5 — compile-time safe

// Use in higher-order functions
func processAll[T, U any](items []T, fn corefuncs.InOutFuncOf[T, U]) []U {
    results := make([]U, len(items))
    for i, item := range items {
        results[i] = fn(item)
    }
    return results
}
```

### Legacy Wrapper Creators (via New)

```go
// Legacy (any-based) wrappers via New creator
legacyWrapper := corefuncs.New.ActionErr("cleanup", func() error {
    return os.Remove("/tmp/file")
})

successCheck := corefuncs.New.IsSuccess("healthcheck", func() bool {
    return ping() == nil
})

namedAction := corefuncs.New.NamedAction("log", func(name string) {
    fmt.Println("executing:", name)
})

legacyInOutErr := corefuncs.New.LegacyInOutErr("transform", func(input any) (any, error) {
    return strings.ToUpper(input.(string)), nil
})

legacyDelegating := corefuncs.New.LegacyResultDelegating("unmarshal", func(target any) error {
    return json.Unmarshal(data, target)
})
```

### Runtime Function Name Extraction

```go
name := corefuncs.GetFuncName(myFunc)         // "myFunc"
fullName := corefuncs.GetFuncFullName(myFunc)  // "package.myFunc"
```

## New Creator Pattern

The `New` variable provides a `newFuncCreator` with two categories:

**Legacy creators** (direct construction):
- `New.ActionErr(name, fn)` → `ActionReturnsErrorFuncWrapper`
- `New.IsSuccess(name, fn)` → `IsSuccessFuncWrapper`
- `New.NamedAction(name, fn)` → `NamedActionFuncWrapper`
- `New.LegacyInOutErr(name, fn)` → `InOutErrFuncWrapper`
- `New.LegacyResultDelegating(name, fn)` → `ResultDelegatingFuncWrapper`

**Generic constructors** (package-level functions — preferred over `New` for generics due to Go type inference limitations):
- `NewInOutErrWrapper[TIn, TOut](name, fn)` → `InOutErrFuncWrapperOf[TIn, TOut]`
- `NewInOutWrapper[TIn, TOut](name, fn)` → `InOutFuncWrapperOf[TIn, TOut]`
- `NewInActionErrWrapper[TIn](name, fn)` → `InActionReturnsErrFuncWrapperOf[TIn]`
- `NewResultDelegatingWrapper[T](name, fn)` → `ResultDelegatingFuncWrapperOf[T]`
- `NewSerializeWrapper[TIn](name, fn)` → `SerializeOutputFuncWrapperOf[TIn]`

> **Note**: Generic wrappers use package-level functions instead of `New.InOutErr().Of(...)` because Go cannot infer type parameters on struct method chains.

## Related Docs

- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
- [Folder Spec](/spec/01-app/folders/10-remaining-packages.md)
- [coredynamic README](/coredata/coredynamic/README.md)
