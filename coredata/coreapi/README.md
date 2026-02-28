# coreapi — Typed API Request/Response

Package `coreapi` provides structured request and response types for API communication. It includes both **legacy dynamic** (`any`-based) and **modern generic** (`[T]`-based) types.

## Architecture

```
coreapi/
├── TypedRequestIn.go              # Generic: TypedRequestIn[T]  (→ GenericRequestIn)
├── TypedRequest.go                # Generic: TypedRequest[T]    (→ SimpleGenericRequest)
├── TypedResponse.go               # Generic: TypedResponse[T]   (→ GenericResponse)
├── TypedResponseResult.go         # Generic: TypedResponseResult[T] (→ GenericResponseResult)
├── TypedSimpleGenericRequest.go   # Generic: TypedSimpleGenericRequest[T]
├── GenericRequestIn.go            # Legacy:  GenericRequestIn   (any-based)
├── GenericResponse.go             # Legacy:  GenericResponse    (any-based)
├── SimpleGenericRequest.go        # Legacy:  SimpleGenericRequest
├── GenericResponseResult.go       # Legacy:  GenericResponseResult
├── RequestAttribute.go            # URL, host, resource, action, auth, search, paging
├── ResponseAttribute.go           # HTTP code/method, count, validity, steps, debug
├── SearchRequest.go               # Search term + match mode flags
├── PageRequest.go                 # Page size + index for pagination
└── PayloadsRequestIn.go           # Raw byte payload request
```

## Type Hierarchy

```
Generic (type-safe, recommended)              Legacy (any-based, backward compat)
──────────────────────────────                ──────────────────────────────────
TypedRequestIn[T]                             GenericRequestIn
  ├─ .Request T                                 └─ .Request any
  ├─ .Attribute *RequestAttribute
  ├─ .Clone()
  └─ .ToGenericRequestIn()

TypedRequest[T]                               SimpleGenericRequest
  ├─ .Request T                                 └─ .Request *SimpleRequest
  ├─ .ToSimpleGenericRequest()
  └─ .ToTypedSimpleGenericRequest()

TypedResponse[T]                              GenericResponse
  ├─ .Response T                                └─ .Response any
  ├─ .TypedResponseResult()
  ├─ .GenericResponseResult()
  └─ .ToGenericResponse()

TypedResponseResult[T]                        GenericResponseResult
  ├─ .Response T                                └─ .Response *SimpleResult
  └─ .ToGenericResponseResult()
```

## Types

### Generic (Typed) — Recommended

| Type | Description |
|------|-------------|
| `TypedRequestIn[T]` | Strongly-typed incoming request with `T` payload |
| `TypedRequest[T]` | Strongly-typed request wrapping `T` directly |
| `TypedResponse[T]` | Strongly-typed response with `T` payload |
| `TypedResponseResult[T]` | Strongly-typed response result |
| `TypedSimpleGenericRequest[T]` | Request wrapping `TypedSimpleRequest[T]` |

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

fmt.Println(resp.Response.ID)   // 1
fmt.Println(resp.Response.Name) // "Alice"

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

### Converting Between Generic and Legacy

```go
// Generic → Legacy
legacyReq := typedReq.ToGenericRequestIn()
legacyResp := typedResp.ToGenericResponse()
legacyResult := typedResp.GenericResponseResult()

// Generic → SimpleGenericRequest (wraps in SimpleRequest)
simpleReq := typedReq.ToSimpleGenericRequest(true, "")

// Generic → TypedSimpleGenericRequest (wraps in TypedSimpleRequest[T])
typedSimpleReq := typedReq.ToTypedSimpleGenericRequest(true, "")
```

### Pagination & Search

```go
req := &coreapi.RequestAttribute{
    Url:          "/api/users",
    ResourceName: "User",
    SearchRequest: &coreapi.SearchRequest{
        SearchTerm: "alice",
    },
    PageRequest: &coreapi.PageRequest{
        PageSize:  20,
        PageIndex: 0,
    },
}
```

## Related Docs

- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
- [Core API Folder Spec](/spec/01-app/folders/05-coredata.md)
- [coredynamic README](/coredata/coredynamic/README.md)
- [corepayload README](/coredata/corepayload/README.md)
