# anycmp — Any-Type Quick Comparison

## Overview

Package `anycmp` provides a lightweight, quick-determination comparator for `any`-typed values. It returns `Equal`, `NotEqual`, or `Inconclusive` without using deep reflection, bytes comparison, or string comparison.

## Function

### `Cmp(left, right any) corecomparator.Compare`

Quick equality determination for two `any` values using pointer identity and nil checks only.

#### Comparison Steps

| Step | Condition | Result |
|------|-----------|--------|
| 1 | `left == right` (same pointer or identical value) | `Equal` |
| 2 | Both are `nil` | `Equal` |
| 3 | One is `nil`, other is not | `NotEqual` |
| 4 | Both are reflection-null (via `isany.NullLeftRight`) | `Equal` |
| 5 | One is reflection-null, other is not | `NotEqual` |
| 6 | Otherwise | `Inconclusive` |

#### Key Design Decisions

- **No deep comparison**: Intentionally avoids `reflect.DeepEqual`, byte-level, or string-level comparison for performance.
- **`Inconclusive` over false negatives**: Returns `Inconclusive` when equality cannot be definitively determined, allowing callers to decide whether to perform deeper comparison.
- **Reflection-null awareness**: Uses `isany.NullLeftRight` to detect typed-nil interfaces (e.g., `(*MyStruct)(nil)` stored as `any`), which the standard `==` operator would miss.

## Dependencies

| Package | Usage |
|---------|-------|
| `corecomparator` | `Compare` result type (`Equal`, `NotEqual`, `Inconclusive`) |
| `isany` | `NullLeftRight` for reflection-based nil detection |

## Usage

```go
import "gitlab.com/auk-go/core/anycmp"

result := anycmp.Cmp(a, b)

if result.IsEqual() {
    // definitely equal (same pointer or both nil)
}

if result.IsInconclusive() {
    // need deeper comparison to determine equality
}
```

## Architecture

```
anycmp/
├── Cmp.go        # Single Cmp function — the package's only export
└── readme.md
```

## File Organization

| File | Purpose |
|------|---------|
| `Cmp.go` | Single `Cmp` function — the package's only export |

## How to Extend Safely

- **New comparison depth**: Add a separate function (e.g., `DeepCmp`) rather than modifying `Cmp` — callers rely on its lightweight, no-reflection guarantee.
- **New return states**: Coordinate with `corecomparator.Compare` enum — do not add package-local result types.
- **Type-specific fast paths**: Add as separate functions (e.g., `CmpString`) that short-circuit before falling back to `Cmp`.

## Related Docs

- [corecomparator readme](../corecomparator/readme.md)
- [corecmp readme](../corecmp/readme.md)
- [Comparison & Sorting spec](../spec/01-app/folders/10-remaining-packages.md)
