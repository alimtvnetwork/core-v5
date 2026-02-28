# Core Interface Conventions

## Overview

The `coreinterface/` package is the single source of truth for all interface contracts in the auk-go ecosystem. It follows strict Go conventions and naming patterns that must be replicated in every downstream repository.

## The `-er` Suffix Convention

Go convention: **interfaces are named with an `-er` suffix** when they represent a capability or behavior.

### Standard Patterns

| Pattern | Naming Rule | Examples |
|---------|-------------|---------|
| **Getter** | `{Field}Getter` | `NameGetter`, `ValueStringGetter`, `LengthGetter`, `ErrorGetter` |
| **Checker** | `Is{Condition}Checker` | `IsValidChecker`, `IsEmptyChecker`, `IsNullChecker`, `IsDefinedChecker` |
| **Has Checker** | `Has{Thing}Checker` | `HasAnyItemChecker`, `HasAll`, `HasAny` |
| **Action** | `{Verb}er` | `Executor`, `Initializer`, `Planner`, `Serializer`, `Disposer` |
| **Binder** | `{Type}ContractsBinder` | `BasicSlicerContractsBinder`, `JsonBytesStringerContractsBinder` |
| **Namer** | `{Scope}Namer` | `EntityTypeNamer`, `CategoryNamer`, `TableNamer`, `ToNamer` |
| **Stringer** | `{Scope}Stringer` | `FullStringer`, `SafeStringer`, `BuildStringer`, `JsonCombineStringer` |

### What `-er` Stands For

The `-er` suffix in Go interfaces denotes **"something that does X"**:
- `Reader` = something that reads
- `Writer` = something that writes
- `Stringer` = something that produces a string
- `Getter` = something that gets a value
- `Checker` = something that checks a condition

This is a Go community convention (from the standard library: `io.Reader`, `io.Writer`, `fmt.Stringer`).

## Interface Composition

Larger interfaces are **composed by embedding** smaller ones:

```go
// Small, focused interfaces
type IsValidChecker interface { IsValid() bool }
type IsInvalidChecker interface { IsInvalid() bool }

// Composed interface
type IsValidInvalidChecker interface {
    IsValidChecker
    IsInvalidChecker
}

// Further composed
type SimpleValidInvalidChecker interface {
    IsValidChecker
    IsInvalidChecker
    InvalidMessageGetter
}
```

## Method Naming Conventions

| Convention | Example | Rule |
|-----------|---------|------|
| Getter methods | `Name() string` | No `Get` prefix for single-value getters |
| Boolean methods | `IsValid() bool`, `HasAny() bool` | `Is` or `Has` prefix |
| Error methods | `InvalidError() error` | Returns error with context |
| Pointer returns | `SafeBytesPtr() *[]byte` | `Ptr` suffix for pointer returns |
| String conversion | `String() string` | Implements `fmt.Stringer` |
| JSON output | `JsonStringMust() string` | `Must` suffix = panics on error |
| Slice returns | `Strings() []string`, `Lines() []string` | Plural noun |

## Error Handling Patterns

1. **Error getters return `error`**: `InvalidError() error`, `ValidationError() error`
2. **Must variants panic**: `SerializeMust() []byte`, `JsonStringMust() string`
3. **Combined result**: Functions return `(value, error)` tuples

## File Organization Convention

| File Name | Contents |
|-----------|----------|
| `all-getters.go` | All `*Getter` interfaces |
| `all-is-checkers.go` | All `Is*Checker` and `Has*Checker` interfaces |
| `all-binders.go` | All `*ContractsBinder` interfaces |
| `all-stringers.go` | All string-output interfaces |
| `all-serializer.go` | Serializer interfaces |
| `all-namers.go` | Naming interfaces |
| `{TypeName}.go` | Complex single-purpose interfaces |

## Where This Convention Applies

- **Every package** in the auk-go ecosystem must follow these naming rules.
- **Every new interface** must be added to `coreinterface/` if it's reusable.
- **Downstream repos** should import from `coreinterface/` rather than redefining.

## CSV Interface Example

```go
// Csver — follows -er convention for CSV generation capability
type Csver interface {
    Csv() string
    CsvOptions(isSkipQuoteOnlyOnExistence bool) string
}

// CsvLiner — follows -er convention for CSV line generation
type CsvLiner interface {
    CsvLines() *[]string
    CsvLinesOptions(isSkipQuoteOnlyOnExistence bool) *[]string
}
```

## Binder Pattern

Binders enforce that a type both implements an interface AND provides a self-returning method for type-safe binding:

```go
type BasicSlicerContractsBinder interface {
    BasicSlicer                                        // must implement BasicSlicer
    AsBasicSliceContractsBinder() BasicSlicerContractsBinder  // self-return for binding
}
```

This pattern ensures compile-time contract verification.

## Related Docs

- [coreinterface folder spec](./folders/04-coreinterface.md)
- [Repo Overview](./00-repo-overview.md)
