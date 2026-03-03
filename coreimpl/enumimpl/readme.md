# enumimpl — Enum Implementation Engine

Package `enumimpl` provides concrete implementations of enum types and diff-checking utilities for dynamic maps. Supports byte, int8, int16, int32, uint16, and string enum types with dynamic maps, diff checking, and JSON serialization.

## Architecture

```
enumimpl/
├── vars.go                            # Package singletons: New, DefaultDiffCheckerImpl, LeftRightDiffCheckerImpl
├── all-interfaces.go                  # Internal + exported interfaces (DifferChecker)
├── funcs.go                           # Function type aliases
├── consts.go                          # Package constants
│
├── DifferChecker (Diff Subsystem)
│   ├── all-interfaces.go              # DifferChecker interface definition
│   ├── differCheckerImpl.go           # DefaultDiffCheckerImpl — default diff strategy
│   ├── leftRightDiffCheckerImpl.go    # LeftRightDiffCheckerImpl — left/right labeled diffs
│   ├── DiffLeftRight.go               # DiffLeftRight struct — left/right value pair with comparison
│   └── DynamicMap.go                  # DynamicMap.DiffRaw*, ShouldDiffMessage*, DiffRawLeftRight*
│
├── Enum Types
│   ├── newCreator.go                  # New.* — factory for all enum types
│   ├── newBasicByteCreator.go         # New.BasicByte.* — byte enum factory
│   ├── newBasicInt8Creator.go         # New.BasicInt8.* — int8 enum factory
│   ├── newBasicInt16Creator.go        # New.BasicInt16.* — int16 enum factory
│   ├── newBasicInt32Creator.go        # New.BasicInt32.* — int32 enum factory
│   ├── newBasicUInt16Creator.go       # New.BasicUInt16.* — uint16 enum factory
│   ├── newBasicStringCreator.go       # New.BasicString.* — string enum factory
│   ├── BasicByte.go                   # BasicByte enum struct
│   ├── BasicByter.go                  # BasicByter interface
│   ├── BasicInt8.go                   # BasicInt8 enum struct
│   ├── BasicInt16.go                  # BasicInt16 enum struct
│   ├── BasicInt32.go                  # BasicInt32 enum struct
│   ├── BasicUInt16.go                 # BasicUInt16 enum struct
│   ├── BasicString.go                 # BasicString enum struct
│   └── numberEnumBase.go              # Shared base for numeric enum types
│
├── DynamicMap & Utilities
│   ├── DynamicMap.go                  # DynamicMap — string→any map with diff, equality, iteration
│   ├── KeyAnyVal.go                   # Key-value pair with any value
│   ├── KeyAnyValues.go                # Slice of KeyAnyVal
│   ├── KeyValInteger.go               # Key-value pair with integer value
│   ├── AllNameValues.go               # All enum names and values
│   ├── NameWithValue.go               # Name-value string formatter
│   ├── Format.go / FormatUsingFmt.go  # Enum formatting utilities
│   ├── OnlySupportedErr.go            # "Only supported" error generator
│   ├── PrependJoin.go                 # Prepend-join string builder
│   ├── JoinPrependUsingDot.go         # Dot-separated join
│   └── UnsupportedNames.go            # Unsupported name collector
│
├── Converters
│   ├── ConvAnyValToInteger.go         # Any → integer conversion
│   ├── convAnyValToString.go          # Any → string conversion
│   ├── IntegersRangesOfAnyVal.go      # Integer ranges from any values
│   ├── stringsToHashset.go            # Strings → Hashset conversion
│   ├── toHashset.go                   # Enum → Hashset conversion
│   ├── toJsonName.go                  # Enum → JSON name
│   ├── toNamer.go                     # Enum → Namer interface
│   ├── toStringPrintableDynamicMap.go # DynamicMap → printable string
│   └── toStringsSliceOfDiffMap.go     # Diff map → string slice
│
├── Error Types
│   └── enumUnmarshallingMappingFailedError.go  # Unmarshalling error type
│
└── enumtype/                          # Enum type metadata
    └── Variant.go                     # Variant type with min/max/ranges
```

## DifferChecker Interface

`DifferChecker` is the strategy interface for comparing two dynamic maps. It controls how individual value differences are reported and how missing keys are handled.

```go
type DifferChecker interface {
    GetSingleDiffResult(isLeft bool, l, r any) any
    GetResultOnKeyMissingInRightExistInLeft(lKey string, lVal any) any
    IsEqual(isRegardless bool, l, r any) bool
}
```

### Methods

| Method | Purpose |
|--------|---------|
| `GetSingleDiffResult(isLeft, l, r)` | Returns the diff result for a single key where both maps have the key but values differ. `isLeft` indicates which side's perspective is being reported. |
| `GetResultOnKeyMissingInRightExistInLeft(lKey, lVal)` | Returns the diff result when a key exists in the left map but is missing from the right map. |
| `IsEqual(isRegardless, l, r)` | Compares two values for equality. When `isRegardless` is true, compares by string representation (type-agnostic). |

### Built-in Implementations

| Singleton | Type | Behavior |
|-----------|------|----------|
| `DefaultDiffCheckerImpl` | `*differCheckerImpl` | Returns raw left or right value on diff. Missing keys return the left value as-is. |
| `LeftRightDiffCheckerImpl` | `*leftRightDiffCheckerImpl` | Returns `DiffLeftRight` JSON string (e.g., `{"Left":5,"Right":6}`). Missing keys return `"5 (type:int) - left - key is missing!"`. |

### Usage with DynamicMap

```go
left := enumimpl.DynamicMap{"a": 1, "b": 3, "c": 5}
right := map[string]any{"a": 1, "b": 4}

// Default diff — returns raw differing values
diffMap := left.DiffRaw(true, right)

// Left/right labeled diff — returns DiffLeftRight JSON strings
diffMap = left.DiffRawUsingDifferChecker(
    enumimpl.LeftRightDiffCheckerImpl,
    true,
    right,
)

// Full diff message with title
msg := left.ShouldDiffLeftRightMessageUsingDifferChecker(
    enumimpl.LeftRightDiffCheckerImpl,
    true,
    "My Diff Title",
    right,
)
```

### Implementing a Custom DifferChecker

```go
type myChecker struct{}

func (c *myChecker) GetSingleDiffResult(isLeft bool, l, r any) any {
    return fmt.Sprintf("changed: %v → %v", l, r)
}

func (c *myChecker) GetResultOnKeyMissingInRightExistInLeft(lKey string, lVal any) any {
    return fmt.Sprintf("removed: %s=%v", lKey, lVal)
}

func (c *myChecker) IsEqual(isRegardless bool, l, r any) bool {
    return reflect.DeepEqual(l, r)
}
```

## Related Docs

- [coreimpl README](../README.md)
- [Folder Spec](/spec/01-app/folders/10-remaining-packages.md)
- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
- [New Creator Pattern](/spec/01-app/21-new-creator-pattern.md)
