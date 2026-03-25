# FuncWrap[T] — Reflection-Based Function Wrapper

`FuncWrap[T]` is a generic wrapper that enables dynamic invocation, argument
validation, type introspection, and equality comparison for Go functions at runtime.

## Architecture

```
FuncWrap[T]
├── FuncWrap.go                ← Core struct, creation, identity methods
├── FuncWrapArgs.go            ← Argument/return type introspection
├── FuncWrapInvoke.go          ← Dynamic invocation methods
├── FuncWrapValidation.go      ← Validation, error reporting
└── newFuncWrapCreator.go      ← Factory methods (Default, Map, Many, etc.)
```

## When to Use FuncWrap

| Scenario | Approach |
|----------|----------|
| Invoking a function whose signature is known at compile time | Call it directly — no FuncWrap needed |
| Invoking a function stored as `any` in test data | `FuncWrapAny` via `NewFuncWrap.Default(fn)` |
| Typed function wrapping with compile-time safety on the Func field | `NewTypedFuncWrap(fn)` |
| Building a registry of named functions | `NewFuncWrap.Map(fn1, fn2, fn3)` |
| Introspecting argument types, names, counts | Any FuncWrap — use Args methods |
| Comparing two functions for structural equality | `fw.IsEqual(other)` |
| Converting struct methods to invocable wrappers | `NewFuncWrap.StructToMap(myStruct)` |

## Creation

### Typed Construction

Use when you know the function type at compile time:

```go
fw := args.NewTypedFuncWrap(func(s string) int { return len(s) })

// fw.Func is typed as func(string) int
// Compile-time error if you pass a non-function
```

### Untyped Construction (via Creator)

Use when the function comes from test data or is stored as `any`:

```go
// Single function
fw := args.NewFuncWrap.Default(myFunc)

// From an existing FuncWrap or FuncWrapGetter — returns as-is
fw := args.NewFuncWrap.Default(existingFuncWrap)

// Invalid (nil or non-function) — returns invalid FuncWrap
fw := args.NewFuncWrap.Default(nil)
fw.IsInvalid() // true
```

### Batch Creation

```go
// Named map — keys are function names
funcMap := args.NewFuncWrap.Map(fn1, fn2, fn3)
// Result: map[string]FuncWrapAny{"fn1": ..., "fn2": ..., "fn3": ...}

// Slice of pointers
funcs := args.NewFuncWrap.Many(fn1, fn2, fn3)
// Result: []*FuncWrapAny

// From struct methods
funcMap, err := args.NewFuncWrap.StructToMap(myStruct)
// Result: map[string]FuncWrapAny with all public methods

// From reflect.Method
fw, err := args.NewFuncWrap.MethodToFunc(reflectMethod)
```

### Invalid FuncWrap

```go
fw := args.NewFuncWrap.Invalid()
fw.IsInvalid() // true
fw.IsValid()   // false
```

## Struct Fields

```go
type FuncWrap[T any] struct {
    Name     string        // Short function name (e.g., "MyFunc")
    FullName string        // Full qualified name (e.g., "pkg.MyFunc")
    Func     T             // The wrapped function value (typed)
    // unexported: isInvalid, rvType, rv, cached arg info...
}
```

## Method Reference

### Identity & Validation

| Method | Signature | Description |
|--------|-----------|-------------|
| `GetFuncName()` | `string` | Short name of the wrapped function |
| `GetPascalCaseFuncName()` | `string` | Function name in PascalCase |
| `HasValidFunc()` | `bool` | True if holds a valid, callable function |
| `IsValid()` | `bool` | True if valid (inverse of IsInvalid) |
| `IsInvalid()` | `bool` | True if nil, marked invalid, or no valid reflect |
| `MustBeValid()` | — | Panics if invalid |
| `ValidationError()` | `error` | Returns error describing invalidity, or nil |
| `InvalidError()` | `error` | Detailed error about why invalid, or nil |
| `GetType()` | `reflect.Type` | The reflect.Type of the function |
| `IsPublicMethod()` | `bool` | True if exported (PkgPath is empty) |
| `IsPrivateMethod()` | `bool` | True if unexported |
| `PkgPath()` | `string` | Full package path of the function |
| `PkgNameOnly()` | `string` | Package name without path |
| `FuncDirectInvokeName()` | `string` | Name suitable for code generation |

### Argument Introspection

| Method | Signature | Description |
|--------|-----------|-------------|
| `ArgsCount()` | `int` | Number of input parameters (-1 if invalid) |
| `InArgsCount()` | `int` | Alias for ArgsCount |
| `ArgsLength()` | `int` | Alias for ArgsCount |
| `OutArgsCount()` | `int` | Number of return values (-1 if invalid) |
| `ReturnLength()` | `int` | Alias for OutArgsCount |
| `GetInArgsTypes()` | `[]reflect.Type` | reflect.Type for each input param (cached) |
| `GetOutArgsTypes()` | `[]reflect.Type` | reflect.Type for each return value (cached) |
| `GetInArgsTypesNames()` | `[]string` | String name of each input type (cached) |
| `GetOutArgsTypesNames()` | `[]string` | String name of each return type (cached) |
| `InArgNames()` | `[]string` | Generated variable names from type names |
| `OutArgNames()` | `[]string` | Generated variable names for return values |
| `InArgNamesEachLine()` | `SimpleSlice` | Input arg names formatted one per line |
| `OutArgNamesEachLine()` | `SimpleSlice` | Output arg names formatted one per line |

### Dynamic Invocation

| Method | Signature | Description |
|--------|-----------|-------------|
| `Invoke(args...)` | `([]any, error)` | Invoke with given arguments; validates types first |
| `InvokeMust(args...)` | `[]any` | Invoke, panic on error |
| `InvokeSkip(skip, args...)` | `([]any, error)` | Invoke with custom stack skip for error reporting |
| `VoidCall()` | `([]any, error)` | Invoke with no arguments |
| `VoidCallNoReturn(args...)` | `error` | Invoke ignoring return values |
| `GetFirstResponseOfInvoke(args...)` | `(any, error)` | Invoke and return only first result |
| `InvokeResultOfIndex(i, args...)` | `(any, error)` | Invoke and return result at index |
| `InvokeError(args...)` | `(funcErr, processingErr)` | Invoke and cast first result to error |
| `InvokeFirstAndError(args...)` | `(first, funcErr, processingErr)` | For `(T, error)` return patterns |

### Type Verification

| Method | Signature | Description |
|--------|-----------|-------------|
| `IsInTypeMatches(args...)` | `bool` | Check if given args match input types |
| `IsOutTypeMatches(outArgs...)` | `bool` | Check if given args match output types |
| `VerifyInArgs(args)` | `(bool, error)` | Verify input args with error details |
| `VerifyOutArgs(args)` | `(bool, error)` | Verify output args with error details |
| `ValidateMethodArgs(args)` | `error` | Validate count + types, detailed error message |
| `InArgsVerifyRv(types)` | `(bool, error)` | Verify using `[]reflect.Type` |
| `OutArgsVerifyRv(types)` | `(bool, error)` | Verify using `[]reflect.Type` |

### Equality

| Method | Signature | Description |
|--------|-----------|-------------|
| `IsEqual(other)` | `bool` | Deep equality: validity, name, visibility, arg counts, arg types |
| `IsNotEqual(other)` | `bool` | Inverse of IsEqual |
| `IsEqualValue(other)` | `bool` | Compare using value (non-pointer) FuncWrap |

## Invocation Flow

When `Invoke(args...)` is called, the following steps occur:

1. **Validation check** — `ValidationError()` ensures FuncWrap is valid
2. **Argument validation** — `ValidateMethodArgs(args)` checks count and type compatibility
3. **Reflection call** — Args are converted to `reflect.Value` and `rv.Call()` is invoked
4. **Panic recovery** — `trydo.WrapPanic` catches any panics during invocation
5. **Error formatting** — On failure, a detailed error with function name, stack trace, and exception message is returned
6. **Result conversion** — `reflect.Value` results are converted back to `[]any`

```
Invoke(args...)
  │
  ├─ ValidationError()     → nil? continue : return error
  ├─ ValidateMethodArgs()  → nil? continue : return error
  ├─ argsToRvFunc()        → convert []any to []reflect.Value
  ├─ trydo.WrapPanic(rv.Call(rvs))
  │   ├─ success → rvToInterfacesFunc(results), nil
  │   └─ panic   → nil, formatted error with stack trace
  └─ return results, error
```

## Error Messages

When argument validation fails, `FuncWrap` produces detailed diagnostic messages:

```
myFunc [Func] =>
  arguments count doesn't match for - count:
    expected : 2
    given    : 3
  expected types listed :
    - string
    - int
  actual given types list :
    - string (value: "hello")
    - int (value: 42)
    - bool (value: true)
```

## Usage Examples

### Basic Introspection

```go
fw := args.NewTypedFuncWrap(strings.Contains)

fmt.Println(fw.GetFuncName())         // "Contains"
fmt.Println(fw.ArgsCount())           // 2
fmt.Println(fw.ReturnLength())        // 1
fmt.Println(fw.GetInArgsTypesNames()) // ["string", "string"]
fmt.Println(fw.IsPublicMethod())      // true
```

### Dynamic Test Invocation

```go
// In test case data
testCases := []struct {
    Func     any
    Args     []any
    Expected any
}{
    {Func: strings.ToUpper, Args: []any{"hello"}, Expected: "HELLO"},
    {Func: strconv.Itoa,    Args: []any{42},      Expected: "42"},
}

for _, tc := range testCases {
    fw := args.NewFuncWrap.Default(tc.Func)
    results, err := fw.Invoke(tc.Args...)

    if err != nil { t.Fatal(err) }
    if results[0] != tc.Expected {
        t.Errorf("got %v, want %v", results[0], tc.Expected)
    }
}
```

### Building a Function Registry

```go
registry := args.NewFuncWrap.Map(
    strings.ToUpper,
    strings.ToLower,
    strings.TrimSpace,
)

// Access by name
upper := registry["ToUpper"]
results, _ := upper.Invoke("hello")
fmt.Println(results[0]) // "HELLO"
```

### Struct Method Extraction

```go
type Calculator struct{}
func (c Calculator) Add(a, b int) int { return a + b }
func (c Calculator) Mul(a, b int) int { return a * b }

methods, err := args.NewFuncWrap.StructToMap(Calculator{})
// methods["Add"], methods["Mul"] are FuncWrapAny instances

addFw := methods["Add"]
results, _ := addFw.Invoke(Calculator{}, 3, 4) // receiver + args
fmt.Println(results[0]) // 7
```

### Equality Comparison

```go
fw1 := args.NewFuncWrap.Default(strings.Contains)
fw2 := args.NewFuncWrap.Default(strings.Contains)
fw3 := args.NewFuncWrap.Default(strings.HasPrefix)

fmt.Println(fw1.IsEqual(fw2)) // true  — same name, args, returns
fmt.Println(fw1.IsEqual(fw3)) // false — different name
```

## Related Docs

- [README.md](README.md) — Full args package documentation
- [spec/01-app/16-testing-guidelines.md](/spec/01-app/16-testing-guidelines.md) — Testing guidelines
- [coretests/](/coretests/) — Parent testing framework
