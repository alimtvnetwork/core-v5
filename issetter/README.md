# issetter — Multi-Valued Boolean Enum

Package `issetter` provides a 6-valued boolean enum type (`Value`) backed by `byte`, supporting states beyond simple true/false: `Uninitialized`, `True`, `False`, `Unset`, `Set`, and `Wildcard`. Implements the full `enuminf` interface contracts including JSON marshalling, range validation, and name-based lookup.

## Architecture

```
issetter/
├── Value.go                     # Value type (byte enum) — 800+ lines of methods
├── New.go                       # New(name) (Value, error) — parse from string
├── NewBool.go                   # NewBool(v, t, f) Value — ternary constructor
├── NewBooleans.go               # NewBooleans(conditions...) Value — combined
├── NewMust.go                   # NewMust(name) Value — panics on error
├── GetBool.go                   # GetBool(bool) Value — bool → True/False
├── GetSet.go                    # GetSet(bool) Value — bool → Set/Unset
├── GetSetByte.go                # GetSetByte(bool) byte
├── GetSetUnset.go               # GetSetUnset(bool) Value
├── GetSetterByComparing.go      # GetSetterByComparing(a, b) Value
├── CombinedBooleans.go          # CombinedBooleans(...bool) Value
├── IsCompareResult.go           # IsCompareResult type & methods
├── IsOutOfRange.go              # IsOutOfRange(byte) bool
├── IntegerEnumRanges.go         # IntegerEnumRanges() []int
├── Min.go / Max.go              # Min/Max Value constants
├── MinByte.go / MaxByte.go      # Min/Max as byte
├── RangesNamesCsv.go            # RangeNamesCsv() string
├── consts.go                    # Yes/No string constants
├── vars.go                      # Lookup maps, name arrays, JSON maps
├── generateDynamicRangesMap.go  # Dynamic ranges map builder
├── jsonBytes.go                 # JSON byte helper
└── toHashset.go                 # Internal hashset helper
```

## Value Constants

| Constant | Byte | Logical Group | Description |
|----------|------|---------------|-------------|
| `Uninitialized` | `0` | Undefined | Default zero value, not yet set |
| `True` | `1` | Positive | Boolean true / accepted |
| `False` | `2` | Negative | Boolean false / rejected |
| `Unset` | `3` | Negative | Explicitly unset |
| `Set` | `4` | Positive | Explicitly set |
| `Wildcard` | `5` | Undefined | Matches anything |

## Logical Groupings

| Group | Values | Methods |
|-------|--------|---------|
| **Positive** | `True`, `Set` | `IsOn()`, `IsAccept()`, `IsSuccess()`, `IsOnLogically()` |
| **Negative** | `False`, `Unset` | `IsOff()`, `IsReject()`, `IsFailed()`, `IsOffLogically()` |
| **Undefined** | `Uninitialized`, `Wildcard` | `IsAsk()`, `IsSkip()`, `IsIndeterminate()`, `IsUndefinedLogically()` |

## Key Methods

### State Checks

| Method | Description |
|--------|-------------|
| `IsTrue()` / `IsFalse()` | Exact match |
| `IsSet()` / `IsUnset()` | Exact match |
| `IsWildcard()` | Exact match |
| `HasInitialized()` | Not `Uninitialized` |
| `IsInitBoolean()` | `True` or `False` |
| `IsDefinedBoolean()` | `True` or `False` |
| `IsDefinedLogically()` | Not `Uninitialized` or `Wildcard` |
| `Boolean()` | Returns Go `bool` (`True` → `true`) |

### Conversion

| Method | Description |
|--------|-------------|
| `ToBooleanValue()` | Converts `Set`→`True`, `Unset`→`False` |
| `ToSetUnsetValue()` | Converts `True`→`Set`, `False`→`Unset` |
| `Value()` / `ValueByte()` | Raw byte value |
| `ValueInt()` / `ValueInt8()` | Integer casts |
| `Name()` / `String()` | Human-readable name |
| `NameValue()` | `"Name = Value"` format |

### Lazy Evaluation

| Method | Description |
|--------|-------------|
| `LazyEvaluateBool(func())` | Runs func only if uninitialized, then sets `True` |
| `LazyEvaluateSet(func())` | Runs func only if uninitialized, then sets `Set` |
| `GetSetBoolOnInvalid(bool)` | Returns cached bool or sets from value |
| `GetSetBoolOnInvalidFunc(func() bool)` | Same with lazy function |

### Wildcard & Conditional

| Method | Description |
|--------|-------------|
| `WildcardApply(bool)` | Returns input if wildcard/unset, else `IsTrue()` |
| `IsWildcardOrBool(bool)` | Always true on wildcard, else returns the bool |
| `ToByteCondition(t, f, inv)` | Maps True/False to custom byte values |

### Enum Interface

| Method | Description |
|--------|-------------|
| `EnumType()` | Returns `enumtype.Byte` |
| `IntegerEnumRanges()` | Valid integer range |
| `RangeNamesCsv()` | CSV of all value names |
| `AllNameValues()` | Slice of all name-value pairs |
| `IsNameEqual(string)` | Compare by name |
| `MarshalJSON()` / `UnmarshalJSON()` | JSON support |
| `Format(string)` | Template formatting with `{type-name}`, `{name}`, `{value}` |

## Display Formats

| Method | True | False | Wildcard | Uninitialized |
|--------|------|-------|----------|---------------|
| `YesNoName()` | Yes | No | * | - |
| `OnOffName()` | On | Off | * | - |
| `TrueFalseName()` | True | False | * | - |
| `SetUnsetName()` | set | unset | * | - |

## Factory Functions

```go
// Parse from string (supports "true", "yes", "1", "set", etc.)
v, err := issetter.New("yes")       // True, nil
v = issetter.NewMust("set")         // Set (panics on error)

// From boolean
v = issetter.GetBool(true)          // True
v = issetter.GetSet(true)           // Set

// Combined — any false → False
v = issetter.CombinedBooleans(true, true, false) // False

// Ternary
v = issetter.NewBool(cond, issetter.Set, issetter.Unset)
```

## JSON Mapping

Supports both quoted and unquoted values:

```
"true", "True", "yes", "Yes", "1" → True
"false", "False", "no", "No"      → False
"set", "Set"                       → Set
"unset", "Unset"                   → Unset
"*", "%", "Wildcard"               → Wildcard
"", "0", "-", "-1"                 → Uninitialized
```

## Related Docs

- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
- [Enum Interface Contracts](/coreinterface/enuminf/README.md)
- [Folder Spec](/spec/01-app/folders/10-remaining-packages.md)
