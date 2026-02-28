# corepayload — Structured Data Transport

Package `corepayload` provides the primary structured data transport system. `PayloadWrapper` carries named, identified payloads with attributes, authentication, and error handling.

## Core Types

| Type | Description |
|------|-------------|
| `PayloadWrapper` | Primary data transport — name, ID, entity, category, JSON payloads, attributes |
| `TypedPayloadWrapper[T]` | Generic wrapper — deserializes payloads into typed `T` with GetAs*, Value*, JSON ops |
| `Attributes` | Key-value pairs, auth info, paging, error wrapper |
| `PayloadsCollection` | Collection of `PayloadWrapper` items |
| `PayloadCreateInstruction` | Builder for creating PayloadWrapper instances |

## Architecture

```
corepayload/
├── TypedPayloadWrapper.go              # Generic: TypedPayloadWrapper[T]  (→ PayloadWrapper)
├── newTypedPayloadWrapperCreator.go    # Generic factory functions (package-level)
├── PayloadWrapper.go                   # Legacy:  PayloadWrapper          (any-based)
├── newPayloadWrapperCreator.go         # Legacy factory: New.PayloadWrapper.*
├── generic_helpers.go                  # Generic helpers: DeserializePayloadTo[T], etc.
├── Attributes.go                       # Key-value pairs, auth, paging
├── PayloadsCollection.go               # Collection of wrappers
└── newCreator.go                       # New Creator root aggregator
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
  ├─ .TypedDataJson() / .TypedDataJsonPtr()
  ├─ .SetTypedData(T) / .Reparse()
  ├─ .ClonePtr() / .Clone()
  ├─ .HasError() / .Error() / .HandleError()
  ├─ .Attributes() / .InitializeAttributesOnNull()
  ├─ .Clear() / .Dispose()
  └─ .ToPayloadWrapper()
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

// Clone
cloned, err := typed.ClonePtr(true)  // deep clone

// Deserialize from raw JSON bytes
typed, err = corepayload.TypedPayloadWrapperDeserialize[User](rawBytes)
typedSlice, err := corepayload.TypedPayloadWrapperDeserializeToMany[User](rawBytes)

// Access underlying wrapper
wrapper := typed.ToPayloadWrapper()
```

### PayloadWrapper — Standard Usage

```go
import "gitlab.com/auk-go/core/coredata/corepayload"

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
```

### Serialize / Deserialize Wrapper

```go
// Serialize entire wrapper
jsonBytes, err := payload.Serialize()

// Deserialize from bytes
restored, err := corepayload.New.PayloadWrapper.Deserialize(jsonBytes)

// Typed deserialization
typed, err := corepayload.TypedPayloadWrapperDeserialize[User](jsonBytes)
```

## Related Docs

- [Data Transport Architecture](/spec/01-app/folders/05-coredata.md)
- [newCreator Convention](/spec/01-app/18-new-creator-convention.md)
- [Go Modernization Plan](/spec/01-app/11-go-modernization.md)
- [coredynamic README](/coredata/coredynamic/README.md)
- [coreapi README](/coredata/coreapi/README.md)
