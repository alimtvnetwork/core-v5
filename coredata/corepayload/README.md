# corepayload — Structured Data Transport

Package `corepayload` provides the primary structured data transport system. `PayloadWrapper` carries named, identified payloads with attributes, authentication, and error handling.

## Entry Points

| Variable | Type | Description |
|----------|------|-------------|
| `New` | `newCreator` | Builder-pattern factory for all payload types |
| `Empty` | `emptyCreator` | Quick empty-instance factory |

## Core Types

| Type | Description |
|------|-------------|
| `PayloadWrapper` | Primary data transport — name, ID, entity, category, JSON payloads, attributes |
| `TypedPayloadWrapper[T]` | Generic wrapper — deserializes payloads into typed `T` with GetAs*, Value*, JSON ops |
| `TypedPayloadCollection[T]` | Generic thread-safe collection of `TypedPayloadWrapper[T]` with ForEach, Filter, AllData |
| `Attributes` | Key-value pairs, auth info, paging, error wrapper, dynamic payloads |
| `PayloadsCollection` | Collection of `PayloadWrapper` items |
| `PayloadCreateInstruction` | Builder for creating PayloadWrapper instances |
| `AuthInfo` | Authentication information container |
| `PagingInfo` | Pagination metadata |
| `SessionInfo` | Session information container |
| `User` | User identity model |
| `UserInfo` | User details with system user support |

## Architecture

```
corepayload/
├── TypedPayloadWrapper.go              # Generic: TypedPayloadWrapper[T]  (→ PayloadWrapper)
├── TypedPayloadCollection.go           # Generic: TypedPayloadCollection[T] (→ PayloadsCollection)
├── typed_collection_funcs.go           # Generic collection helper functions
├── newTypedPayloadWrapperCreator.go    # Generic factory functions (package-level)
├── PayloadWrapper.go                   # Legacy:  PayloadWrapper          (any-based)
├── PayloadWrapperGetters.go            # PayloadWrapper read-only accessors
├── PayloadWrapperJson.go               # PayloadWrapper JSON operations
├── newPayloadWrapperCreator.go         # Legacy factory: New.PayloadWrapper.*
├── generic_helpers.go                  # Generic helpers: DeserializePayloadTo[T], etc.
├── Attributes.go                       # Key-value pairs, auth, paging, error
├── AttributesGetters.go                # Attributes read-only accessors
├── AttributesSetters.go                # Attributes mutation methods
├── AttributesJson.go                   # Attributes JSON serialization
├── PayloadsCollection.go               # Collection of wrappers
├── newPayloadsCollectionCreator.go     # PayloadsCollection factory
├── AuthInfo.go                         # Authentication info
├── PagingInfo.go                       # Pagination metadata
├── SessionInfo.go                      # Session info
├── User.go / UserInfo.go              # User identity models
├── emptyCreator.go                     # Empty-instance factory
├── newCreator.go                       # New Creator root aggregator
└── vars.go                            # Package-level variables (New, Empty)
```

## New Creator Pattern

```go
// New.PayloadWrapper — primary factory
New.PayloadWrapper.Record(name, id, taskName, category, record)
New.PayloadWrapper.Records(name, id, taskName, category, records)
New.PayloadWrapper.NameIdRecord(name, id, record)
New.PayloadWrapper.NameIdCategory(name, id, category, record)
New.PayloadWrapper.UsingCreateInstruction(instruction)
New.PayloadWrapper.Deserialize(rawBytes)
New.PayloadWrapper.Empty()

// New.PayloadsCollection
New.PayloadsCollection.Deserialize(rawBytes)

// New.Attributes
New.Attributes.Empty()
New.Attributes.All(authInfo, kvPairs, anyKV, pagingInfo, dynamicPayloads, fromTo, basicErr)

// New.User
New.User.*(...)
```

## Type Hierarchy

```
Generic (type-safe, recommended)              Legacy (any-based, backward compat)
──────────────────────────────                ──────────────────────────────────
TypedPayloadWrapper[T]                        PayloadWrapper
  ├─ .TypedData() / .Data() T                   └─ .Value() any / .Payloads []byte
  ├─ .GetAs*(String/Int/Int64/Float64/Float32/Bool/Bytes/Strings)
  ├─ .Value*(String/Int/Bool)
  ├─ .Json() / .JsonPtr() / .JsonString() / .PrettyJsonString()
  ├─ .MarshalJSON() / .UnmarshalJSON()
  ├─ .Serialize() / .SerializeMust()
  ├─ .TypedDataJson() / .TypedDataJsonPtr() / .TypedDataJsonBytes()
  ├─ .SetTypedData(T) / .SetTypedDataMust(T) / .Reparse()
  ├─ .SetName() / .SetIdentifier() / .SetEntityType() / .SetCategoryName()
  ├─ .ClonePtr(bool) / .Clone(bool)
  ├─ .HasError() / .Error() / .HandleError()
  ├─ .Name() / .Identifier() / .IdString() / .IdInteger()
  ├─ .EntityType() / .CategoryName() / .TaskTypeName()
  ├─ .Attributes() / .InitializeAttributesOnNull()
  ├─ .Length() / .IsNull() / .IsEmpty() / .HasItems() / .HasSafeItems()
  ├─ .DynamicPayloads() / .PayloadsString()
  ├─ .Clear() / .Dispose()
  └─ .ToPayloadWrapper() / .PayloadWrapperValue()

TypedPayloadCollection[T]                     PayloadsCollection
  ├─ .Items() / .Length() / .IsEmpty() / .HasItems()
  ├─ .First() / .Last() / .FirstOrDefault() / .LastOrDefault() / .SafeAt()
  ├─ .Add() / .AddLock() / .Adds() / .AddCollection() / .RemoveAt()
  ├─ .ForEach() / .ForEachData() / .ForEachBreak()
  ├─ .Filter() / .FilterByData() / .FirstByFilter() / .FirstByData()
  ├─ .FirstByName() / .FirstById() / .CountFunc()
  ├─ .Skip() / .Take()
  ├─ .AllData() / .AllNames() / .AllIdentifiers()
  ├─ .Clone() / .CloneMust() / .ConcatNew()
  ├─ .Clear() / .Dispose()
  └─ .ToPayloadsCollection()
```

## Usage

### TypedPayloadWrapper[T] — Generic (Recommended)

```go
type User struct {
    Name  string
    Email string
}

// Create from existing PayloadWrapper
typed, err := corepayload.NewTypedPayloadWrapper[User](wrapper)
fmt.Println(typed.TypedData().Name)  // strongly typed — no assertions

// Create directly from typed data
typed, err = corepayload.NewTypedPayloadWrapperFrom[User](
    "user-create", "usr-123", "User",
    User{Name: "Alice", Email: "alice@example.com"},
)

// Factory functions (package-level, mirror New.PayloadWrapper.*)
typed, err = corepayload.TypedPayloadWrapperRecord[User](
    "user-create", "usr-123", "task", "category",
    User{Name: "Alice"},
)
typed, err = corepayload.TypedPayloadWrapperNameIdRecord[User](
    "user-create", "usr-123", User{Name: "Alice"},
)
typed, err = corepayload.TypedPayloadWrapperAll[User](
    "name", "id", "task", "User", "category",
    false, myUser, myAttrs,
)

// GetAs* helpers
str, ok := typed.GetAsString()
num, ok := typed.GetAsInt()
f64, ok := typed.GetAsFloat64()

// Value* convenience (with safe defaults)
fmt.Println(typed.ValueString())  // fmt fallback
fmt.Println(typed.ValueInt())     // InvalidValue fallback

// JSON ops on typed data specifically
jsonResult := typed.TypedDataJson()
jsonBytes, err := typed.TypedDataJsonBytes()

// Mutate typed data
err = typed.SetTypedData(User{Name: "Bob"})

// Mutate metadata
typed.SetName("user-update")
typed.SetIdentifier("usr-456")

// Clone
cloned, err := typed.ClonePtr(true)  // deep clone

// Deserialize from raw JSON bytes
typed, err = corepayload.TypedPayloadWrapperDeserialize[User](rawBytes)
typedSlice, err := corepayload.TypedPayloadWrapperDeserializeToMany[User](rawBytes)

// Access underlying wrapper
wrapper := typed.ToPayloadWrapper()
```

### TypedPayloadCollection[T] — Generic Collection

```go
// Create
col := corepayload.NewTypedPayloadCollection[User](10)
col.Add(typedWrapper)

// From existing PayloadsCollection
col = corepayload.TypedPayloadCollectionFromPayloads[User](payloadsCol)

// Iterate
col.ForEach(func(index int, item *corepayload.TypedPayloadWrapper[User]) {
    fmt.Println(item.Data().Name)
})
col.ForEachData(func(index int, data User) {
    fmt.Println(data.Name)
})

// Filter
admins := col.FilterByData(func(u User) bool { return u.Role == "admin" })
first := col.FirstByName("user-create")
byId := col.FirstById("usr-123")

// Extract all typed data
allUsers := col.AllData()  // []User
allNames := col.AllNames() // []string

// Convert back to legacy
legacy := col.ToPayloadsCollection()
```

### PayloadWrapper — Standard Usage

```go
import "github.com/alimtvnetwork/core/coredata/corepayload"

// Create via instruction
payload, err := corepayload.New.PayloadWrapper.UsingCreateInstruction(
    &corepayload.PayloadCreateInstruction{
        Name:       "user-create",
        Identifier: "usr-123",
        EntityType: "User",
        Payloads:   myStruct,  // auto-serialized to JSON bytes
    },
)

// Access metadata
fmt.Println(payload.PayloadName())       // "user-create"
fmt.Println(payload.IdString())          // "usr-123"
fmt.Println(payload.PayloadEntityType()) // "User"

// Deserialize payloads
var user User
err := payload.Deserialize(&user)

// Error handling
if payload.HasError() {
    log.Fatal(payload.Error())
}

// Attributes
attrs := payload.InitializeAttributesOnNull()
attrs.AddOrUpdateString("role", "admin")

// Clone
cloned, err := payload.ClonePtr(true) // deep clone
```

### Package-Level Generic Helpers

```go
// Deserialize without creating a TypedPayloadWrapper
user, err := corepayload.DeserializePayloadTo[User](wrapper)
users, err := corepayload.DeserializePayloadToSlice[User](wrapper)
user = corepayload.DeserializePayloadToMust[User](wrapper) // panics on error

// Attributes deserialization
config, err := corepayload.DeserializeAttributesPayloadTo[AppConfig](attrs)
config = corepayload.DeserializeAttributesPayloadToMust[AppConfig](attrs) // panics on error
configs, err := corepayload.DeserializeAttributesPayloadToSlice[AppConfig](attrs)
```

### Serialize / Deserialize Wrapper

```go
// Serialize entire wrapper
jsonBytes, err := payload.Serialize()

// Deserialize from bytes
restored, err := corepayload.New.PayloadWrapper.Deserialize(jsonBytes)

// Typed deserialization
typed, err := corepayload.TypedPayloadWrapperDeserialize[User](jsonBytes)
typedFromResult, err := corepayload.TypedPayloadWrapperDeserializeUsingJsonResult[User](jsonResult)
```

## Related Docs

- [Data Transport Architecture](/spec/01-app/folders/05-coredata.md)
- [newCreator Convention](/spec/01-app/18-new-creator-convention.md)
- [Go Modernization Plan](/spec/01-app/11-go-modernization.md)
- [coredynamic README](/coredata/coredynamic/README.md)
- [coreapi README](/coredata/coreapi/README.md)
