# corepayload — Structured Data Transport

Package `corepayload` provides the primary structured data transport system. `PayloadWrapper` carries named, identified payloads with attributes, authentication, and error handling.

## Core Types

| Type | Description |
|------|-------------|
| `PayloadWrapper` | Primary data transport — name, ID, entity, category, JSON payloads, attributes |
| `TypedPayloadWrapper[T]` | Generic wrapper — deserializes payloads into typed `T` |
| `Attributes` | Key-value pairs, auth info, paging, error wrapper |
| `PayloadsCollection` | Collection of `PayloadWrapper` items |
| `PayloadCreateInstruction` | Builder for creating PayloadWrapper instances |

## Usage

### PayloadWrapper — Standard Usage

```go
import "gitlab.com/auk-go/core/coredata/corepayload"

// Create via instruction
payload := corepayload.New.PayloadWrapper.UsingInstruction(
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

// GetAs* helpers
str, ok := typed.GetAsString()
num, ok := typed.GetAsInt()

// Access underlying wrapper
wrapper := typed.ToPayloadWrapper()
```

### Package-Level Generic Helpers

```go
// Deserialize without creating a TypedPayloadWrapper
user, err := corepayload.DeserializePayloadTo[User](wrapper)
users, err := corepayload.DeserializePayloadToSlice[User](wrapper)
user = corepayload.DeserializePayloadToMust[User](wrapper) // panics on error
```

### Serialize / Deserialize Wrapper

```go
// Serialize entire wrapper
jsonBytes, err := payload.Serialize()

// Deserialize from bytes
restored, err := corepayload.New.PayloadWrapper.Deserialize(jsonBytes)
```

## Related Docs

- [Data Transport Architecture](/spec/01-app/folders/05-coredata.md)
- [newCreator Convention](/spec/01-app/18-new-creator-convention.md)
