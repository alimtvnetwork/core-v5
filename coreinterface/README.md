# coreinterface — Shared Interface Contracts

## Overview

The `coreinterface` package defines the shared interface contracts used throughout the codebase. It provides fine-grained, composable interfaces for checkers, getters, setters, stringers, serializers, reflectors, and more. Types across all packages implement these interfaces to ensure consistent API surfaces.

## Sub-packages

| Package | Description |
|---------|-------------|
| `baseactioninf/` | Action execution interfaces (`IsExecutableChecker`) |
| `corepubsubinf/` | Pub/sub messaging interfaces |
| `entityinf/` | Entity-level interfaces (identity, lifecycle) |
| `enuminf/` | Enum-specific interfaces (`IsValidChecker`, `RangeValidateChecker`) |
| `errcoreinf/` | Error-core interfaces |
| `loggerinf/` | Logger interfaces |
| `pathextendinf/` | Path extension interfaces |
| `payloadinf/` | Payload/data transfer interfaces |
| `serializerinf/` | Serialization/deserialization contracts |

## Interface Categories

### Checkers (`all-is-checkers.go`, `all-has-checkers.go`)

Predicate interfaces for state validation:

- **Validity**: `IsValidChecker`, `IsInvalidChecker`, `IsValidInvalidChecker`
- **Nullability**: `IsNilChecker`, `IsNullChecker`, `IsNullOrEmptyChecker`, `IsAnyNullChecker`
- **State**: `IsEmptyChecker`, `IsDefinedChecker`, `IsSuccessChecker`, `IsFailedChecker`, `IsCompletedChecker`
- **Enablement**: `IsEnabledChecker`, `IsDisabledChecker`, `IsEnableAllChecker`, `IsEnableAnyChecker`
- **Containment**: `HasKeyChecker`, `HasAnyItemChecker`, `StringHasChecker`, `StringHasAllChecker`
- **Range**: `RangeValidateChecker`, `IsWithinRange*Checker`, `IsOutOfRange*Checker`
- **Dynamic**: `IsDynamicContainsChecker`, `IsDynamicItemValidChecker`, `IsDynamicNullChecker`
- **Type-specific**: `IsByteValueValidChecker`, `IsInt8ValueValidChecker`, etc.

### Getters (`all-getters.go`)

Value accessor interfaces:

- **Identity**: `NameGetter`, `IdentifierGetter`, `TypeNameGetter`, `CategoryNameGetter`
- **Values**: `ValueStringGetter`, `ValueIntegerGetter`, `ValueByteGetter`, `ValueAnyItemGetter`
- **Collections**: `StringsGetter`, `ListStringsGetter`, `SafeStringsGetter`, `AllValuesStringsGetter`
- **Errors**: `ErrorGetter`, `InvalidErrorGetter`, `ValidationErrorGetter`
- **Reflection**: `ReflectTypeGetter`, `ReflectKindGetter`, `ReflectValueGetter`
- **Maps**: `MapStringAnyGetter`, `MapStringStringGetter`, `HashmapGetter`, `KeysHashsetGetter`

### Setters (`all-setters.go`)

Mutator interfaces for setting values on types.

### Stringers (`all-stringers.go`)

String conversion interfaces:

- `AllKeysStringer`, `AllKeysSortedStringer`
- `JsonCombineStringer`, `MustJsonStringer`
- `FullStringer`, `SafeStringer`, `BuildStringer`
- `NameValueStringer`, `ToValueStringer`, `ToNumberStringer`

### Compilers & Builders

- `StringCompiler`, `BytesCompiler`, `Compiler`
- `CoreDefiner`, `FmtCompiler`
- `StringFinalizer`, `Committer`

### Serialization (`all-serializer.go`)

JSON and byte serialization contracts.

### Reflection (`all-reflection.go`)

Reflection-based type checking and conversion.

### Other

- **Binders** (`all-binders.go`): Contract binding interfaces
- **Changes** (`all-changes-related.go`): Change tracking interfaces
- **Instructions** (`all-instructions.go`): Instruction execution interfaces
- **Namers** (`all-namers.go`): Naming convention interfaces
- **Key-Value** (`all-keyval-definer.go`): Key-value pair interfaces

## Design Principle

Interfaces are intentionally **small and composable** — most define a single method. This follows the Interface Segregation Principle, allowing types to implement only the contracts they need. Composite interfaces embed smaller ones (e.g., `StringHasCombineChecker` embeds `StringHasChecker`, `StringHasAllChecker`, `StringHasAnyChecker`).

## Contributors

## Issues for Future Reference
