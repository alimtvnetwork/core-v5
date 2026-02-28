# coreapi — Typed API Request/Response

Package `coreapi` provides structured request and response types for API communication. It includes both **legacy dynamic** (`any`-based) and **modern generic** (`[T]`-based) types.

## Types

### Generic (Typed) — Recommended

| Type | Description |
|------|-------------|
| `TypedRequestIn[T]` | Strongly-typed incoming request with `T` payload |
| `TypedRequest[T]` | Strongly-typed request wrapping `T` directly |
| `TypedResponse[T]` | Strongly-typed response with `T` payload |
| `TypedResponseResult[T]` | Strongly-typed response result |

### Legacy (Dynamic)

| Type | Description |
|------|-------------|
| `GenericRequestIn` | Request with `any` payload |
| `GenericResponse` | Response with `any` payload |
| `SimpleGenericRequest` | Request wrapping `*coredynamic.SimpleRequest` |
| `GenericResponseResult` | Response wrapping `*coredynamic.SimpleResult` |

### Supporting Types

| Type | Description |
|------|-------------|
| `RequestAttribute` | URL, host, resource, action, auth, search, paging metadata |
| `ResponseAttribute` | HTTP code/method, count, validity, steps, debug info |
| `SearchRequest` | Search term with match mode flags |
| `PageRequest` | Page size and index for pagination |
| `PayloadsRequestIn` | Raw byte payload request |

## Usage

### Generic Request/Response (Recommended)

```go
import "gitlab.com/auk-go/core/coredata/coreapi"

type UserInput struct {
    Name  string
    Email string
}

// Create a typed request
req := coreapi.NewTypedRequestIn[UserInput](
    &coreapi.RequestAttribute{
        Url:          "/api/users",
        ResourceName: "User",
        ActionName:   "Create",
        IsValid:      true,
    },
    UserInput{Name: "Alice", Email: "alice@example.com"},
)

fmt.Println(req.Request.Name)  // "Alice" — compile-time safe
fmt.Println(req.Request.Email) // "alice@example.com"

// Create a typed response
type UserOutput struct {
    ID   int
    Name string
}

resp := coreapi.NewTypedResponse[UserOutput](
    &coreapi.ResponseAttribute{IsValid: true, HttpCode: 200},
    UserOutput{ID: 1, Name: "Alice"},
)

// Clone
clone := req.Clone()

// Backward compatibility
legacyReq := req.ToGenericRequestIn()
```

### Invalid Requests/Responses

```go
invalidReq := coreapi.InvalidTypedRequestIn[UserInput](nil)
invalidResp := coreapi.InvalidTypedResponse[UserOutput](nil)
```

## Related Docs

- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
- [Core API Folder Spec](/spec/01-app/folders/05-coredata.md)
