# corejson — JSON Serialize/Deserialize Pipeline

Package `corejson` provides a complete JSON serialization and deserialization pipeline with rich error handling, type-safe results, and the struct-as-namespace pattern.

## Core Types

| Type | Description |
|------|-------------|
| `Result` | JSON bytes + error, with safe accessors and pretty-print |
| `ResultsCollection` | Collection of `Result` items |
| `BytesCollection` | Lightweight collection of byte slices |

## Entry Points

| Namespace | Description |
|-----------|-------------|
| `corejson.Serialize.*` | Serialize any value to JSON (bytes, string, result) |
| `corejson.Deserialize.*` | Deserialize JSON bytes/string into Go types |
| `corejson.New(value)` | Create a `Result` from any value |
| `corejson.NewPtr(value)` | Create a `*Result` from any value |
| `corejson.NewResult.*` | Advanced result creation (from bytes, errors, types) |
| `corejson.AnyTo.*` | Convert any type to JSON result |
| `corejson.Empty.*` | Empty result/collection factories |

## Usage

### Serialization

```go
import "gitlab.com/auk-go/core/coredata/corejson"

type User struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email,omitempty"`
}

user := User{Name: "Alice", Age: 30}

// To JSON string
jsonStr, err := corejson.Serialize.ToString(user)
// `{"name":"Alice","age":30}`

// To JSON bytes
jsonBytes, err := corejson.Serialize.Raw(user)

// To Result (bytes + error in one object)
result := corejson.New(user)
```

### Deserialization

```go
var restored User

// From bytes
err := corejson.Deserialize.UsingBytes(jsonBytes, &restored)

// From string
err = corejson.Deserialize.UsingString(jsonStr, &restored)

// Must variant (panics on error)
corejson.Deserialize.UsingBytesMust(jsonBytes, &restored)

// Deep copy via JSON round-trip
source := User{Name: "Bob", Age: 25}
target := User{}
err = corejson.Deserialize.FromTo(source, &target)
```

### Result Type

```go
result := corejson.NewPtr(user)

// Safe access
fmt.Println(result.HasError())         // false
fmt.Println(result.HasIssuesOrEmpty()) // false
bytes := result.SafeValues()           // []byte — never nil
jsonStr := result.JsonString()         // string
pretty := result.PrettyJsonString()    // formatted string

// Error handling
result.HandleError() // logs error if present

// Unmarshal from result
var another User
err := result.Deserialize(&another)
```

### Error Handling

```go
// Invalid input produces error result
badResult := corejson.New(make(chan int))
fmt.Println(badResult.HasError())    // true
fmt.Println(badResult.ErrorString()) // marshaling error

// Meaningful errors
err := badResult.MeaningfulError()
```

## Related Docs

- [Folder Spec](/spec/01-app/folders/05-coredata.md)
- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
