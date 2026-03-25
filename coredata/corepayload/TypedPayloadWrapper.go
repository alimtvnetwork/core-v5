package corepayload

import (
	"encoding/json"
	"fmt"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/defaulterr"
)

// TypedPayloadWrapper is a generic version of PayloadWrapper where T represents
// the deserialized type of the Payloads field.
//
// It wraps a standard PayloadWrapper and provides typed access to the payload data
// via the TypedData() method and GetAs* accessors.
//
// Usage:
//
//	type User struct { Name string; Email string }
//
//	// Create from an existing PayloadWrapper
//	typed, err := corepayload.NewTypedPayloadWrapper[User](wrapper)
//	fmt.Println(typed.TypedData().Name)  // strongly typed
//
//	// Create directly
//	typed, err = corepayload.NewTypedPayloadWrapperFrom[User](
//	    "user-create", "usr-123", "User",
//	    User{Name: "Alice", Email: "alice@example.com"},
//	)
type TypedPayloadWrapper[T any] struct {
	Wrapper   *PayloadWrapper
	typedData T
	parsed    bool
}

// NewTypedPayloadWrapper creates a TypedPayloadWrapper by deserializing the
// PayloadWrapper's Payloads bytes into T.
func NewTypedPayloadWrapper[T any](wrapper *PayloadWrapper) (*TypedPayloadWrapper[T], error) {
	if wrapper == nil {
		return nil, defaulterr.NilResult
	}

	var data T
	err := corejson.Deserialize.UsingBytes(wrapper.Payloads, &data)

	if err != nil {
		return nil, err
	}

	return &TypedPayloadWrapper[T]{
		Wrapper:   wrapper,
		typedData: data,
		parsed:    true,
	}, nil
}

// NewTypedPayloadWrapperFrom creates a TypedPayloadWrapper directly from typed data.
//
// The data is serialized into the inner PayloadWrapper's Payloads field.
func NewTypedPayloadWrapperFrom[T any](
	name string,
	identifier string,
	entityType string,
	data T,
) (*TypedPayloadWrapper[T], error) {
	jsonBytes, err := corejson.Serialize.Raw(data)

	if err != nil {
		return nil, err
	}

	wrapper := &PayloadWrapper{
		Name:       name,
		Identifier: identifier,
		EntityType: entityType,
		Payloads:   jsonBytes,
	}

	return &TypedPayloadWrapper[T]{
		Wrapper:   wrapper,
		typedData: data,
		parsed:    true,
	}, nil
}

// NewTypedPayloadWrapperFromInstruction creates a TypedPayloadWrapper from a
// PayloadCreateInstruction, automatically serializing the typed data.
func NewTypedPayloadWrapperFromInstruction[T any](
	name string,
	identifier string,
	taskTypeName string,
	entityType string,
	categoryName string,
	hasManyRecords bool,
	data T,
	attributes *Attributes,
) (*TypedPayloadWrapper[T], error) {
	jsonBytes, err := corejson.Serialize.Raw(data)

	if err != nil {
		return nil, err
	}

	wrapper := &PayloadWrapper{
		Name:           name,
		Identifier:     identifier,
		TaskTypeName:   taskTypeName,
		EntityType:     entityType,
		CategoryName:   categoryName,
		HasManyRecords: hasManyRecords,
		Payloads:       jsonBytes,
		Attributes:     attributes,
	}

	return &TypedPayloadWrapper[T]{
		Wrapper:   wrapper,
		typedData: data,
		parsed:    true,
	}, nil
}

// NewTypedPayloadWrapperMust creates a TypedPayloadWrapper or panics on error.
func NewTypedPayloadWrapperMust[T any](wrapper *PayloadWrapper) *TypedPayloadWrapper[T] {
	result, err := NewTypedPayloadWrapper[T](wrapper)

	if err != nil {
		panic(err)
	}

	return result
}

// =============================================================================
// Core accessors
// =============================================================================

// TypedData returns the deserialized, strongly-typed payload data.
func (it *TypedPayloadWrapper[T]) TypedData() T {
	return it.typedData
}

// Data is an alias for TypedData.
func (it *TypedPayloadWrapper[T]) Data() T {
	return it.typedData
}

// IsParsed returns whether the typed data has been successfully parsed.
func (it *TypedPayloadWrapper[T]) IsParsed() bool {
	return it != nil && it.parsed
}

// Name returns the payload name from the underlying wrapper.
func (it *TypedPayloadWrapper[T]) Name() string {
	if it == nil || it.Wrapper == nil {
		return ""
	}

	return it.Wrapper.Name
}

// Identifier returns the identifier from the underlying wrapper.
func (it *TypedPayloadWrapper[T]) Identifier() string {
	if it == nil || it.Wrapper == nil {
		return ""
	}

	return it.Wrapper.Identifier
}

// IdString returns the identifier (alias for Identifier).
func (it *TypedPayloadWrapper[T]) IdString() string {
	return it.Identifier()
}

// IdInteger returns the identifier parsed as integer.
func (it *TypedPayloadWrapper[T]) IdInteger() int {
	if it == nil || it.Wrapper == nil {
		return constants.InvalidValue
	}

	return it.Wrapper.IdInteger()
}

// EntityType returns the entity type from the underlying wrapper.
func (it *TypedPayloadWrapper[T]) EntityType() string {
	if it == nil || it.Wrapper == nil {
		return ""
	}

	return it.Wrapper.EntityType
}

// CategoryName returns the category name from the underlying wrapper.
func (it *TypedPayloadWrapper[T]) CategoryName() string {
	if it == nil || it.Wrapper == nil {
		return ""
	}

	return it.Wrapper.CategoryName
}

// TaskTypeName returns the task type name from the underlying wrapper.
func (it *TypedPayloadWrapper[T]) TaskTypeName() string {
	if it == nil || it.Wrapper == nil {
		return ""
	}

	return it.Wrapper.TaskTypeName
}

// HasManyRecords returns whether the wrapper signals multiple records.
func (it *TypedPayloadWrapper[T]) HasManyRecords() bool {
	if it == nil || it.Wrapper == nil {
		return false
	}

	return it.Wrapper.HasManyRecords
}

// HasSingleRecord returns whether the wrapper signals a single record.
func (it *TypedPayloadWrapper[T]) HasSingleRecord() bool {
	return !it.HasManyRecords()
}

// Attributes returns the Attributes from the underlying wrapper.
func (it *TypedPayloadWrapper[T]) Attributes() *Attributes {
	if it == nil || it.Wrapper == nil {
		return nil
	}

	return it.Wrapper.Attributes
}

// InitializeAttributesOnNull initializes Attributes on the underlying wrapper if nil.
func (it *TypedPayloadWrapper[T]) InitializeAttributesOnNull() *Attributes {
	if it == nil || it.Wrapper == nil {
		return nil
	}

	it.Wrapper.InitializeAttributesOnNull()

	return it.Wrapper.Attributes
}

// =============================================================================
// Error handling
// =============================================================================

// HasError returns whether the underlying wrapper has an error.
func (it *TypedPayloadWrapper[T]) HasError() bool {
	if it == nil || it.Wrapper == nil {
		return false
	}

	return it.Wrapper.HasError()
}

// IsEmpty returns whether the underlying wrapper is empty.
func (it *TypedPayloadWrapper[T]) IsEmpty() bool {
	if it == nil || it.Wrapper == nil {
		return true
	}

	return it.Wrapper.IsEmpty()
}

// HasItems returns whether the wrapper has payload items.
func (it *TypedPayloadWrapper[T]) HasItems() bool {
	return !it.IsEmpty()
}

// HasSafeItems returns true if non-empty and no error.
func (it *TypedPayloadWrapper[T]) HasSafeItems() bool {
	return it.HasItems() && !it.HasError()
}

// Error returns the error from the underlying wrapper.
func (it *TypedPayloadWrapper[T]) Error() error {
	if it == nil || it.Wrapper == nil {
		return nil
	}

	return it.Wrapper.Error()
}

// HandleError panics if the underlying wrapper has an error.
func (it *TypedPayloadWrapper[T]) HandleError() {
	if it.HasError() {
		it.Wrapper.HandleError()
	}
}

// =============================================================================
// String representations
// =============================================================================

// String returns a string representation.
func (it *TypedPayloadWrapper[T]) String() string {
	if it == nil || it.Wrapper == nil {
		return ""
	}

	return it.Wrapper.String()
}

// PrettyJsonString returns a pretty-printed JSON string.
func (it *TypedPayloadWrapper[T]) PrettyJsonString() string {
	if it == nil || it.Wrapper == nil {
		return ""
	}

	return it.Wrapper.PrettyJsonString()
}

// JsonString returns a compact JSON string.
func (it *TypedPayloadWrapper[T]) JsonString() string {
	if it == nil || it.Wrapper == nil {
		return ""
	}

	return it.Wrapper.JsonString()
}

// =============================================================================
// JSON operations
// =============================================================================

// Json returns a corejson.Result for the underlying wrapper.
func (it *TypedPayloadWrapper[T]) Json() corejson.Result {
	if it == nil || it.Wrapper == nil {
		return corejson.Result{}
	}

	return it.Wrapper.Json()
}

// JsonPtr returns a pointer to a corejson.Result.
func (it *TypedPayloadWrapper[T]) JsonPtr() *corejson.Result {
	if it == nil || it.Wrapper == nil {
		return corejson.Empty.ResultPtr()
	}

	return it.Wrapper.JsonPtr()
}

// MarshalJSON implements the json.Marshaler interface, delegating to the wrapper.
func (it *TypedPayloadWrapper[T]) MarshalJSON() ([]byte, error) {
	if it == nil || it.Wrapper == nil {
		return nil, defaulterr.NilResult
	}

	return it.Wrapper.MarshalJSON()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// It unmarshals into the wrapper, then re-parses typed data.
func (it *TypedPayloadWrapper[T]) UnmarshalJSON(data []byte) error {
	if it == nil {
		return defaulterr.NilResult
	}

	if it.Wrapper == nil {
		it.Wrapper = &PayloadWrapper{}
	}

	err := json.Unmarshal(data, it.Wrapper)

	if err != nil {
		it.parsed = false

		return err
	}

	err = corejson.Deserialize.UsingBytes(it.Wrapper.Payloads, &it.typedData)
	it.parsed = err == nil

	return err
}

// Serialize serializes the entire wrapper to JSON bytes.
func (it *TypedPayloadWrapper[T]) Serialize() ([]byte, error) {
	if it == nil || it.Wrapper == nil {
		return nil, defaulterr.NilResult
	}

	return it.Wrapper.Serialize()
}

// SerializeMust serializes to JSON bytes, panics on error.
func (it *TypedPayloadWrapper[T]) SerializeMust() []byte {
	if it == nil || it.Wrapper == nil {
		panic(defaulterr.NilResult)
	}

	return it.Wrapper.SerializeMust()
}

// TypedDataJson returns a corejson.Result for the typed data only (not the wrapper).
func (it *TypedPayloadWrapper[T]) TypedDataJson() corejson.Result {
	return corejson.New(it.typedData)
}

// TypedDataJsonPtr returns a pointer to a corejson.Result for the typed data.
func (it *TypedPayloadWrapper[T]) TypedDataJsonPtr() *corejson.Result {
	return corejson.NewPtr(it.typedData)
}

// TypedDataJsonBytes serializes the typed data to JSON bytes.
func (it *TypedPayloadWrapper[T]) TypedDataJsonBytes() ([]byte, error) {
	return json.Marshal(it.typedData)
}

// =============================================================================
// GetAs* type assertion methods
// =============================================================================

// GetAsString attempts to retrieve the typed data as a string.
func (it *TypedPayloadWrapper[T]) GetAsString() (string, bool) {
	val, ok := any(it.typedData).(string)

	return val, ok
}

// GetAsInt attempts to retrieve the typed data as an int.
func (it *TypedPayloadWrapper[T]) GetAsInt() (int, bool) {
	val, ok := any(it.typedData).(int)

	return val, ok
}

// GetAsInt64 attempts to retrieve the typed data as an int64.
func (it *TypedPayloadWrapper[T]) GetAsInt64() (int64, bool) {
	val, ok := any(it.typedData).(int64)

	return val, ok
}

// GetAsFloat64 attempts to retrieve the typed data as a float64.
func (it *TypedPayloadWrapper[T]) GetAsFloat64() (float64, bool) {
	val, ok := any(it.typedData).(float64)

	return val, ok
}

// GetAsFloat32 attempts to retrieve the typed data as a float32.
func (it *TypedPayloadWrapper[T]) GetAsFloat32() (float32, bool) {
	val, ok := any(it.typedData).(float32)

	return val, ok
}

// GetAsBool attempts to retrieve the typed data as a bool.
func (it *TypedPayloadWrapper[T]) GetAsBool() (bool, bool) {
	val, ok := any(it.typedData).(bool)

	return val, ok
}

// GetAsBytes attempts to retrieve the typed data as []byte.
func (it *TypedPayloadWrapper[T]) GetAsBytes() ([]byte, bool) {
	val, ok := any(it.typedData).([]byte)

	return val, ok
}

// GetAsStrings attempts to retrieve the typed data as []string.
func (it *TypedPayloadWrapper[T]) GetAsStrings() ([]string, bool) {
	val, ok := any(it.typedData).([]string)

	return val, ok
}

// =============================================================================
// Value* convenience methods
// =============================================================================

// ValueString returns the typed data as a string, with fmt.Sprintf fallback.
func (it *TypedPayloadWrapper[T]) ValueString() string {
	str, isString := any(it.typedData).(string)

	if isString {
		return str
	}

	return fmt.Sprintf("%v", it.typedData)
}

// ValueInt returns the typed data as int, or constants.InvalidValue on failure.
func (it *TypedPayloadWrapper[T]) ValueInt() int {
	val, ok := any(it.typedData).(int)

	if ok {
		return val
	}

	return constants.InvalidValue
}

// ValueBool returns the typed data as bool, or false on failure.
func (it *TypedPayloadWrapper[T]) ValueBool() bool {
	val, ok := any(it.typedData).(bool)

	if ok {
		return val
	}

	return false
}

// =============================================================================
// Setters (mutate wrapper metadata)
// =============================================================================

// SetName sets the Name on the underlying wrapper.
func (it *TypedPayloadWrapper[T]) SetName(name string) {
	if it != nil && it.Wrapper != nil {
		it.Wrapper.Name = name
	}
}

// SetIdentifier sets the Identifier on the underlying wrapper.
func (it *TypedPayloadWrapper[T]) SetIdentifier(identifier string) {
	if it != nil && it.Wrapper != nil {
		it.Wrapper.Identifier = identifier
	}
}

// SetEntityType sets the EntityType on the underlying wrapper.
func (it *TypedPayloadWrapper[T]) SetEntityType(entityType string) {
	if it != nil && it.Wrapper != nil {
		it.Wrapper.EntityType = entityType
	}
}

// SetCategoryName sets the CategoryName on the underlying wrapper.
func (it *TypedPayloadWrapper[T]) SetCategoryName(categoryName string) {
	if it != nil && it.Wrapper != nil {
		it.Wrapper.CategoryName = categoryName
	}
}

// SetTypedData replaces the typed data and re-serializes into the wrapper's Payloads.
func (it *TypedPayloadWrapper[T]) SetTypedData(data T) error {
	if it == nil || it.Wrapper == nil {
		return defaulterr.NilResult
	}

	jsonBytes, err := corejson.Serialize.Raw(data)

	if err != nil {
		return err
	}

	it.typedData = data
	it.Wrapper.Payloads = jsonBytes
	it.parsed = true

	return nil
}

// SetTypedDataMust replaces the typed data, panics on serialization error.
func (it *TypedPayloadWrapper[T]) SetTypedDataMust(data T) {
	err := it.SetTypedData(data)

	if err != nil {
		panic(err)
	}
}

// =============================================================================
// Clone and conversion
// =============================================================================

// ClonePtr returns a deep clone of the TypedPayloadWrapper.
func (it *TypedPayloadWrapper[T]) ClonePtr(isDeepClone bool) (*TypedPayloadWrapper[T], error) {
	if it == nil {
		return nil, nil
	}

	clonedWrapper, err := it.Wrapper.ClonePtr(isDeepClone)

	if err != nil {
		return nil, err
	}

	// Re-parse typed data from cloned wrapper
	var data T
	err = corejson.Deserialize.UsingBytes(clonedWrapper.Payloads, &data)

	if err != nil {
		return nil, err
	}

	return &TypedPayloadWrapper[T]{
		Wrapper:   clonedWrapper,
		typedData: data,
		parsed:    true,
	}, nil
}

// Clone returns a value clone.
func (it *TypedPayloadWrapper[T]) Clone(isDeepClone bool) (TypedPayloadWrapper[T], error) {
	clonedPtr, err := it.ClonePtr(isDeepClone)

	if err != nil {
		return TypedPayloadWrapper[T]{}, err
	}

	if clonedPtr == nil {
		return TypedPayloadWrapper[T]{}, defaulterr.NilResult
	}

	return *clonedPtr, nil
}

// ToPayloadWrapper returns the non-generic PayloadWrapper for backward compatibility.
func (it *TypedPayloadWrapper[T]) ToPayloadWrapper() *PayloadWrapper {
	if it == nil {
		return nil
	}

	return it.Wrapper
}

// PayloadWrapperValue returns the underlying PayloadWrapper (alias for ToPayloadWrapper).
func (it *TypedPayloadWrapper[T]) PayloadWrapperValue() *PayloadWrapper {
	return it.ToPayloadWrapper()
}

// Reparse re-deserializes the wrapper's Payloads into the typed data.
// Use after manually modifying the wrapper's Payloads bytes.
func (it *TypedPayloadWrapper[T]) Reparse() error {
	if it == nil || it.Wrapper == nil {
		return defaulterr.NilResult
	}

	var data T
	err := corejson.Deserialize.UsingBytes(it.Wrapper.Payloads, &data)

	if err != nil {
		it.parsed = false

		return err
	}

	it.typedData = data
	it.parsed = true

	return nil
}

// DynamicPayloads returns the raw payload bytes from the underlying wrapper.
func (it *TypedPayloadWrapper[T]) DynamicPayloads() []byte {
	if it == nil || it.Wrapper == nil {
		return []byte{}
	}

	return it.Wrapper.Payloads
}

// PayloadsString returns the raw payload bytes as a string.
func (it *TypedPayloadWrapper[T]) PayloadsString() string {
	if it == nil || it.Wrapper == nil {
		return ""
	}

	return it.Wrapper.PayloadsString()
}

// Length returns the byte length of the payloads.
func (it *TypedPayloadWrapper[T]) Length() int {
	if it == nil || it.Wrapper == nil {
		return 0
	}

	return it.Wrapper.Length()
}

// IsNull returns whether the typed wrapper or its inner wrapper is nil.
func (it *TypedPayloadWrapper[T]) IsNull() bool {
	return it == nil || it.Wrapper == nil
}

// Clear clears the underlying wrapper and resets parsed state.
func (it *TypedPayloadWrapper[T]) Clear() {
	if it == nil {
		return
	}

	if it.Wrapper != nil {
		it.Wrapper.Clear()
	}

	var zero T
	it.typedData = zero
	it.parsed = false
}

// Dispose clears and nils the underlying wrapper.
func (it *TypedPayloadWrapper[T]) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
	it.Wrapper = nil
}
