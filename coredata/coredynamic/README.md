# coredynamic — Dynamic Type Wrappers & Generic Collections

Package `coredynamic` provides dynamic type wrappers for runtime values, strongly-typed generic collections, and reflection-based utilities.

## Core Types

| Type | Description |
|------|-------------|
| `Dynamic` | Wraps `any` with reflection, type inspection, and conversion |
| `TypedDynamic[T]` | Generic, compile-time safe wrapper for typed values |
| `SimpleRequest` | Dynamic request with validity flag and message |
| `TypedSimpleRequest[T]` | Generic typed request with `GetAs*` methods |
| `Collection[T]` | Generic, thread-safe collection with LINQ-style operations |
| `DynamicCollection` | Collection of `Dynamic` values |

## Usage

### TypedDynamic[T] — Generic Wrapper

```go
import "gitlab.com/auk-go/core/coredata/coredynamic"

// Create a typed dynamic value
d := coredynamic.NewTypedDynamic[string]("hello", true)
fmt.Println(d.Data())    // "hello" (typed as string)
fmt.Println(d.IsValid()) // true

// Serialize to JSON
bytes, err := d.JsonBytes()

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
fmt.Println(req.IsValid())   // true

// GetAs* type assertion helpers
str, ok := req.GetAsString()     // false (T is UserInput)
val, ok := req.GetAsInt()        // false
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
