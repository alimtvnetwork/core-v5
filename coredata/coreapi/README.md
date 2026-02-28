# coreapi — Typed API Request/Response

Package `coreapi` provides structured request and response types for API communication using **generic** (`[T]`-based) types for compile-time type safety.

## Architecture

```
coreapi/
├── TypedRequestIn.go              # TypedRequestIn[T] — generic incoming request
├── TypedRequest.go                # TypedRequest[T] — generic request wrapper
├── TypedResponse.go               # TypedResponse[T] — generic response
├── TypedResponseResult.go         # TypedResponseResult[T] — generic response result
├── TypedSimpleGenericRequest.go   # TypedSimpleGenericRequest[T] — wraps TypedSimpleRequest[T]
├── InvalidRequestAttribute.go     # Invalid factory for RequestAttribute
├── InvalidResponseAttribute.go    # Invalid factory for ResponseAttribute
├── RequestAttribute.go            # URL, host, resource, action, auth, search, paging
├── ResponseAttribute.go           # HTTP code/method, count, validity, steps, debug
├── SearchRequest.go               # Search term + match mode flags
├── PageRequest.go                 # Page size + index for pagination
├── PayloadsRequestIn.go           # Raw byte payload request
└── README.md
```

## Type Hierarchy

```
TypedRequestIn[T]
  ├─ .Request T
  ├─ .Attribute *RequestAttribute
  ├─ .Clone()
  └─ .TypedSimpleGenericRequest(isValid, msg)

TypedRequest[T]
  ├─ .Request T
  ├─ .Clone()
  └─ .ToTypedSimpleGenericRequest(isValid, msg)

TypedResponse[T]
  ├─ .Response T
  ├─ .Clone()
  └─ .TypedResponseResult()

TypedResponseResult[T]
  ├─ .Response T
  ├─ .Clone() / .ClonePtr()
  ├─ .IsValid() / .IsInvalid() / .Message()
  └─ .ToTypedResponse()

TypedSimpleGenericRequest[T]
  ├─ .Attribute *RequestAttribute
  ├─ .Request *TypedSimpleRequest[T]
  ├─ .IsValid() / .IsInvalid()
  ├─ .Data() / .Message() / .InvalidError()
  └─ .Clone()
```

## Types

### Generic (Typed)

| Type | Description |
|------|-------------|
| `TypedRequestIn[T]` | Strongly-typed incoming request with `T` payload |
| `TypedRequest[T]` | Strongly-typed request wrapping `T` directly |
| `TypedResponse[T]` | Strongly-typed response with `T` payload |
| `TypedResponseResult[T]` | Strongly-typed response result with validity/message |
| `TypedSimpleGenericRequest[T]` | Request wrapping `TypedSimpleRequest[T]` with validation |

### Supporting Types

| Type | Description |
|------|-------------|
| `RequestAttribute` | URL, host, resource, action, auth, search, paging metadata |
| `ResponseAttribute` | HTTP code/method, count, validity, steps, debug info |
| `SearchRequest` | Search term with match mode flags |
| `PageRequest` | Page size and index for pagination |
| `PayloadsRequestIn` | Raw byte payload request |

## Usage

### Generic Request/Response

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

// Convert to TypedSimpleGenericRequest
typedSimpleReq := req.TypedSimpleGenericRequest(true, "")
```

### Invalid Requests/Responses

```go
invalidReq := coreapi.InvalidTypedRequestIn[UserInput](nil)
invalidResp := coreapi.InvalidTypedResponse[UserOutput](nil)
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
