# corefuncs — Function Type Definitions

Package `corefuncs` defines reusable function type signatures for callbacks, processors, and functional composition. It includes both **legacy** (`any`-based) and **generic** (`[T]`-based) types.

## Architecture

```
corefuncs/
├── genericFuncs.go                            # Generic function types: InOutFuncOf[T,U], etc.
├── funcs.go                                   # Legacy function types: ExecFunc, InOutFunc, etc.
├── GetFunc.go / GetFuncName.go                # Runtime function name extraction
├── ActionReturnsErrorFuncWrapper.go           # Wrapper: func() error with name
├── InActionReturnsErrFuncWrapperOf.go         # Generic: func(T) error with name
├── InOutErrFuncWrapper.go                     # Wrapper: func(any) (any, error) with name
├── InOutErrFuncWrapperOf.go                   # Generic: func(T) (U, error) with name
├── InOutFuncWrapperOf.go                      # Generic: func(T) U with name
├── IsSuccessFuncWrapper.go                    # Wrapper: func() bool with name
├── NamedActionFuncWrapper.go                  # Wrapper: func() with name
├── ResultDelegatingFuncWrapper.go             # Wrapper: func(any) error with name
├── ResultDelegatingFuncWrapperOf.go           # Generic: func(T) error with name
└── newCreator.go                              # New Creator pattern
```

## Function Types

### Generic (Type-Safe) — Recommended

| Type | Signature |
|------|-----------|
| `InOutFuncOf[TIn, TOut]` | `func(TIn) TOut` |
| `InOutErrFuncOf[TIn, TOut]` | `func(TIn) (TOut, error)` |
| `SerializeOutputFuncOf[TIn]` | `func(TIn) ([]byte, error)` |

### Legacy (any-based)

| Type | Signature |
|------|-----------|
| `ExecFunc` | `func()` |
| `ActionFunc` | `func()` |
| `IsBooleanFunc` | `func() bool` |
| `IsApplyFunc` | `func() bool` |
| `InOutFunc` | `func(any) any` |
| `InOutErrFunc` | `func(any) (any, error)` |
| `ActionReturnsErrorFunc` | `func() error` |
| `ResultDelegatingFunc` | `func(any) error` |
| `PayloadProcessorFunc` | `func([]byte) error` |
| `StringerActionFunc` | `func() string` |

### Named Wrappers

Named wrappers pair a function with a name for logging, tracing, and debugging:

| Wrapper | Inner Type | Description |
|---------|-----------|-------------|
| `ActionReturnsErrorFuncWrapper` | `func() error` | Named error-returning action |
| `InOutErrFuncWrapperOf[T, U]` | `func(T) (U, error)` | Generic named transform |
| `InOutFuncWrapperOf[T, U]` | `func(T) U` | Generic named pure transform |
| `IsSuccessFuncWrapper` | `func() bool` | Named boolean check |
| `NamedActionFuncWrapper` | `func()` | Named void action |
| `ResultDelegatingFuncWrapperOf[T]` | `func(T) error` | Generic named processor |

## Usage

### Generic Function Types

```go
import "gitlab.com/auk-go/core/corefuncs"

// Strongly typed transformation
var transform corefuncs.InOutFuncOf[string, int] = func(s string) int {
    return len(s)
}

result := transform("hello") // 5 — compile-time safe

// Serializer
var serialize corefuncs.SerializeOutputFuncOf[MyStruct] = func(m MyStruct) ([]byte, error) {
    return json.Marshal(m)
}

// Use in higher-order functions
func processAll[T, U any](items []T, fn corefuncs.InOutFuncOf[T, U]) []U {
    results := make([]U, len(items))
    for i, item := range items {
        results[i] = fn(item)
    }
    return results
}
```

### Named Function Wrappers

```go
// Package-level generic constructors (preferred — clean type inference)
wrapper := corefuncs.NewInOutErrWrapper[string, int](
    "parseAge",
    func(s string) (int, error) {
        return strconv.Atoi(s)
    },
)

fmt.Println(wrapper.Name)           // "parseAge"
result, err := wrapper.Action("25") // 25, nil

// Legacy (any-based) wrappers via New creator
legacyWrapper := corefuncs.New.ActionErr("cleanup", func() error {
    return os.Remove("/tmp/file")
})
```

### Legacy Function Types

```go
var exec corefuncs.ExecFunc = func() {
    fmt.Println("executed")
}

var check corefuncs.IsBooleanFunc = func() bool {
    return true
}

var transform corefuncs.InOutFunc = func(input any) any {
    return strings.ToUpper(input.(string))
}
```

### Runtime Function Name Extraction

```go
name := corefuncs.GetFuncName(myFunc)     // "myFunc"
fullName := corefuncs.GetFuncFullName(myFunc) // "package.myFunc"
```

## Related Docs

- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
- [Folder Spec](/spec/01-app/folders/10-remaining-packages.md)
- [coredynamic README](/coredata/coredynamic/README.md)
