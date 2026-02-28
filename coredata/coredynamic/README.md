# coredynamic — Dynamic Type Wrappers & Generic Collections

Package `coredynamic` provides dynamic type wrappers for runtime values, strongly-typed generic collections, and reflection-based utilities. It follows the **generic-first** principle.

## Architecture

```
coredynamic/
├── TypedDynamic.go              # Generic: TypedDynamic[T]        (→ Dynamic)
├── TypedSimpleRequest.go        # Generic: TypedSimpleRequest[T]  (→ SimpleRequest)
├── TypedSimpleResult.go         # Generic: TypedSimpleResult[T]   (→ SimpleResult)
├── Dynamic.go                   # Legacy:  Dynamic                (any-based, reflection)
├── SimpleRequest.go             # Legacy:  SimpleRequest          (any-based)
├── SimpleResult.go              # Legacy:  SimpleResult           (any-based)
├── Collection.go                # Generic: Collection[T]          (thread-safe list)
├── DynamicCollection.go         # Legacy:  DynamicCollection
├── AnyCollection.go             # Legacy:  AnyCollection
├── KeyVal.go                    # Dynamic key-value pair
├── LeftRight.go                 # Left/Right pair wrapper
├── MapAnyItems.go               # Dynamic map with paging
└── newCreator.go                # New Creator pattern
```

## Type Hierarchy

```
Generic (type-safe, recommended)              Legacy (any-based, backward compat)
──────────────────────────────                ──────────────────────────────────
TypedDynamic[T]                               Dynamic
  ├─ .Data() T                                  └─ .Data() any
  ├─ .GetAs*(String/Int/Int64/Float64/Bool/Bytes/Strings)
  ├─ .Value*(String/Int/Int64/Bool)
  ├─ .Json() / .JsonPtr() / .JsonBytes()
  ├─ .MarshalJSON() / .UnmarshalJSON()
  ├─ .Bytes() / .Deserialize()
  ├─ .ClonePtr() / .NonPtr() / .Ptr()
  └─ .ToDynamic()

TypedSimpleRequest[T]                         SimpleRequest
  ├─ .Data() / .Request() / .Value() T          └─ .Data() any
  ├─ .GetAs*(String/Int/Int64/Float64/Float32/Bool/Bytes/Strings)
  ├─ .InvalidError() / .Message()
  ├─ .Json() / .JsonPtr() / .MarshalJSON()
  ├─ .ToTypedDynamic() / .ToDynamic()
  └─ .ToSimpleRequest()

TypedSimpleResult[T]                          SimpleResult
  ├─ .Data() / .Result() T                      └─ .Result any
  ├─ .GetAs*(String/Int/Int64/Float64/Bool/Bytes/Strings)
  ├─ .InvalidError() / .Message()
  ├─ .Json() / .JsonPtr() / .MarshalJSON()
  ├─ .ClonePtr() / .ToTypedDynamic()
  └─ .ToSimpleResult()
```

## Usage

### TypedDynamic[T] — Generic Wrapper (Recommended)

```go
import "gitlab.com/auk-go/core/coredata/coredynamic"

// Create a typed dynamic value
d := coredynamic.NewTypedDynamic[string]("hello", true)
fmt.Println(d.Data())    // "hello" (typed as string)
fmt.Println(d.IsValid()) // true

// GetAs* type assertion helpers
str, ok := d.GetAsString()     // "hello", true
num, ok := d.GetAsInt()        // 0, false

// Value* convenience methods
fmt.Println(d.ValueString())   // "hello"
fmt.Println(d.ValueInt())      // -1 (InvalidValue)
fmt.Println(d.ValueBool())     // false

// JSON operations
bytes, err := d.JsonBytes()
jsonStr, err := d.JsonString()
result := d.Json()
resultPtr := d.JsonPtr()

// Raw bytes
rawBytes, ok := d.Bytes()

// Deserialize from JSON
err = d.Deserialize([]byte(`"world"`))

// Clone
clone := d.Clone()
clonePtr := d.ClonePtr()

// Convert to legacy Dynamic
legacy := d.ToDynamic()
```

### TypedSimpleRequest[T] — Generic Request

```go
type UserInput struct {
    Name string
    Age  int
}

req := coredynamic.NewTypedSimpleRequestValid[UserInput](
    UserInput{Name: "Alice", Age: 30},
)

fmt.Println(req.Data().Name) // "Alice" — strongly typed
fmt.Println(req.Request())   // same as Data()
fmt.Println(req.IsValid())   // true

// Validation
if req.IsInvalid() {
    err := req.InvalidError()
    log.Fatal(req.Message())
}

// GetAs* (useful when T is any or interface type)
str, ok := req.GetAsString()
num, ok := req.GetAsInt64()

// JSON
jsonResult := req.Json()
jsonBytes, err := req.JsonBytes()

// Conversions
typedDynamic := req.ToTypedDynamic()
legacyDynamic := req.ToDynamic()
legacyRequest := req.ToSimpleRequest()
```

### TypedSimpleResult[T] — Generic Result

```go
type UserOutput struct {
    ID   int
    Name string
}

result := coredynamic.NewTypedSimpleResultValid[UserOutput](
    UserOutput{ID: 1, Name: "Alice"},
)

fmt.Println(result.Data().Name)   // "Alice" — compile-time safe
fmt.Println(result.Result().ID)   // 1 (alias for Data)
fmt.Println(result.IsValid())     // true

// Invalid result
invalidResult := coredynamic.InvalidTypedSimpleResult[UserOutput]("user not found")
fmt.Println(invalidResult.IsInvalid()) // true
fmt.Println(invalidResult.Message())   // "user not found"
err := invalidResult.InvalidError()    // errors.New("user not found")

// Clone
clone := result.ClonePtr()

// Conversions
legacyResult := result.ToSimpleResult()
typedDynamic := result.ToTypedDynamic()
```

### Collection[T] — Generic Collections

```go
// Create collections via the New creator pattern
col := coredynamic.New.Collection.String.Cap(10)
col.Add("hello")
col.Add("world")

fmt.Println(col.Length())  // 2
fmt.Println(col.First())  // "hello"

// Map, Filter, Reduce
mapped := coredynamic.Map[string, int](col, func(s string) int {
    return len(s)
})

reduced := coredynamic.Reduce[string, string](col, "", func(acc, item string) string {
    return acc + item
})

// Distinct, Sort, Reverse
col.Reverse()
distinct := coredynamic.Distinct[string](col)
```

### Dynamic — Legacy Wrapper

```go
d := coredynamic.NewDynamic(myValue, true)
fmt.Println(d.IsValid())
fmt.Println(d.ReflectTypeName())
fmt.Println(d.Length())   // for slices/maps/arrays

// Type checking
d.IsMap()
d.IsSliceOrArray()
d.IsPrimitive()
d.IsNumber()
```

## Related Docs

- [Folder Spec](/spec/01-app/folders/05-coredata.md)
- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
- [coreapi README](/coredata/coreapi/README.md)
- [Go Modernization Plan](/spec/01-app/11-go-modernization.md)
