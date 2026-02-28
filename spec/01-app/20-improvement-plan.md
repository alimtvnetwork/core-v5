# Improvement Plan — Phased Roadmap

> A prioritized, phase-by-phase plan for modernizing the `core` framework.

## Phase 1: Foundation (Current — In Progress)

### 1.1 Complete `interface{}` → `any` Migration
- **Status**: ~60% done
- **Remaining**: `coredata/corepayload/`, `coreinterface/`, `coretests/coretestcases/CaseV1.go`, remaining `internal/` files
- **Effort**: ~2-3 sessions
- **Risk**: Low (alias, no behavioral change)

### 1.2 Fix Known Bugs
- [ ] `PayloadWrapper.IsIdentifier` compares `it.Name` instead of `it.Identifier` (line 443)
- [ ] Remaining `convertinteranl` / `refeflectcore` typo references
- **Effort**: 1 session

### 1.3 Complete Go Version Update
- [ ] Update `go.mod` to Go 1.24
- [ ] Update `makefile` GoVersion
- [ ] Run `go mod tidy` and verify
- **Effort**: 1 session

---

## Phase 2: Generics — Core Collections

### 2.1 Create `Collection[T]` (New Generic Collection)
**Location**: `coredata/coredynamic/` (new file `Collection.go`)

```go
type Collection[T any] struct {
    items []T
}

func NewCollection[T any](capacity int) *Collection[T] {
    return &Collection[T]{
        items: make([]T, 0, capacity),
    }
}

func EmptyCollection[T any]() *Collection[T] {
    return NewCollection[T](0)
}
```

**Methods to include**:
- `Add(item T)`, `AddMany(items ...T)`, `AddNonNil(item *T)` (for pointer types)
- `At(index int) T`, `First() T`, `Last() T`, `FirstOrDefault() (T, bool)`
- `Items() []T`, `Length() int`, `IsEmpty() bool`, `HasAnyItem() bool`
- `Skip(n int) []T`, `Take(n int) []T`, `Limit(n int) []T`
- `SkipCollection(n int) *Collection[T]`, `TakeCollection(n int) *Collection[T]`
- `Loop(func(index int, item T) bool)`, `Filter(func(T) bool) *Collection[T]`
- `Map[U any](func(T) U) *Collection[U]` (separate file, generic function)
- `RemoveAt(index int) bool`
- `GetPagedCollection(pageSize int) []*Collection[T]`
- JSON serialization/deserialization

### 2.2 Pre-Built Type Aliases
**Location**: `coredata/coredynamic/` (new file `CollectionTypes.go`)

```go
// Common collection types — ready to use without specifying T
type StringCollection = Collection[string]
type IntCollection = Collection[int]
type ByteCollection = Collection[byte]
type BoolCollection = Collection[bool]
type Float64Collection = Collection[float64]
type ByteSliceCollection = Collection[[]byte]

// Map collection types
type StringMapCollection = Collection[map[string]string]
type AnyMapCollection = Collection[map[string]any]
type IntMapCollection = Collection[map[string]int]

// Factory shortcuts
func NewStringCollection(cap int) *StringCollection { return NewCollection[string](cap) }
func NewIntCollection(cap int) *IntCollection { return NewCollection[int](cap) }
// ... etc
```

### 2.3 Backward Compatibility
- `DynamicCollection` and `AnyCollection` remain but are **deprecated**
- Add deprecation comments pointing to `Collection[T]`
- Internal methods can delegate to `Collection[any]` where appropriate
- **Effort**: 2-3 sessions

---

## Phase 3: Generics — Payload & Attributes

### 3.1 Generic Deserialize Helpers for Attributes

Add generic helper functions (not methods, because Go doesn't support generic methods on non-generic types):

```go
// coredata/corepayload/generic_helpers.go
package corepayload

import "gitlab.com/auk-go/core/coredata/corejson"

// DeserializePayloadTo deserializes the dynamic payload bytes into T
func DeserializePayloadTo[T any](wrapper *PayloadWrapper) (T, error) {
    var result T
    if wrapper == nil || len(wrapper.Payloads) == 0 {
        return result, defaulterr.NilResult
    }
    err := corejson.Deserialize.UsingBytes(wrapper.Payloads, &result)
    return result, err
}

// DeserializeAttributesPayloadTo deserializes attributes' dynamic payload into T
func DeserializeAttributesPayloadTo[T any](attr *Attributes) (T, error) {
    var result T
    if attr == nil || len(attr.DynamicPayloads) == 0 {
        return result, defaulterr.NilResult
    }
    err := corejson.Deserialize.UsingBytes(attr.DynamicPayloads, &result)
    return result, err
}

// DeserializePayloadToSlice deserializes the payload into []T
func DeserializePayloadToSlice[T any](wrapper *PayloadWrapper) ([]T, error) {
    var result []T
    if wrapper == nil || len(wrapper.Payloads) == 0 {
        return []T{}, defaulterr.NilResult
    }
    err := corejson.Deserialize.UsingBytes(wrapper.Payloads, &result)
    return result, err
}
```

**Usage**:
```go
user, err := corepayload.DeserializePayloadTo[User](wrapper)
users, err := corepayload.DeserializePayloadToSlice[User](wrapper)
config, err := corepayload.DeserializeAttributesPayloadTo[AppConfig](attrs)
```

### 3.2 Why Generic Functions (Not Methods)

Go doesn't allow generic methods on non-generic structs. So we can't do:
```go
// ❌ Not possible in Go
func (it *PayloadWrapper) DeserializeTo[T any]() (T, error)
```

Instead, we use **package-level generic functions** that accept the struct as a parameter. This is the idiomatic Go approach.

### 3.3 Optional: TypedPayloadWrapper[T]

For cases where the payload type is known at compile time:

```go
type TypedPayloadWrapper[T any] struct {
    *PayloadWrapper
}

func NewTypedPayload[T any](wrapper *PayloadWrapper) *TypedPayloadWrapper[T] {
    return &TypedPayloadWrapper[T]{PayloadWrapper: wrapper}
}

func (it *TypedPayloadWrapper[T]) TypedPayload() (T, error) {
    return DeserializePayloadTo[T](it.PayloadWrapper)
}

func (it *TypedPayloadWrapper[T]) TypedPayloadMust() T {
    result, err := it.TypedPayload()
    if err != nil { panic(err) }
    return result
}
```

**Effort**: 2 sessions

---

## Phase 4: Test Coverage Expansion

### Priority Order (by risk/usage):

| Priority | Package | Reason |
|----------|---------|--------|
| P0 | `conditional/` | Generic helpers used everywhere, zero tests |
| P0 | `errcore/` | Error system is critical infrastructure |
| P0 | `converters/` | Type conversions affect data integrity |
| P1 | `coretaskinfo/` | Metadata system, many nil-guard paths |
| P1 | `coredata/corepayload/` | Payload system is the data transport backbone |
| P1 | `regexnew/` | Lazy regex with concurrency — needs validation |
| P2 | `coremath/` | Simple but critical correctness |
| P2 | `coresort/` | Sorting correctness |
| P2 | `coreutils/` | Utility functions |
| P3 | `mutexbykey/` | Concurrency — important but specialized |
| P3 | `namevalue/` | Simple data types |
| P3 | `pagingutil/` | Paging calculations |
| P3 | `typesconv/` | Type conversions |
| P3 | `coreappend/` | Append utilities |
| P3 | `coreunique/` | Uniqueness utilities |

### Per-Package Test Template

Each package test follows the structure in [Testing Guidelines](/spec/01-app/16-testing-guidelines.md):

```
tests/integratedtests/{package}tests/
├── {Feature}_testcases.go
├── {Feature}_test.go
└── (optional) testWrapper.go
```

**Effort**: 1-2 sessions per P0 package, 1 session per P1-P3 package

---

## Phase 5: Refactoring Large Files

| File | Lines | Action |
|------|-------|--------|
| `PayloadWrapper.go` | 842 | Split: `PayloadWrapperGetters.go`, `PayloadWrapperSetters.go`, `PayloadWrapperJson.go` |
| `Attributes.go` | 768 | Split: `AttributesGetters.go`, `AttributesSetters.go`, `AttributesJson.go` |
| `Info.go` | 646 | Split: `InfoGetters.go`, `InfoIncludeExclude.go`, `InfoJson.go` |
| `DynamicCollection.go` | 636 | Deprecate and replace with `Collection[T]` |
| `AnyCollection.go` | 707 | Deprecate and replace with `Collection[any]` |
| `Dynamic.go` | 673 | Split: `DynamicGetters.go`, `DynamicReflect.go`, `DynamicJson.go` |
| `BaseTestCase.go` | 435 | Split: `BaseTestCaseAssert.go`, `BaseTestCaseParams.go` |

**Effort**: 2-3 sessions

---

## Phase 6: Value Receiver Migration

Migrate read-only methods from pointer to value receivers, package by package:
- Start with small packages (`coreversion/`, `issetter/`)
- Graduate to larger packages (`coretaskinfo/`, `corepayload/`)
- Always verify interface satisfaction after changes

**Effort**: Ongoing, 1-2 files per session alongside other work

---

## Summary Timeline

| Phase | Focus | Sessions | Dependencies |
|-------|-------|----------|-------------|
| 1 | Foundation (any, bugs, Go version) | 3-4 | None |
| 2 | Generic Collection[T] | 3-4 | Phase 1 |
| 3 | Generic Payload helpers | 2 | Phase 1 |
| 4 | Test coverage | 8-10 | None (parallel) |
| 5 | File splitting | 2-3 | None (parallel) |
| 6 | Value receivers | Ongoing | Phase 2 |
