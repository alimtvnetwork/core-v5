# Phase 6: Value Receiver Migration — coredata/corepayload

## Date: 2026-03-02

## Rule
Convert pointer receiver methods to value receivers when:
1. The method is **read-only** (does not mutate the receiver)
2. The method does **not** check `it == nil` (nil-safe methods must remain pointer receivers)

Methods that check `it == nil` MUST stay pointer receivers because value receivers cannot be nil.

## Changes Applied

### PayloadWrapper.go (9 methods)
| Method | Rationale |
|---|---|
| `All()` | Returns fields, no nil check |
| `PayloadName()` | Returns `it.Name` |
| `PayloadCategory()` | Returns `it.CategoryName` |
| `PayloadTaskType()` | Returns `it.TaskTypeName` |
| `PayloadEntityType()` | Returns `it.EntityType` |
| `PayloadDynamic()` | Returns `it.Payloads` |
| `Value()` | Returns `it.Payloads` as `any` |
| `JsonModel()` | Returns value copy |
| `JsonModelAny()` | Delegates to `JsonModel()` |

### AttributesJson.go (8 methods)
| Method | Rationale |
|---|---|
| `JsonString()` | Delegates to `JsonPtr()` |
| `JsonStringMust()` | Delegates to `JsonPtr()` |
| `String()` | Delegates to `JsonString()` |
| `PrettyJsonString()` | Delegates to `JsonPtr()` |
| `Json()` | Returns `corejson.New(it)` |
| `JsonPtr()` | Returns `corejson.NewPtr(it)` |
| `JsonModel()` | Returns value copy (changed from `*Attributes` to `Attributes`) |
| `JsonModelAny()` | Delegates to `JsonModel()` |

### User.go (2 methods)
| Method | Rationale |
|---|---|
| `IdentifierInteger()` | Field access, no nil check |
| `IdentifierUnsignedInteger()` | Delegates to `IdentifierInteger()` |

### AuthInfo.go (7 methods)
| Method | Rationale |
|---|---|
| `IdentifierInteger()` | Field access, no nil check |
| `IdentifierUnsignedInteger()` | Delegates to above |
| `String()` | Delegates to `JsonPtr()` |
| `PrettyJsonString()` | Delegates to `JsonPtr()` |
| `Json()` | Returns `corejson.New(it)` |
| `JsonPtr()` | Returns `corejson.NewPtr(it)` |
| `Clone()` | Returns new struct (also fixed: was missing `Identifier` in clone) |

### SessionInfo.go (2 methods)
| Method | Rationale |
|---|---|
| `IdentifierInteger()` | Field access, no nil check |
| `IdentifierUnsignedInteger()` | Delegates to above |

### payloadProperties.go (4 methods)
| Method | Rationale |
|---|---|
| `BasicError()` | Delegates to wrapper |
| `IdInteger()` | Delegates to wrapper |
| `IdUnsignedInteger()` | Delegates to wrapper |
| `HasManyRecord()` | Returns field |

## Methods NOT Converted (must remain pointer receivers)
- All `IsEmpty()`, `IsNull()`, `HasError()` — check `it == nil`
- All `Set*()`, `Add*()`, `Clear()`, `Dispose()` — mutate receiver
- `ParseInjectUsingJson*()` — mutate receiver via Unmarshal
- `MarshalJSON()`, `UnmarshalJSON()` — interface contracts
- `ClonePtr()` — checks `it == nil`
- `PayloadWrapper.AllSafe()` — checks `it.IsNull()`

## Bug Fix
- `AuthInfo.Clone()` was missing `Identifier` field in the cloned struct
