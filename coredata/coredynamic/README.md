# coredynamic — Dynamic Type Wrappers & Generic Collections

Package `coredynamic` provides dynamic type wrappers for runtime values, strongly-typed generic collections, and reflection-based utilities. It follows the **generic-first** principle.

## Architecture

```
coredynamic/
├── TypedDynamic.go              # Generic: TypedDynamic[T]        (→ Dynamic)
├── TypedSimpleRequest.go        # Generic: TypedSimpleRequest[T]  (→ SimpleRequest)
├── TypedSimpleResult.go         # Generic: TypedSimpleResult[T]   (→ SimpleResult)
├── Dynamic.go                   # Legacy:  Dynamic                (any-based, reflection)
├── DynamicGetters.go            # Dynamic read-only accessors, type checks, value extraction
├── DynamicReflect.go            # Dynamic reflection ops, loops, filters, conversion
├── DynamicJson.go               # Dynamic JSON serialization/deserialization
├── DynamicStatus.go             # Dynamic status helpers
├── SimpleRequest.go             # Legacy:  SimpleRequest          (any-based)
├── SimpleResult.go              # Legacy:  SimpleResult           (any-based)
├── Collection.go                # Generic: Collection[T]          (thread-safe list)
├── CollectionMethods.go         # Collection mutators, clone, capacity, reorder, search
├── CollectionMap.go             # Package-level Map, FlatMap, Reduce
├── CollectionDistinct.go        # Package-level Distinct, Unique, DistinctCount, IsDistinct
├── CollectionSort.go            # SortFunc, SortAsc/Desc, IsSorted (cmp.Ordered)
├── CollectionGroupBy.go         # GroupBy operations
├── CollectionLock.go            # Mutex-protected method variants
├── CollectionSearch.go          # Search and query methods
├── CollectionTypes.go           # Type-specific collection helpers
├── DynamicCollection.go         # Legacy:  DynamicCollection
├── AnyCollection.go             # Legacy:  AnyCollection
├── KeyVal.go                    # Dynamic key-value pair
├── KeyValCollection.go          # Collection of KeyVal
├── LeftRight.go                 # Left/Right pair wrapper
├── MapAnyItems.go               # Dynamic map with paging
├── CastTo.go / CastedResult.go  # Type casting utilities
├── ReflectSetFromTo.go          # Reflection-based assignment
├── newCreator.go                # New Creator pattern
└── vars.go                      # Package-level variables (New)
```

## Type Hierarchy

```
Generic (type-safe, recommended)              Legacy (any-based, backward compat)
──────────────────────────────                ──────────────────────────────────
TypedDynamic[T]                               Dynamic
  ├─ .Data() / .Value() T                       └─ .Data() / .Value() any
  ├─ .GetAs*(String/Int/Int64/Uint/Float64/Float32/Bool/Bytes/Strings)
  ├─ .Value*(String/Int/Int64/Bool)
  ├─ .Json() / .JsonPtr() / .JsonResult() / .JsonBytes() / .JsonString()
  ├─ .MarshalJSON() / .UnmarshalJSON() / .ValueMarshal()
  ├─ .Bytes() / .Deserialize()
  ├─ .Clone() / .ClonePtr() / .NonPtr() / .Ptr()
  ├─ .JsonModel() / .JsonModelAny()
  └─ .ToDynamic()

TypedSimpleRequest[T]                         SimpleRequest
  ├─ .Data() / .Request() / .Value() T          └─ .Request() / .Value() any
  ├─ .GetAs*(String/Int/Int64/Float64/Float32/Bool/Bytes/Strings)
  ├─ .InvalidError() / .Message() / .String()
  ├─ .Json() / .JsonPtr() / .JsonBytes() / .JsonResult()
  ├─ .MarshalJSON() / .JsonModel() / .JsonModelAny()
  ├─ .Clone()
  ├─ .ToTypedDynamic() / .ToDynamic()
  └─ .ToSimpleRequest()

TypedSimpleResult[T]                          SimpleResult
  ├─ .Data() / .Result() T                      └─ .Result any (field)
  ├─ .GetAs*(String/Int/Int64/Float64/Bool/Bytes/Strings)
  ├─ .InvalidError() / .Message() / .String()
  ├─ .Json() / .JsonPtr() / .JsonBytes() / .JsonResult()
  ├─ .MarshalJSON() / .JsonModel() / .JsonModelAny()
  ├─ .Clone() / .ClonePtr()
  ├─ .ToTypedDynamic() / .ToDynamic()
  └─ .ToSimpleResult()
```

## Usage

### TypedDynamic[T] — Generic Wrapper (Recommended)

```go
import "gitlab.com/auk-go/core/coredata/coredynamic"

// Create a typed dynamic value
d := coredynamic.NewTypedDynamic[string]("hello", true)
fmt.Println(d.Data())    // "hello" (typed as string)
fmt.Println(d.Value())   // "hello" (alias for Data)
fmt.Println(d.IsValid()) // true

// GetAs* type assertion helpers
str, ok := d.GetAsString()     // "hello", true
num, ok := d.GetAsInt()        // 0, false
uid, ok := d.GetAsUint()       // 0, false
f32, ok := d.GetAsFloat32()    // 0, false

// Value* convenience methods
fmt.Println(d.ValueString())   // "hello"
fmt.Println(d.ValueInt())      // -1 (InvalidValue)
fmt.Println(d.ValueInt64())    // -1 (InvalidValue)
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

// Clone
cloned := req.Clone()

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
clone := result.Clone()
clonePtr := result.ClonePtr()

// Conversions
legacyResult := result.ToSimpleResult()
typedDynamic := result.ToTypedDynamic()
legacyDynamic := result.ToDynamic()
```

### Collection[T] — Generic Collections

```go
// Create collections via the New creator pattern
col := coredynamic.New.Collection.String.Cap(10)
col.Add("hello")
col.Add("world")

fmt.Println(col.Length())  // 2
fmt.Println(col.First())  // "hello"

// Map, FlatMap, Reduce (package-level generic functions)
mapped := coredynamic.Map(col, func(s string) int {
    return len(s)
})

reduced := coredynamic.Reduce(col, "", func(acc, item string) string {
    return acc + item
})

flattened := coredynamic.FlatMap(col, func(s string) []rune {
    return []rune(s)
})

// Distinct (requires comparable T)
distinct := coredynamic.Distinct(col)
count := coredynamic.DistinctCount(col)

// Sort (requires cmp.Ordered T)
coredynamic.SortAsc(col)
coredynamic.SortDesc(col)
sorted := coredynamic.SortedAsc(col)  // non-mutating

// Sort with custom comparator (any T)
col.SortFunc(func(a, b string) bool { return a < b })

// Reverse, Filter, Clone
col.Reverse()
filtered := col.Filter(func(s string) bool { return len(s) > 3 })
cloned := col.Clone()
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
d.IsPointer()
d.IsStringType()
d.IsStruct()

// Value extraction
d.ValueString()
d.ValueInt()
d.ValueInt64()
d.ValueBool()
d.ValueStrings()
rawBytes, ok := d.Bytes()
```

## Related Docs

- [Folder Spec](/spec/01-app/folders/05-coredata.md)
- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
- [coreapi README](/coredata/coreapi/README.md)
- [Go Modernization Plan](/spec/01-app/11-go-modernization.md)
