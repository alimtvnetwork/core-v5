# corefuncs — Function Type Definitions

Package `corefuncs` defines reusable function type signatures for callbacks, processors, and functional composition. It includes both **legacy** (`any`-based) and **generic** (`[T]`-based) types.

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

## Related Docs

- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
- [Folder Spec](/spec/01-app/folders/10-remaining-packages.md)
