# Coding Guidelines

## Receiver Types: Prefer Value Receivers (Future Direction)

> **Status**: Guideline for new code. Existing pointer receivers will be migrated incrementally.

### Rationale

- Value receivers are simpler, safer (no nil panics without guards), and communicate immutability.
- Pointer receivers should only be used when **mutating** the struct or when the struct is **large** (>~5 fields with complex types).
- All nil-safety guards (`if it == nil`) become unnecessary with value receivers.

### When to Use Pointer Receivers

- The method **modifies** the receiver (setter, initializer).
- The struct is **large** and copying would be expensive.
- The method must satisfy an interface that requires pointer receiver.
- The method implements `json.Marshaler` / `json.Unmarshaler`.

### When to Use Value Receivers

- The method is a **getter** (read-only).
- The method returns a **computed value** or **formatted string**.
- The struct is small (≤5 simple fields).
- `Json()`, `String()`, `Clone()`, `ToPtr()` — always value receivers.

### Example

```go
// ✅ Good: Value receiver for read-only methods
func (it Info) Name() string { return it.RootName }
func (it Info) Json() corejson.Result { return corejson.New(it) }
func (it Info) Clone() Info { return Info{...} }

// ✅ Good: Pointer receiver for mutation
func (it *Info) SetSecure() *Info { it.ExcludeOptions = ...; return it }
```

### Migration Plan

1. New code: Follow this guideline immediately.
2. Existing code: Migrate during refactoring passes — do NOT change receiver type in isolation (may break interface satisfaction).

---

## `interface{}` → `any` Migration

All new code must use `any` instead of `interface{}`. This is a Go 1.18+ alias and is semantically identical. See [Go Modernization Plan](/spec/01-app/11-go-modernization.md).

---

## One File Per Function

Each public function or method group gets its own `.go` file, named after the function/struct. This keeps files small and focused.

---

## Struct-as-Namespace Pattern

Group related operations on unexported struct types, exposed via package-level `var`:

```go
// unexported struct
type newCreator struct{}

// package-level var
var New newCreator

// usage: corepayload.New.PayloadWrapper.Empty()
```

---

## Zero-Nil Safety

- Return empty slices/maps instead of `nil`.
- All pointer-receiver methods must have nil guards if the receiver could be nil.
- Use `IsNull()` / `IsEmpty()` / `IsDefined()` consistently.

---

## Interface Naming

Follow Go's `-er` suffix convention:

| Pattern | Example |
|---------|---------|
| `*Getter` | `NameGetter`, `ValueGetter` |
| `*Checker` | `HasErrorChecker` |
| `*Binder` | `ContractsBinder`, `AttributesBinder` |
| `*er` | `Serializer`, `Csver`, `Stringer` |

---

## The `newCreator` Convention (Hierarchical Factory Pattern)

This is the **most important architectural pattern** in the codebase. Instead of flat `NewX()` functions, we decompose object creation into a tree of small factory structs exposed via a package-level `var New`:

```go
// vars.go
var New = newCreator{}

// newCreator.go — root aggregator
type newCreator struct {
    Widget newWidgetCreator
    Config newConfigCreator
}

// newWidgetCreator.go — one file per sub-creator
type newWidgetCreator struct{}

func (it newWidgetCreator) Empty() *Widget {
    return &Widget{Items: []string{}}
}

func (it newWidgetCreator) Create(name string) *Widget {
    return &Widget{Name: name, Items: []string{}}
}
```

**Usage**: `mypkg.New.Widget.Empty()` — IDE autocomplete guides users through the tree.

See the full guide: **[newCreator Convention](/spec/01-app/18-new-creator-convention.md)**

---

## Conditional Formatting & Readability

### Prefer Positive Conditions

Use **positive** boolean variables (`isInvalid`, `isEmpty`, `hasError`) rather than negating a variable inline (`!isValid`, `!isEmpty`). This improves readability and makes intent explicit.

```go
// ✅ Good: Positive condition via renamed variable
items, isValid := input.GetAsStrings("items")
isInvalid := !isValid

if isInvalid {
    errcore.HandleErrMessage("GetAsStrings 'items' failed")
}

// ❌ Bad: Negation inline
items, isValid := input.GetAsStrings("items")
if !isValid {
    errcore.HandleErrMessage("GetAsStrings 'items' failed")
}
```

**Exception**: When the variable is used only once and the meaning is obvious (e.g. `if !ok {`), inline negation is acceptable.

### Blank Line Before `return`

Always insert a blank line before a `return` statement when it is preceded by a line of code. This visually separates the function's exit point from its logic.

```go
// ✅ Good: Blank line before return
result := compute(input)

return result

// ✅ Good: Early return guard (no blank line needed after opening `{`)
func (it Info) Name() string {
    if it.IsEmpty() {
        return ""
    }

    return it.RootName
}

// ❌ Bad: No blank line before return
result := compute(input)
return result
```

**Exception**: Single-line function bodies do not need a blank line before `return`:

```go
func (it Info) Name() string { return it.RootName }
```

### Blank Line Rules for Control Flow Blocks

These rules apply uniformly to **all** control flow statements: `if`, `for`, `switch`, `select`, and `range`.

1. **Before the statement**: Always insert a blank line before a control flow statement when preceded by a line of code or a closing `}` (unless that `}` immediately closes an outer block).
2. **After `}`**: Insert a blank line after `}` only if the next line is **not** another `}` closing a parent block.
3. **Consecutive control flow**: When two control flow blocks appear back-to-back with no intervening code, a single blank line separates them.

```go
// ✅ Good: Spacing around if
items, isValid := input.GetAsStrings("items")
isInvalid := !isValid

if isInvalid {
    errcore.HandleErrMessage("GetAsStrings 'items' failed")
}

search, hasSearch := input.GetAsString("search")
isSearchMissing := !hasSearch

if isSearchMissing {
    errcore.HandleErrMessage("GetAsString 'search' failed")
}

// ✅ Good: Spacing around for
col := coredynamic.New.Collection.String.From(items)

for i := 0; i < col.Length(); i++ {
    process(col.SafeAt(i))
}

result := col.First()

// ✅ Good: Spacing around switch
kind := reflect.TypeOf(value).Kind()

switch kind {
case reflect.String:
    handleString(value)
case reflect.Int:
    handleInt(value)
default:
    handleOther(value)
}

// ✅ Good: Spacing around select
timeout := time.After(5 * time.Second)

select {
case msg := <-ch:
    process(msg)
case <-timeout:
    return ErrTimeout
}

// ✅ Good: Spacing around range
names := []string{"a", "b", "c"}

for _, name := range names {
    fmt.Println(name)
}

// ✅ Good: No blank line before closing parent }
for _, item := range items {
    if item == "" {
        continue
    }
}

// ❌ Bad: No breathing room
items, isValid := input.GetAsStrings("items")
isInvalid := !isValid
if isInvalid {
    errcore.HandleErrMessage("GetAsStrings 'items' failed")
}
search, hasSearch := input.GetAsString("search")
if !hasSearch {
    errcore.HandleErrMessage("GetAsString 'search' failed")
}
for i := 0; i < len(items); i++ {
    process(items[i])
}
```

---

## Function Call Argument Formatting

When a function call has **multiple arguments**, each argument must be placed on its own line — including the first argument. The closing parenthesis sits on its own line, aligned with the function call indentation.

```go
// ✅ Good: Each argument on its own line, first argument on the next line
verifyDefaultErr(
    t,
    0,
    "NilResult error is not nil",
    defaulterr.NilResult,
)

errcore.AssertDiffOnMismatch(
    t,
    caseIndex,
    tc.Title,
    actLines,
    expectedLines,
)

req := coreapi.NewTypedSimpleGenericRequest[string](
    attr,
    simpleReq,
)

// ✅ Good: Single argument can stay on the same line
Write-Success "All tests passed"
fmt.Println(value)

// ❌ Bad: Multiple arguments on the same line as function name
verifyDefaultErr(t, 0, "NilResult error is not nil", defaulterr.NilResult)

// ❌ Bad: First argument on the same line as function name
verifyDefaultErr(t,
    0,
    "NilResult error is not nil",
    defaulterr.NilResult,
)
```

**Exception**: Single-argument calls or very short two-argument calls where both fit comfortably on one line (e.g., `fmt.Sprintf("%v", value)`).

---

## Variable Naming Conventions

### Avoid Numbered Suffixes

Do **not** use numbered variable names like `val1`, `val2`, `var1`, `var2`. Use descriptive names that convey meaning.

```go
// ✅ Good: Descriptive parameter names
func VarTwo(
    isIncludeType bool,
    firstName string,
    firstValue any,
    secondName string,
    secondValue any,
) string { ... }

// ❌ Bad: Numbered suffixes
func VarTwo(
    isIncludeType bool,
    var1 string,
    val1 any,
    var2 string,
    val2 any,
) string { ... }
```

### Naming Guidelines

| Pattern | Good | Bad |
|---------|------|-----|
| Loop variables | `item`, `name`, `key` | `v`, `x`, `tmp` |
| Boolean flags | `isValid`, `hasError` | `ok2`, `flag` |
| Positional params | `firstName`, `secondValue` | `val1`, `val2` |
| Iterators | `index`, `offset` | `i2`, `j2` |

**Exception**: Single-letter variables are acceptable in very short scopes (e.g., `i` in a `for` loop, `k`/`v` in a map range).

---

## Related Docs

- [Design Philosophy](/spec/01-app/00-repo-overview.md)
- [Interface Conventions](/spec/01-app/14-core-interface-conventions.md)
- [Go Modernization Plan](/spec/01-app/11-go-modernization.md)
- [newCreator Convention](/spec/01-app/18-new-creator-convention.md)
- [Testing Guidelines](/spec/01-app/16-testing-guidelines.md)
