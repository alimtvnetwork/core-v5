package corepayload

import (
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

// newTypedPayloadWrapperCreator provides typed factory methods for TypedPayloadWrapper[T].
//
// Because Go does not allow generic methods on non-generic types,
// these are package-level generic functions rather than methods on a creator struct.
// They mirror the newPayloadWrapperCreator API with compile-time type safety.

// TypedPayloadWrapperFrom creates a TypedPayloadWrapper[T] from typed data
// with name, identifier, and entity type.
//
// Usage:
//
//	typed, err := corepayload.TypedPayloadWrapperFrom[User](
//	    "user-create", "usr-123", "User",
//	    User{Name: "Alice"},
//	)
func TypedPayloadWrapperFrom[T any](
	name string,
	identifier string,
	entityType string,
	data T,
) (*TypedPayloadWrapper[T], error) {
	return NewTypedPayloadWrapperFrom[T](name, identifier, entityType, data)
}

// TypedPayloadWrapperRecord creates a TypedPayloadWrapper[T] with auto-detected entity type.
//
// Usage:
//
//	typed, err := corepayload.TypedPayloadWrapperRecord[User](
//	    "user-create", "usr-123", "task", "category",
//	    User{Name: "Alice"},
//	)
func TypedPayloadWrapperRecord[T any](
	name string,
	identifier string,
	taskTypeName string,
	categoryName string,
	data T,
) (*TypedPayloadWrapper[T], error) {
	entityType := reflectinternal.ReflectType.SafeName(data)

	return NewTypedPayloadWrapperFromInstruction[T](
		name,
		identifier,
		taskTypeName,
		entityType,
		categoryName,
		false,
		data,
		nil,
	)
}

// TypedPayloadWrapperRecords creates a TypedPayloadWrapper[T] for multiple records.
//
// Usage:
//
//	typed, err := corepayload.TypedPayloadWrapperRecords[[]User](
//	    "users-list", "batch-1", "task", "category",
//	    users,
//	)
func TypedPayloadWrapperRecords[T any](
	name string,
	identifier string,
	taskTypeName string,
	categoryName string,
	data T,
) (*TypedPayloadWrapper[T], error) {
	entityType := reflectinternal.ReflectType.SafeTypeNameOfSliceOrSingle(false, data)

	return NewTypedPayloadWrapperFromInstruction[T](
		name,
		identifier,
		taskTypeName,
		entityType,
		categoryName,
		true,
		data,
		nil,
	)
}

// TypedPayloadWrapperNameIdRecord creates a TypedPayloadWrapper[T] with name, ID, and data.
//
// Usage:
//
//	typed, err := corepayload.TypedPayloadWrapperNameIdRecord[User](
//	    "user-create", "usr-123",
//	    User{Name: "Alice"},
//	)
func TypedPayloadWrapperNameIdRecord[T any](
	name string,
	identifier string,
	data T,
) (*TypedPayloadWrapper[T], error) {
	entityType := reflectinternal.ReflectType.SafeName(data)

	return NewTypedPayloadWrapperFromInstruction[T](
		name,
		identifier,
		entityType,
		entityType,
		"",
		false,
		data,
		nil,
	)
}

// TypedPayloadWrapperNameIdCategory creates a TypedPayloadWrapper[T] with name, ID, category, and data.
func TypedPayloadWrapperNameIdCategory[T any](
	name string,
	identifier string,
	categoryName string,
	data T,
) (*TypedPayloadWrapper[T], error) {
	entityType := reflectinternal.ReflectType.SafeName(data)

	return NewTypedPayloadWrapperFromInstruction[T](
		name,
		identifier,
		entityType,
		entityType,
		categoryName,
		false,
		data,
		nil,
	)
}

// TypedPayloadWrapperAll creates a TypedPayloadWrapper[T] with all fields specified.
func TypedPayloadWrapperAll[T any](
	name string,
	identifier string,
	taskTypeName string,
	entityTypeName string,
	categoryName string,
	hasManyRecords bool,
	data T,
	attributes *Attributes,
) (*TypedPayloadWrapper[T], error) {
	return NewTypedPayloadWrapperFromInstruction[T](
		name,
		identifier,
		taskTypeName,
		entityTypeName,
		categoryName,
		hasManyRecords,
		data,
		attributes,
	)
}

// TypedPayloadWrapperDeserialize deserializes raw JSON bytes into a TypedPayloadWrapper[T].
//
// First deserializes to PayloadWrapper, then parses Payloads into T.
func TypedPayloadWrapperDeserialize[T any](rawBytes []byte) (*TypedPayloadWrapper[T], error) {
	wrapper, err := New.PayloadWrapper.Deserialize(rawBytes)

	if err != nil {
		return nil, err
	}

	return NewTypedPayloadWrapper[T](wrapper)
}

// TypedPayloadWrapperDeserializeUsingJsonResult deserializes a corejson.Result into TypedPayloadWrapper[T].
func TypedPayloadWrapperDeserializeUsingJsonResult[T any](
	jsonResult *corejson.Result,
) (*TypedPayloadWrapper[T], error) {
	wrapper, err := New.PayloadWrapper.DeserializeUsingJsonResult(jsonResult)

	if err != nil {
		return nil, err
	}

	return NewTypedPayloadWrapper[T](wrapper)
}

// TypedPayloadWrapperDeserializeToMany deserializes raw JSON bytes into a slice
// of TypedPayloadWrapper[T].
func TypedPayloadWrapperDeserializeToMany[T any](
	rawBytes []byte,
) ([]*TypedPayloadWrapper[T], error) {
	wrappers, err := New.PayloadWrapper.DeserializeToMany(rawBytes)

	if err != nil {
		return nil, err
	}

	typedWrappers := make([]*TypedPayloadWrapper[T], 0, len(wrappers))

	for _, wrapper := range wrappers {
		typed, typedErr := NewTypedPayloadWrapper[T](wrapper)

		if typedErr != nil {
			return nil, typedErr
		}

		typedWrappers = append(typedWrappers, typed)
	}

	return typedWrappers, nil
}
