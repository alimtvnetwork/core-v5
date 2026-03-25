# coreappend

String and slice assembly helpers that prepend, append, and join arbitrary values with nil-skipping and custom formatting.

## Architecture

```
coreappend/
├── AppendAnyItemsToStringSkipOnNil.go          # Join items with an append suffix, skip nil
├── PrependAnyItemsToStringSkipOnNil.go         # Join items with a prepend prefix, skip nil
├── PrependAppendAnyItemsToStringSkipOnNil.go   # Join items with both prepend and append, skip nil
├── PrependAppendAnyItemsToStringsSkipOnNil.go  # Core: build []string from any items, skip nil
├── PrependAppendAnyItemsToStringsUsingFunc.go  # Build []string via custom compiler func
└── MapStringStringToMapStringToAnyItems.go     # Merge map[string]any into map[string]string
```

## Features

### String Assembly (nil-skipping)

| Function | Description |
|---|---|
| `AppendAnyItemsToStringSkipOnNil` | Join variadic items with an append item at the end, skip nil values |
| `PrependAnyItemsToStringSkipOnNil` | Join variadic items with a prepend item at the start, skip nil values |
| `PrependAppendAnyItemsToStringSkipOnNil` | Join variadic items with both prepend and append items, skip nil values |
| `PrependAppendAnyItemsToStringsSkipOnNil` | Core builder — returns `[]string` from any items, skipping nil |

### Custom Formatting

| Function | Description |
|---|---|
| `PrependAppendAnyItemsToStringsUsingFunc` | Build `[]string` using a custom `func(any) string` compiler, with optional empty-string skipping |

### Map Merging

| Function | Description |
|---|---|
| `MapStringStringAppendMapStringToAnyItems` | Merge `map[string]any` values (via `Sprintf`) into an existing `map[string]string`, optionally skipping empty |

## Usage Examples

```go
import "github.com/alimtvnetwork/core/coreappend"

// Append item at end, skip nil
result := coreappend.AppendAnyItemsToStringSkipOnNil(
    ", ",       // joiner
    "suffix",   // appendItem
    "a", nil, "b",
)
// result: "a, b, suffix"

// Prepend + append with nil-skipping
result := coreappend.PrependAppendAnyItemsToStringSkipOnNil(
    " | ",
    "START", "END",
    "x", nil, "y",
)
// result: "START | x | y | END"

// Custom compiler function
slice := coreappend.PrependAppendAnyItemsToStringsUsingFunc(
    true,  // skip empty strings
    func(item any) string { return fmt.Sprintf("[%v]", item) },
    "pre", "post",
    "a", "b",
)
// slice: ["[pre]", "[a]", "[b]", "[post]"]
```

## Related Docs

- [constants](../constants/readme.md) — `SprintValueFormat` used by formatters
