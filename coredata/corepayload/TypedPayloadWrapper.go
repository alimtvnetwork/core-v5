package corepayload

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/defaulterr"
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
//	typed = corepayload.NewTypedPayloadWrapperFrom[User](
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

// TypedData returns the deserialized, strongly-typed payload data.
func (it *TypedPayloadWrapper[T]) TypedData() T {
	return it.typedData
}

// PayloadWrapper returns the underlying PayloadWrapper.
func (it *TypedPayloadWrapper[T]) PayloadWrapper() *PayloadWrapper {
	if it == nil {
		return nil
	}

	return it.Wrapper
}

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

// GetAsFloat64 attempts to retrieve the typed data as a float64.
func (it *TypedPayloadWrapper[T]) GetAsFloat64() (float64, bool) {
	val, ok := any(it.typedData).(float64)

	return val, ok
}

// GetAsBool attempts to retrieve the typed data as a bool.
func (it *TypedPayloadWrapper[T]) GetAsBool() (bool, bool) {
	val, ok := any(it.typedData).(bool)

	return val, ok
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

// EntityType returns the entity type from the underlying wrapper.
func (it *TypedPayloadWrapper[T]) EntityType() string {
	if it == nil || it.Wrapper == nil {
		return ""
	}

	return it.Wrapper.EntityType
}

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

// ToPayloadWrapper returns the non-generic PayloadWrapper for backward compatibility.
func (it *TypedPayloadWrapper[T]) ToPayloadWrapper() *PayloadWrapper {
	if it == nil {
		return nil
	}

	return it.Wrapper
}
