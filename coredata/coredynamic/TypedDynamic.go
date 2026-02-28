package coredynamic

import (
	"encoding/json"
	"fmt"

	"gitlab.com/auk-go/core/coredata/corejson"
)

// TypedDynamic is a generic, strongly-typed wrapper around a value of type T.
//
// Unlike Dynamic (which wraps interface{}), TypedDynamic[T] provides
// compile-time type safety and eliminates the need for type assertions.
//
// Usage:
//
//	d := coredynamic.NewTypedDynamic[string]("hello", true)
//	fmt.Println(d.Data())    // "hello" (typed as string)
//	fmt.Println(d.IsValid()) // true
type TypedDynamic[T any] struct {
	innerData T
	isValid   bool
}

// NewTypedDynamic creates a valid TypedDynamic with the given data and validity flag.
func NewTypedDynamic[T any](data T, isValid bool) TypedDynamic[T] {
	return TypedDynamic[T]{
		innerData: data,
		isValid:   isValid,
	}
}

// NewTypedDynamicValid creates a valid TypedDynamic.
func NewTypedDynamicValid[T any](data T) TypedDynamic[T] {
	return TypedDynamic[T]{
		innerData: data,
		isValid:   true,
	}
}

// NewTypedDynamicPtr creates a pointer to a TypedDynamic.
func NewTypedDynamicPtr[T any](data T, isValid bool) *TypedDynamic[T] {
	d := NewTypedDynamic(data, isValid)

	return &d
}

// InvalidTypedDynamic creates an invalid TypedDynamic with a zero-value T.
func InvalidTypedDynamic[T any]() TypedDynamic[T] {
	return TypedDynamic[T]{
		isValid: false,
	}
}

// Data returns the underlying typed data.
func (it TypedDynamic[T]) Data() T {
	return it.innerData
}

// IsValid returns whether the dynamic value is valid.
func (it TypedDynamic[T]) IsValid() bool {
	return it.isValid
}

// IsInvalid returns whether the dynamic value is invalid.
func (it TypedDynamic[T]) IsInvalid() bool {
	return !it.isValid
}

// String returns a string representation of the inner data.
func (it TypedDynamic[T]) String() string {
	return fmt.Sprintf("%v", it.innerData)
}

// JsonBytes serializes the inner data to JSON bytes.
func (it TypedDynamic[T]) JsonBytes() ([]byte, error) {
	return json.Marshal(it.innerData)
}

// JsonResult returns a corejson.Result for the inner data.
func (it TypedDynamic[T]) JsonResult() corejson.Result {
	return corejson.New(it.innerData)
}

// Clone returns a copy of the TypedDynamic.
//
// Note: T is copied by value. For pointer types, only the pointer is copied.
func (it TypedDynamic[T]) Clone() TypedDynamic[T] {
	return TypedDynamic[T]{
		innerData: it.innerData,
		isValid:   it.isValid,
	}
}

// ToDynamic converts to the non-generic Dynamic for backward compatibility.
func (it TypedDynamic[T]) ToDynamic() Dynamic {
	return NewDynamic(it.innerData, it.isValid)
}
