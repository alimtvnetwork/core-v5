# defaulterr

Pre-built sentinel error variables for common failure scenarios. Each error is typed via `errcore` error types for consistent categorization across the codebase.

## Architecture

```
defaulterr/
└── defaulterr.go   # All sentinel error variables
```

## Error Variables

| Variable | Error Type | Description |
|---|---|---|
| `Marshalling` | `MarshallingFailedType` | Cannot marshal object to serialized form |
| `UnMarshalling` | `UnMarshallingFailedType` | Cannot unmarshal data to object form |
| `UnMarshallingPlusCannotFindingEnumMap` | `UnMarshallingFailedType` | Cannot find value in enum map |
| `MarshallingFailedDueToNilOrEmpty` | `UnMarshallingFailedType` | Nil/empty object prevents marshalling |
| `UnmarshallingFailedDueToNilOrEmpty` | `UnMarshallingFailedType` | Nil/empty data prevents unmarshalling |
| `CannotProcessNilOrEmpty` | `CannotBeNilOrEmptyType` | Generic nil/empty rejection |
| `OutOfRange` | `OutOfRangeType` | Value out of valid range |
| `NegativeDataCannotProcess` | `CannotBeNegativeType` | Negative value rejection |
| `NilResult` | `NullResultType` | Nil result encountered |
| `UnexpectedValue` | `UnexpectedValueType` | Unexpected value(s) encountered |
| `CannotRemoveFromEmptyCollection` | `CannotRemoveIndexesFromEmptyCollectionType` | Remove from empty collection |
| `CannotConvertStringToByte` | `FailedToConvertType` | String-to-byte conversion failure |
| `AttributeNull` | `NullResultType` | Nil attribute |
| `JsonResultNull` | `CannotBeNilOrEmptyType` | Nil JSON result |
| `KeyNotExistInMap` | `KeyNotExistInMapType` | Map key lookup failure |

## Usage Examples

```go
import "github.com/alimtvnetwork/core/defaulterr"

if data == nil {
    return defaulterr.CannotProcessNilOrEmpty
}

if index > length {
    return defaulterr.OutOfRange
}
```

## Related Docs

- [errcore](../errcore/readme.md) — error type definitions used to construct these sentinels
