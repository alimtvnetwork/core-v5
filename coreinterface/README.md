# coreinterface — Shared Interface Contracts

## Overview

The `coreinterface` package defines the shared interface contracts used throughout the codebase. It provides fine-grained, composable interfaces for checkers, getters, setters, stringers, serializers, reflectors, and more. Types across all packages implement these interfaces to ensure consistent API surfaces.

## Sub-packages

| Package | Description |
|---------|-------------|
| `baseactioninf/` | Action execution interfaces (`ExecAny`, `IsExecutableChecker`) |
| `corepubsubinf/` | Pub/sub messaging interfaces (subscribers, publishers, communicators) |
| `entityinf/` | Entity-level interfaces (identity, lifecycle, task definitions) |
| `enuminf/` | Enum interfaces (`BasicEnumer`, `SimpleEnumer`, `RangeValidateChecker`) |
| `errcoreinf/` | Error-core interfaces (wrappers, collections, should-be assertions) |
| `loggerinf/` | Logger interfaces (standard, single, meta-attributes, conditional) |
| `pathextendinf/` | Path extension interfaces |
| `payloadinf/` | Payload/data transfer interfaces |
| `serializerinf/` | Serialization/deserialization contracts (JSON, bytes, field-level) |

## Interface Categories

### Checkers (`all-is-checkers.go`, `all-has-checkers.go`)

Predicate interfaces for state validation:

- **Validity**: `IsValidChecker`, `IsInvalidChecker`, `IsValidInvalidChecker`, `IsSuccessValidator`
- **Nullability**: `IsNilChecker`, `IsNullChecker`, `IsNullOrEmptyChecker`, `IsAnyNullChecker`
- **State**: `IsEmptyChecker`, `IsDefinedChecker`, `IsSuccessChecker`, `IsFailedChecker`, `IsCompletedChecker`
- **Enablement**: `IsEnabledChecker`, `IsDisabledChecker`, `IsEnableAllChecker`, `IsEnableAnyChecker`
- **Containment**: `HasKeyChecker`, `HasAnyItemChecker`, `StringHasChecker`, `StringHasAllChecker`
- **Range**: `RangeValidateChecker`, `IsWithinRange*Checker`, `IsOutOfRange*Checker`
- **Dynamic**: `IsDynamicContainsChecker`, `IsDynamicItemValidChecker`, `IsDynamicNullChecker`
- **Type-specific**: `IsByteValueValidChecker`, `IsInt8ValueValidChecker`, etc.
- **Reflection**: `IsReflectionTypeChecker`, `IsReflectKindChecker`, `IsPointerChecker`
- **Flags**: `IsEnableDisableConditionChecker`, `IsFlagsEnabledByNamesChecker`

### Getters (`all-getters.go`)

Value accessor interfaces:

- **Identity**: `NameGetter`, `IdentifierGetter`, `TypeNameGetter`, `CategoryNameGetter`
- **Values**: `ValueStringGetter`, `ValueIntegerGetter`, `ValueByteGetter`, `ValueAnyItemGetter`
- **Collections**: `StringsGetter`, `ListStringsGetter`, `SafeStringsGetter`, `AllValuesStringsGetter`
- **Errors**: `ErrorGetter`, `InvalidErrorGetter`, `ValidationErrorGetter`, `MeaningFullErrorGetter`
- **Reflection**: `ReflectTypeGetter`, `ReflectKindGetter`, `ReflectValueGetter`
- **Maps**: `MapStringAnyGetter`, `MapStringStringGetter`, `HashmapGetter`, `KeysHashsetGetter`
- **Length**: `LengthGetter`, `RawPayloadsGetter`

### Setters (`all-setters.go`)

Mutator interfaces for setting values on types.

### Stringers (`all-stringers.go`)

String conversion interfaces:

- `AllKeysStringer`, `AllKeysSortedStringer`
- `JsonCombineStringer`, `MustJsonStringer`
- `FullStringer`, `SafeStringer`, `BuildStringer`
- `NameValueStringer`, `ToValueStringer`, `ToNumberStringer`

### Compilers & Builders (`all-is-checkers.go`)

- `StringCompiler`, `BytesCompiler`, `Compiler`
- `FmtCompiler`, `StringFinalizer`, `Committer`
- `BytesCompilerIf`, `MustBytesCompiler`
- `IfStringCompiler`

### Serialization (`all-serializer.go`)

JSON and byte serialization contracts.

### Reflection (`all-reflection.go`)

Reflection-based type checking and conversion.

### Other Categories

- **Binders** (`all-binders.go`): Contract binding interfaces
- **Changes** (`all-changes-related.go`): Change tracking interfaces
- **Instructions** (`all-instructions.go`): Instruction execution interfaces
- **Namers** (`all-namers.go`): Naming convention interfaces
- **Key-Value** (`all-keyval-definer.go`): Key-value pair interfaces

### Implementation-Level Interfaces

Some interface contracts are defined alongside their implementations rather than in `coreinterface`:

| Interface | Location | Purpose |
|-----------|----------|---------|
| [`DifferChecker`](/coreimpl/enumimpl/readme.md#differchecker-interface) | `coreimpl/enumimpl/` | Strategy interface for `DynamicMap` diff comparison — controls how value differences and missing keys are reported |

## Standalone Interface Files

Single-purpose interfaces with dedicated files:

| File | Interface | Purpose |
|------|-----------|---------|
| `Cloner.go` | `Cloner[T]` | Generic type-safe cloning |
| `Disposer.go` | `Disposer` | Resource disposal |
| `CoreDefiner.go` | `CoreDefiner` | Core definition contract |
| `StringCompiler.go` | `StringCompiler` | String compilation |
| `TableNamer.go` | `TableNamer` | Database table name provider |
| `CsvLiner.go` | `CsvLiner` | CSV line generation |
| `SlicePager.go` | `SlicePager` | Slice pagination |
| `DynamicAdder.go` | `DynamicAdder` | Dynamic item addition |
| `DynamicLinq.go` | `DynamicLinq` | LINQ-style dynamic queries |
| `ReflectSetter.go` | `ReflectSetter` | Reflection-based value setting |
| `ErrorHandler.go` | `ErrorHandler` | Error handling contract |

## Design Principle

Interfaces are intentionally **small and composable** — most define a single method. This follows the Interface Segregation Principle, allowing types to implement only the contracts they need. Composite interfaces embed smaller ones (e.g., `StringHasCombineChecker` embeds `StringHasChecker`, `StringHasAllChecker`, `StringHasAnyChecker`).

```go
// Composition example
type IsSuccessValidator interface {
    IsValidChecker    // IsValid() bool
    IsSuccessChecker  // IsSuccess() bool
    IsFailedChecker   // IsFailed() bool
}
```

## How to Extend Safely

- **New checker**: Add to `all-is-checkers.go` or `all-has-checkers.go` based on prefix convention.
- **New getter**: Add to `all-getters.go`.
- **Complex interface**: Create a dedicated file (e.g., `MyContract.go`).
- **New sub-package**: Create a directory under `coreinterface/` for domain-specific contracts (e.g., `loggerinf/`).
- **Generic interface**: Use type parameters where appropriate (see `Cloner[T]`).

## Related Docs

- [Repo Overview](../spec/01-app/00-repo-overview.md)
- [coregeneric spec](../spec/01-app/folders/05a-coregeneric.md)
