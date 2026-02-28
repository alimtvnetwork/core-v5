package coredynamic

import (
	"encoding/json"
	"fmt"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/defaulterr"
)

// TypedDynamic is a generic, strongly-typed wrapper around a value of type T.
//
// Unlike Dynamic (which wraps any), TypedDynamic[T] provides
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

// InvalidTypedDynamicPtr creates an invalid TypedDynamic pointer.
func InvalidTypedDynamicPtr[T any]() *TypedDynamic[T] {
	d := InvalidTypedDynamic[T]()

	return &d
}

// Data returns the underlying typed data.
func (it TypedDynamic[T]) Data() T {
	return it.innerData
}

// Value is an alias for Data, mirroring Dynamic.Value().
func (it TypedDynamic[T]) Value() T {
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

// =============================================================================
// JSON operations
// =============================================================================

// JsonBytes serializes the inner data to JSON bytes.
func (it TypedDynamic[T]) JsonBytes() ([]byte, error) {
	return json.Marshal(it.innerData)
}

// JsonResult returns a corejson.Result for the inner data.
func (it TypedDynamic[T]) JsonResult() corejson.Result {
	return corejson.New(it.innerData)
}

// Json returns a corejson.Result (alias matching Dynamic.Json()).
func (it TypedDynamic[T]) Json() corejson.Result {
	return corejson.New(it.innerData)
}

// JsonPtr returns a pointer to a corejson.Result.
func (it TypedDynamic[T]) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it.innerData)
}

// JsonString serializes the inner data to a JSON string.
func (it TypedDynamic[T]) JsonString() (string, error) {
	jsonBytes, err := json.Marshal(it.innerData)

	if err != nil {
		return constants.EmptyString, err
	}

	return string(jsonBytes), nil
}

// MarshalJSON implements the json.Marshaler interface.
func (it TypedDynamic[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.innerData)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (it *TypedDynamic[T]) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &it.innerData)

	it.isValid = err == nil

	return err
}

// ValueMarshal serializes the inner data to JSON bytes.
func (it TypedDynamic[T]) ValueMarshal() ([]byte, error) {
	return corejson.Serialize.ToBytesErr(it.innerData)
}

// Bytes attempts to retrieve the inner data as raw bytes.
// If the data is []byte, it returns directly. Otherwise, it marshals to JSON.
func (it TypedDynamic[T]) Bytes() ([]byte, bool) {
	rawBytes, isBytes := any(it.innerData).([]byte)

	if isBytes {
		return rawBytes, true
	}

	marshalled, err := json.Marshal(it.innerData)

	return marshalled, err == nil
}

// =============================================================================
// GetAs* type assertion methods
// =============================================================================

// GetAsString attempts to retrieve the data as a string.
func (it TypedDynamic[T]) GetAsString() (string, bool) {
	val, ok := any(it.innerData).(string)

	return val, ok
}

// GetAsInt attempts to retrieve the data as an int.
func (it TypedDynamic[T]) GetAsInt() (int, bool) {
	val, ok := any(it.innerData).(int)

	return val, ok
}

// GetAsInt64 attempts to retrieve the data as an int64.
func (it TypedDynamic[T]) GetAsInt64() (int64, bool) {
	val, ok := any(it.innerData).(int64)

	return val, ok
}

// GetAsUint attempts to retrieve the data as a uint.
func (it TypedDynamic[T]) GetAsUint() (uint, bool) {
	val, ok := any(it.innerData).(uint)

	return val, ok
}

// GetAsFloat64 attempts to retrieve the data as a float64.
func (it TypedDynamic[T]) GetAsFloat64() (float64, bool) {
	val, ok := any(it.innerData).(float64)

	return val, ok
}

// GetAsFloat32 attempts to retrieve the data as a float32.
func (it TypedDynamic[T]) GetAsFloat32() (float32, bool) {
	val, ok := any(it.innerData).(float32)

	return val, ok
}

// GetAsBool attempts to retrieve the data as a bool.
func (it TypedDynamic[T]) GetAsBool() (bool, bool) {
	val, ok := any(it.innerData).(bool)

	return val, ok
}

// GetAsBytes attempts to retrieve the data as []byte.
func (it TypedDynamic[T]) GetAsBytes() ([]byte, bool) {
	val, ok := any(it.innerData).([]byte)

	return val, ok
}

// GetAsStrings attempts to retrieve the data as []string.
func (it TypedDynamic[T]) GetAsStrings() ([]string, bool) {
	val, ok := any(it.innerData).([]string)

	return val, ok
}

// =============================================================================
// Value* convenience methods (mirroring Dynamic.Value*)
// =============================================================================

// ValueString returns the data as a string, or fmt.Sprintf fallback.
func (it TypedDynamic[T]) ValueString() string {
	str, isString := any(it.innerData).(string)

	if isString {
		return str
	}

	return fmt.Sprintf("%v", it.innerData)
}

// ValueInt returns the data as int, or constants.InvalidValue on failure.
func (it TypedDynamic[T]) ValueInt() int {
	val, ok := any(it.innerData).(int)

	if ok {
		return val
	}

	return constants.InvalidValue
}

// ValueInt64 returns the data as int64, or constants.InvalidValue on failure.
func (it TypedDynamic[T]) ValueInt64() int64 {
	val, ok := any(it.innerData).(int64)

	if ok {
		return val
	}

	return constants.InvalidValue
}

// ValueBool returns the data as bool, or false on failure.
func (it TypedDynamic[T]) ValueBool() bool {
	val, ok := any(it.innerData).(bool)

	if ok {
		return val
	}

	return false
}

// =============================================================================
// Clone and conversion
// =============================================================================

// Clone returns a copy of the TypedDynamic.
//
// Note: T is copied by value. For pointer types, only the pointer is copied.
func (it TypedDynamic[T]) Clone() TypedDynamic[T] {
	return TypedDynamic[T]{
		innerData: it.innerData,
		isValid:   it.isValid,
	}
}

// ClonePtr returns a pointer to a copy of the TypedDynamic.
func (it *TypedDynamic[T]) ClonePtr() *TypedDynamic[T] {
	if it == nil {
		return nil
	}

	cloned := it.Clone()

	return &cloned
}

// NonPtr returns a value copy (mirrors Dynamic.NonPtr()).
func (it TypedDynamic[T]) NonPtr() TypedDynamic[T] {
	return it
}

// Ptr returns a pointer to the TypedDynamic (mirrors Dynamic.Ptr()).
func (it *TypedDynamic[T]) Ptr() *TypedDynamic[T] {
	return it
}

// ToDynamic converts to the non-generic Dynamic for backward compatibility.
func (it TypedDynamic[T]) ToDynamic() Dynamic {
	return NewDynamic(it.innerData, it.isValid)
}

// Deserialize unmarshals JSON bytes into the inner data.
func (it *TypedDynamic[T]) Deserialize(jsonBytes []byte) error {
	if it == nil {
		return defaulterr.UnmarshallingFailedDueToNilOrEmpty
	}

	err := json.Unmarshal(jsonBytes, &it.innerData)

	it.isValid = err == nil

	return err
}

// JsonModel returns the inner data for JSON serialization (mirrors Dynamic.JsonModel()).
func (it TypedDynamic[T]) JsonModel() T {
	return it.innerData
}

// JsonModelAny returns the inner data as any (mirrors Dynamic.JsonModelAny()).
func (it TypedDynamic[T]) JsonModelAny() any {
	return it.innerData
}
