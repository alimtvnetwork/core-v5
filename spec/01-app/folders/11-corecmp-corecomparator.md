# corecmp & corecomparator

## Folder Purpose

Comparison infrastructure for the auk-go ecosystem. `corecomparator` defines the `Compare` result type; `corecmp` provides concrete comparison functions for all built-in types.

## Package Breakdown

### `corecomparator/`

Defines the `Compare` enum (`byte`-backed) with 7 values: `Equal`, `LeftGreater`, `LeftGreaterEqual`, `LeftLess`, `LeftLessEqual`, `NotEqual`, `Inconclusive`. Includes JSON marshaling, operator symbols, logical grouping methods, and embeddable case-sensitivity structs.

### `corecmp/`

Concrete comparison functions returning `corecomparator.Compare`:

| Category | Functions |
|----------|-----------|
| **Primitives** | `Integer`, `Integer8/16/32/64`, `Byte`, `Time` |
| **Pointers** | `IntegerPtr`, `Integer8/16/32/64Ptr`, `BytePtr`, `TimePtr` |
| **Any** | `AnyItem` — interface equality via `==`, returns `Inconclusive` for non-comparable types |
| **Equality** | `IsIntegersEqual(Ptr)`, `IsStringsEqual(Ptr)`, `IsStringsEqualWithoutOrder` |
| **Versioning** | `VersionSliceByte`, `VersionSliceInteger` |

## Responsibilities

1. Provide a single, consistent comparison result type (`Compare`) used across all packages.
2. Offer nil-safe pointer comparisons for all numeric types.
3. Support any-type comparison with explicit `Inconclusive` for non-comparable values.
4. Enable version comparison via slice-based byte/integer comparators.
5. Provide JSON serialization for comparison results.

## Key Design Decisions

- **`Inconclusive` over panic**: `AnyItem` returns `Inconclusive` instead of panicking when `==` cannot determine equality for complex types.
- **Pointer nil-safety**: All `*Ptr` variants handle nil inputs gracefully — both nil = `Equal`, one nil = `NotEqual`.
- **Logical grouping methods**: `IsLeftGreaterEqualLogically()` returns true for `Equal`, `LeftGreater`, and `LeftGreaterEqual`, enabling range-style checks.

## Key Files

### `corecomparator/`

| File | Purpose |
|------|---------|
| `Compare.go` | `Compare` type definition, all methods, JSON support |
| `consts.go` | Package-level constants |
| `Ranges.go` | Full range of Compare values |
| `RangeNamesCsv.go` | CSV name representation |
| `BaseIsCaseSensitive.go` | Embeddable case-sensitive flag |
| `BaseIsIgnoreCase.go` | Embeddable case-insensitive flag |
| `Min.go` / `Max.go` | Min/Max helpers |
| `MinLength.go` | Minimum length constraint |

### `corecmp/`

| File | Purpose |
|------|---------|
| `AnyItem.go` | `any` type comparison via `==` |
| `Integer.go` | `int` comparison |
| `Integer8.go` – `Integer64.go` | Sized integer comparisons |
| `IntegerPtr.go` – `Integer64Ptr.go` | Nil-safe pointer comparisons |
| `Byte.go` / `BytePtr.go` | Byte comparisons |
| `Time.go` / `TimePtr.go` | Time comparisons |
| `IsIntegersEqual.go` / `IsIntegersEqualPtr.go` | Integer equality shortcuts |
| `IsStringsEqual.go` / `IsStringsEqualPtr.go` | String equality shortcuts |
| `IsStringsEqualWithoutOrder.go` | Order-independent string slice equality |
| `VersionSliceByte.go` / `VersionSliceInteger.go` | Version comparisons |

## Related Docs

- [corecmp Readme](../../corecmp/Readme.md)
- [corecomparator Readme](../../corecomparator/Readme.md)
- [Remaining Packages spec](./10-remaining-packages.md)
- [Folder Map](../01-folder-map.md)
