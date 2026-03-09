# 03 — Args Reference (`coretests/args/`)

The `args` package provides typed input holders for test cases. All types are generic with `any`-based aliases for backward compatibility.

---

## args.Map

The **primary** input/output holder. A `map[string]any` with typed accessors.

### Declaration

```go
type Map map[string]any
```

### Creating

```go
// In ArrangeInput (test input):
ArrangeInput: args.Map{
    "name":    "Alice",
    "age":     30,
    "isAdmin": true,
}

// In ExpectedInput (test expectations):
ExpectedInput: args.Map{
    "isValid":  true,
    "fullName": "Alice Smith",
    "count":    1,
}

// In test runner (actual results):
actual := args.Map{
    "isValid":  result.IsValid(),
    "fullName": result.FullName(),
    "count":    result.Count(),
}
```

### Typed Getters (return `(T, bool)`)

| Method | Returns | Example |
|--------|---------|---------|
| `GetAsString("key")` | `(string, bool)` | `msg, ok := input.GetAsString("message")` |
| `GetAsInt("key")` | `(int, bool)` | `n, ok := input.GetAsInt("count")` |
| `GetAsBool("key")` | `(bool, bool)` | `flag, ok := input.GetAsBool("isEnabled")` |
| `GetAsStrings("key")` | `([]string, bool)` | `items, ok := input.GetAsStrings("tags")` |
| `GetAsAnyItems("key")` | `([]any, bool)` | `items, ok := input.GetAsAnyItems("data")` |
| `Get("key")` | `(any, bool)` | `val, ok := input.Get("anything")` |

### Default Getters

| Method | Returns | Example |
|--------|---------|---------|
| `GetAsStringDefault("key")` | `string` | Returns `""` if missing |
| `GetAsIntDefault("key", 0)` | `int` | Returns default if missing |
| `GetAsBoolDefault("key", false)` | `bool` | Returns default if missing |

### Presence Checks

| Method | Meaning |
|--------|---------|
| `Has("key")` | Key exists (value may be nil) |
| `HasDefined("key")` | Key exists AND value is non-nil |
| `HasDefinedAll("a", "b")` | All keys exist and are non-nil |
| `IsKeyMissing("key")` | Key does not exist |
| `IsKeyInvalid("key")` | Key missing OR value is nil |

### Positional Access (aliases for common keys)

| Method | Checks keys (in order) |
|--------|----------------------|
| `FirstItem()` | `"first"`, `"f1"`, `"p1"`, `"1"` |
| `SecondItem()` | `"second"`, `"f2"`, `"p2"`, `"2"` |
| `ThirdItem()` | `"third"`, `"f3"`, `"p3"`, `"3"` |
| `WorkFunc()` | `"func"`, `"work.func"`, `"workFunc"` |
| `Expected()` | `"expected"`, `"expects"`, `"expect"` |

### Function Invocation

```go
// Direct invocation via reflection:
results, err := input.Invoke(arg1, arg2)

// Invoke with all defined (non-func) values:
results, err := input.InvokeWithValidArgs()

// Invoke with specific named args:
results, err := input.InvokeArgs("first", "second")
```

### CompileToStrings (for assertions)

Converts map to sorted `"key : value"` lines using `%v` formatting:

```go
m := args.Map{"isZero": false, "value": 5}
m.CompileToStrings()
// Returns: []string{"isZero : false", "value : 5"}
```

This is what `ShouldBeEqualMap` uses internally to compare expected vs actual.

### Formatting Rules

```go
// ✅ GOOD — multi-line for 2+ entries:
ArrangeInput: args.Map{
    "name": "Alice",
    "age":  30,
}

// ✅ OK — inline for single entry:
ExpectedInput: args.Map{"isValid": true}

// ❌ BAD — inline for 2+ entries:
ExpectedInput: args.Map{"isValid": true, "count": 5}
```

---

## args.One[T1] through args.Six[T1...T6]

Positional argument holders for 1–6 typed parameters.

### Structure (args.One example)

```go
type One[T1 any] struct {
    First  T1
    Expect any
}
```

### Aliases

```go
type OneAny   = One[any]
type TwoAny   = Two[any, any]
type ThreeAny = Three[any, any, any]
// ... through SixAny
```

### Key Methods

| Method | Exists on | Purpose |
|--------|-----------|---------|
| `FirstItem()` | All | Returns `First` as `any` |
| `SecondItem()` | Two+ | Returns `Second` as `any` |
| `HasFirst()` | All | Non-nil check |
| `ValidArgs()` | All | Returns all defined args as `[]any` |
| `Args(upTo)` | All | Returns first N args |
| `Expected()` | All | Returns `Expect` field |

### When to Use

Use in `ArrangeInput` when test data is fundamentally positional (not key-value):

```go
// ✅ GOOD — positional data in ArrangeInput:
ArrangeInput: &args.TwoAny{
    First:  "hello",
    Second: 42,
    Expect: "hello-42",
}

// ❌ BAD — positional data in ExpectedInput (use args.Map instead):
ExpectedInput: &args.OneAny{First: "result"}
```

**Rule**: Positional types are **allowed in ArrangeInput** but **prohibited in ExpectedInput**. Always use `args.Map` for expectations.

---

## args.Dynamic[T]

A map-based argument holder with a typed `Expect` field. Combines the flexibility of `args.Map` with typed expectations.

### Structure

```go
type Dynamic[T any] struct {
    Params Map
    Expect T
}
```

### Alias

```go
type DynamicAny = Dynamic[any]
```

### When to Use

When the input is inherently dynamic (variable number of parameters) but the expected result has a known type:

```go
ArrangeInput: &args.Dynamic[bool]{
    Params: args.Map{
        "input": "test@example.com",
        "func":  validateEmail,
    },
    Expect: true,
}
```

### Key Methods

All `args.Map` methods are delegated through `Params`:
- `Get`, `GetAsString`, `GetAsInt`, `HasDefined`, etc.
- `Invoke`, `InvokeWithValidArgs` (function invocation)

---

## args.Holder[T]

A 6-slot positional holder with a typed `WorkFunc` field and a fallback `Hashmap`.

### Structure

```go
type Holder[T any] struct {
    First    any
    Second   any
    Third    any
    Fourth   any
    Fifth    any
    Sixth    any
    WorkFunc T
    Expect   any
    Hashmap  Map
}
```

### Alias

```go
type HolderAny = Holder[any]
```

### When to Use

When you need both positional parameters AND a function reference in the same test case:

```go
ArrangeInput: &args.HolderAny{
    First:    "input-data",
    Second:   42,
    WorkFunc: myProcessFunc,
    Hashmap:  args.Map{"extra": "config"},
}
```

---

## args.LeftRight[TLeft, TRight]

Semantic two-item holder for cases where left/right directionality matters.

### Structure

```go
type LeftRight[TLeft, TRight any] struct {
    Left   TLeft
    Right  TRight
    Expect any
}
```

### Alias

```go
type LeftRightAny = LeftRight[any, any]
```

### When to Use

Comparison tests, diff tests, migration tests — anywhere "left vs right" is meaningful:

```go
ArrangeInput: &args.LeftRightAny{
    Left:   originalConfig,
    Right:  modifiedConfig,
    Expect: "3 differences",
}
```

---

## args.ThreeFunc / args.ThreeFuncAny

Variant of `Three` that includes a `WorkFunc` field for function invocation tests:

```go
type ThreeFunc[T1, T2, T3 any] struct {
    First    T1
    Second   T2
    Third    T3
    WorkFunc any
    Expect   any
}
```

### Example

```go
ArrangeInput: args.ThreeFuncAny{
    First:    "arg1",
    Second:   "arg2",
    Third:    "arg3",
    WorkFunc: myFunction,
}
```

---

## Decision Matrix: Which Args Type?

| Scenario | Use |
|----------|-----|
| Key-value input (most common) | `args.Map` |
| Positional function arguments (2-6 params) | `args.Two` – `args.Six` |
| Variable number of params + typed expect | `args.Dynamic[T]` |
| Positional params + function reference | `args.Holder[T]` |
| Left/right comparison | `args.LeftRight[T1, T2]` |
| Function params + function ref (3 args) | `args.ThreeFuncAny` |
| Expected output (always) | `args.Map` ✅ |
