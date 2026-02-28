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

## Related Docs

- [Design Philosophy](/spec/01-app/00-repo-overview.md)
- [Interface Conventions](/spec/01-app/14-core-interface-conventions.md)
- [Go Modernization Plan](/spec/01-app/11-go-modernization.md)
- [newCreator Convention](/spec/01-app/18-new-creator-convention.md)
- [Testing Guidelines](/spec/01-app/16-testing-guidelines.md)
