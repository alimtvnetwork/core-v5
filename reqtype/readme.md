# reqtype

Byte-based enum representing **request types** for CRUD operations, HTTP methods, file mutations, service lifecycle actions, and inheritance strategies. Built on `enumimpl.BasicByte` with full JSON marshalling.

## Architecture

```
reqtype/
├── Request.go                    # Request enum type + all methods (~850 lines)
├── vars.go                       # Ranges array, RangesMap, logical grouping maps, BasicEnumImpl
├── consts.go                     # Package constants
├── ResultStatus.go               # Result struct for range validation
├── Min.go / Max.go               # Min/Max enum boundaries
├── start.go / end.go             # Range boundary helpers
├── RangesInBetween.go            # Generate slice of Request values in a range
├── RangesStrings.go              # Convert requests to string names
├── RangesString.go               # Join request names with joiner
├── RangesStringDefaultJoiner.go  # Join with default joiner
├── RangesNotMeet.go              # Format "range not met" message
├── RangesNotMeetError.go         # Return error for range not met
├── RangesNotSupportedFor.go      # Return error for unsupported requests
├── RangesOnlySupportedFor.go     # Return error for only-supported requests
└── RangesInvalidErr.go           # Return invalid-range error
```

## Enum Values

### CRUD Operations

| Value | Description |
|---|---|
| `Create`, `Read`, `Update`, `Delete`, `Drop` | Basic CRUD |
| `CreateOrUpdate`, `CreateOrSkipOnExist` | Conditional create |
| `UpdateOrSkipOnNonExist`, `UpdateOnExist` | Conditional update |
| `DeleteOrSkipOnNonExist`, `DropOrSkipOnNonExist`, `DropOnExist`, `DropCreate` | Conditional delete/drop |
| `ExistCheck`, `SkipOnExist` | Existence checks |

### Content Mutation

| Value | Description |
|---|---|
| `Append`, `AppendLines`, `AppendByCompare`, `AppendLinesByCompare` | Append operations |
| `Prepend`, `PrependLines`, `CreateOrAppend`, `CreateOrPrepend` | Prepend operations |
| `Rename`, `Change`, `Merge`, `MergeLines`, `Touch` | Modification operations |
| `Overwrite`, `Override`, `Enforce` | Override strategies |

### HTTP Methods

| Value | Description |
|---|---|
| `GetHttp`, `PutHttp`, `PostHttp`, `DeleteHttp`, `PatchHttp` | REST verbs |

### Service Lifecycle

| Value | Description |
|---|---|
| `Start`, `Stop`, `Restart`, `Reload`, `StopSleepStart` | Basic lifecycle |
| `Suspend`, `Pause`, `Resumed` | State transitions |
| `TryRestart3Times`, `TryRestart5Times`, `TryStart3Times`, etc. | Retry variants |

### Inheritance

| Value | Description |
|---|---|
| `InheritOnly`, `InheritPlusOverride`, `DynamicAction` | Inheritance strategies |

## Logical Grouping Methods

| Method | Description |
|---|---|
| `IsCreateLogically()` | Create, CreateOrUpdate, CreateOrSkipOnExist, DropCreate |
| `IsCreateOrUpdateLogically()` | All create + update variants |
| `IsDropLogically()` | All drop/delete variants |
| `IsCrudOnlyLogically()` | Standard CRUD + conditional variants |
| `IsReadOrEditLogically()` | Read + update + rename + change |
| `IsOverrideOrOverwriteOrEnforce()` | All override strategies |

## Usage Examples

```go
import "github.com/alimtvnetwork/core/reqtype"

req := reqtype.Create

// Type checks
req.IsCreate()              // true
req.IsCreateLogically()     // true
req.IsCrudOnlyLogically()   // true

// Name and value
req.Name()                  // "CreateUsingAliasMap"
req.Value()                 // 1

// Range utilities
err := reqtype.RangesOnlySupportedFor(
    "this operation",
    reqtype.Create, reqtype.Update,
)

// Ranges as strings
csv := reqtype.RangesStringDefaultJoiner(
    reqtype.Create, reqtype.Read,
)
```

## Related Docs

- [enumimpl](../coreimpl/enumimpl/readme.md) — underlying enum implementation
- [errcore](../errcore/readme.md) — error types used by range helpers
