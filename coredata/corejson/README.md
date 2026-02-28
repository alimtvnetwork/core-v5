# corejson — JSON Serialize/Deserialize Pipeline

Package `corejson` provides a complete JSON serialization and deserialization pipeline with rich error handling, type-safe results, and the struct-as-namespace pattern.

## Architecture

```
corejson/
├── vars.go                        # Package-level singletons: Serialize, Deserialize, CastAny, etc.
├── serializerLogic.go             # Serialize.* — any → JSON bytes/string/Result
├── deserializerLogic.go           # Deserialize.* — JSON → Go types
├── castingAny.go                  # CastAny.* — any-to-any via JSON round-trip
├── anyTo.go                       # AnyTo.* — any → Result conversion
├── Result.go                      # Result type: bytes + error + type name
├── ResultCollection.go            # ResultsCollection: slice of Result
├── ResultsPtrCollection.go        # ResultsPtrCollection: slice of *Result
├── BytesCollection.go             # BytesCollection: lightweight [][]byte wrapper
├── MapResults.go                  # MapResults: map[string]Result
├── New.go / NewPtr.go             # Quick constructors: New(val), NewPtr(val)
├── newResultCreator.go            # NewResult.* — advanced Result factories
├── emptyCreator.go                # Empty.* — zero-value factories
├── funcs.go                       # Standalone helpers: BytesToPrettyString, etc.
├── KeyAny.go                      # Key-value pair with any value
├── KeyWithJsoner.go               # Key-value pair with Jsoner interface
├── KeyWithResult.go               # Key-value pair with Result
└── all-interfaces.go              # Jsoner, SimpleJsoner, JsonStringer, etc.
```

## Core Types

| Type | Description |
|------|-------------|
| `Result` | JSON bytes + error, with safe accessors and pretty-print |
| `ResultsCollection` | Collection of `Result` items |
| `ResultsPtrCollection` | Collection of `*Result` items |
| `BytesCollection` | Lightweight collection of byte slices |
| `MapResults` | Map of string → Result with aggregate operations |

## Entry Points

| Namespace | Description |
|-----------|-------------|
| `corejson.Serialize.*` | Serialize any value to JSON (bytes, string, result) |
| `corejson.Deserialize.*` | Deserialize JSON bytes/string into Go types |
| `corejson.CastAny.*` | Cast any → any via JSON serialization round-trip |
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

// Using any value
result = corejson.Serialize.UsingAny(user)
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

// From Result
err = corejson.Deserialize.Apply(result, &restored)

// Deep copy via JSON round-trip
source := User{Name: "Bob", Age: 25}
target := User{}
err = corejson.Deserialize.FromTo(source, &target)
```

### CastAny — Type Casting via JSON

```go
// Cast any value to a target type via JSON serialize → deserialize round-trip.
// Useful when you have map[string]any or similar dynamic data.
var user User
err := corejson.CastAny.FromToDefault(dynamicMap, &user)

// Same as FromToDefault
err = corejson.CastAny.FromToReflection(dynamicMap, &user)
```

### Result Type

```go
result := corejson.NewPtr(user)

// Safe access
fmt.Println(result.HasError())         // false
fmt.Println(result.HasIssuesOrEmpty()) // false
bytes := result.SafeValues()           // []byte — never nil
bytes = result.SafeBytes()             // alias for SafeValues
jsonStr := result.JsonString()         // string
pretty := result.PrettyJsonString()    // formatted string

// Error handling
result.HandleError() // logs error if present
err := result.MeaningfulError()        // nil if no error, error otherwise

// Unmarshal from result
var another User
err = result.Deserialize(&another)
```

### MapResults — Named Result Collection

```go
mapResults := corejson.New.MapResults.Empty()
mapResults.Set("user", corejson.New(user))
mapResults.Set("config", corejson.New(config))

// Retrieve
result, has := mapResults.Get("user")
allKeys := mapResults.Keys()
errStrings := mapResults.GetErrorsStrings()
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
- [coredynamic README](/coredata/coredynamic/README.md)
- [corepayload README](/coredata/corepayload/README.md)
