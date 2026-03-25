# coreimpl — Core Implementations

Package `coreimpl` provides concrete implementations of core interface contracts. Currently contains `enumimpl`, the enum implementation engine supporting byte, int8, int16, int32, uint16, and string enum types with dynamic maps, diff checking, and JSON serialization.

## Architecture

```
coreimpl/
└── enumimpl/                              # Enum implementation engine (see enumimpl/readme.md)
    ├── vars.go                            # Package singletons: New, DefaultDiffCheckerImpl, LeftRightDiffCheckerImpl
    ├── newCreator.go                      # New.* — factory for all enum types
    ├── newBasicByteCreator.go             # New.BasicByte.* — byte enum factory
    ├── newBasicInt8Creator.go             # New.BasicInt8.* — int8 enum factory
    ├── newBasicInt16Creator.go            # New.BasicInt16.* — int16 enum factory
    ├── newBasicInt32Creator.go            # New.BasicInt32.* — int32 enum factory
    ├── newBasicUInt16Creator.go           # New.BasicUInt16.* — uint16 enum factory
    ├── newBasicStringCreator.go           # New.BasicString.* — string enum factory
    ├── BasicByte.go                       # BasicByte enum struct
    ├── BasicByter.go                      # BasicByter interface
    ├── BasicInt8.go                       # BasicInt8 enum struct
    ├── BasicInt16.go                      # BasicInt16 enum struct
    ├── BasicInt32.go                      # BasicInt32 enum struct
    ├── BasicUInt16.go                     # BasicUInt16 enum struct
    ├── BasicString.go                     # BasicString enum struct
    ├── numberEnumBase.go                  # Shared base for numeric enum types
    ├── DynamicMap.go                      # DynamicMap — string→any map for enum metadata
    ├── DiffLeftRight.go                   # Left/right diff comparison for dynamic maps
    ├── differCheckerImpl.go               # Default diff checker implementation
    ├── leftRightDiffCheckerImpl.go        # Left/right diff checker implementation
    ├── KeyAnyVal.go                       # Key-value pair with any value
    ├── KeyAnyValues.go                    # Slice of KeyAnyVal
    ├── KeyValInteger.go                   # Key-value pair with integer value
    ├── AllNameValues.go                   # All enum names and values
    ├── NameWithValue.go                   # Name-value string formatter
    ├── Format.go / FormatUsingFmt.go      # Enum formatting utilities
    ├── OnlySupportedErr.go               # "Only supported" error generator
    ├── PrependJoin.go                     # Prepend-join string builder
    ├── JoinPrependUsingDot.go            # Dot-separated join
    ├── UnsupportedNames.go               # Unsupported name collector
    ├── all-interfaces.go                  # Enum interface contracts
    ├── consts.go                          # Package constants
    ├── funcs.go                           # Standalone helper functions
    ├── enumtype/                          # Enum type metadata
    │   └── Variant.go                     # Variant type with min/max/ranges
    ├── ConvAnyValToInteger.go             # Any → integer conversion
    ├── convAnyValToString.go              # Any → string conversion
    ├── IntegersRangesOfAnyVal.go          # Integer ranges from any values
    ├── stringsToHashset.go               # Strings → Hashset conversion
    ├── toHashset.go                       # Enum → Hashset conversion
    ├── toJsonName.go                      # Enum → JSON name
    ├── toNamer.go                         # Enum → Namer interface
    ├── toStringPrintableDynamicMap.go     # DynamicMap → printable string
    ├── toStringsSliceOfDiffMap.go         # Diff map → string slice
    └── enumUnmarshallingMappingFailedError.go # Unmarshalling error type
```

## Usage

### Creating an Enum Type

```go
import "github.com/alimtvnetwork/core/coreimpl/enumimpl"

// Byte enum
byteEnum := enumimpl.New.BasicByte.CreateUsingSlicePlusAliasMapOptions(
    firstItem,
    nameToValueMap,
    aliasMap,
)

// String enum
stringEnum := enumimpl.New.BasicString.CreateUsingSlicePlusAliasMapOptions(
    firstItem,
    nameToValueMap,
    aliasMap,
)
```

### DifferChecker — Map Diff Strategy

[`DifferChecker`](enumimpl/readme.md#differchecker-interface) is the strategy interface used by `DynamicMap` to control how value differences and missing keys are reported during map comparison.

```go
type DifferChecker interface {
    GetSingleDiffResult(isLeft bool, l, r any) any
    GetResultOnKeyMissingInRightExistInLeft(lKey string, lVal any) any
    IsEqual(isRegardless bool, l, r any) bool
}
```

Two built-in implementations are provided as package singletons:

| Singleton | Behavior |
|-----------|----------|
| `DefaultDiffCheckerImpl` | Returns raw differing values; missing keys return left value as-is |
| `LeftRightDiffCheckerImpl` | Returns `DiffLeftRight` JSON (e.g., `{"Left":5,"Right":6}`); missing keys include type annotation |

```go
// Compare two dynamic maps using left/right labeled diffs
left := enumimpl.DynamicMap{"a": 1, "b": 3}
right := map[string]any{"a": 1, "b": 4}

diffMap := left.DiffRawUsingDifferChecker(
    enumimpl.LeftRightDiffCheckerImpl,
    true,
    right,
)
```

See [enumimpl/readme.md](enumimpl/readme.md) for full interface docs, custom implementation examples, and DynamicMap integration details.

## Related Docs

- [enumimpl README](enumimpl/readme.md)
- [Folder Spec](/spec/01-app/folders/10-remaining-packages.md)
- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
- [New Creator Pattern](/spec/01-app/21-new-creator-pattern.md)
