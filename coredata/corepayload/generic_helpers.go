package corepayload

import (
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/defaulterr"
)

// DeserializePayloadTo deserializes the PayloadWrapper's Payloads bytes into T.
//
// Usage:
//
//	user, err := corepayload.DeserializePayloadTo[User](wrapper)
func DeserializePayloadTo[T any](wrapper *PayloadWrapper) (T, error) {
	var result T
	if wrapper == nil || len(wrapper.Payloads) == 0 {
		return result, defaulterr.NilResult
	}

	err := corejson.Deserialize.UsingBytes(wrapper.Payloads, &result)
	return result, err
}

// DeserializePayloadToMust deserializes the PayloadWrapper's Payloads bytes into T or panics.
//
// Usage:
//
//	user := corepayload.DeserializePayloadToMust[User](wrapper)
func DeserializePayloadToMust[T any](wrapper *PayloadWrapper) T {
	result, err := DeserializePayloadTo[T](wrapper)
	if err != nil {
		panic(err)
	}
	return result
}

// DeserializePayloadToSlice deserializes the PayloadWrapper's Payloads bytes into []T.
//
// Usage:
//
//	users, err := corepayload.DeserializePayloadToSlice[User](wrapper)
func DeserializePayloadToSlice[T any](wrapper *PayloadWrapper) ([]T, error) {
	var result []T
	if wrapper == nil || len(wrapper.Payloads) == 0 {
		return []T{}, defaulterr.NilResult
	}

	err := corejson.Deserialize.UsingBytes(wrapper.Payloads, &result)
	return result, err
}

// DeserializePayloadToSliceMust deserializes the PayloadWrapper's Payloads into []T or panics.
func DeserializePayloadToSliceMust[T any](wrapper *PayloadWrapper) []T {
	result, err := DeserializePayloadToSlice[T](wrapper)
	if err != nil {
		panic(err)
	}
	return result
}

// DeserializeAttributesPayloadTo deserializes the Attributes' DynamicPayloads bytes into T.
//
// Usage:
//
//	config, err := corepayload.DeserializeAttributesPayloadTo[AppConfig](attrs)
func DeserializeAttributesPayloadTo[T any](attr *Attributes) (T, error) {
	var result T
	if attr == nil || len(attr.DynamicPayloads) == 0 {
		return result, defaulterr.NilResult
	}

	err := corejson.Deserialize.UsingBytes(attr.DynamicPayloads, &result)
	return result, err
}

// DeserializeAttributesPayloadToMust deserializes Attributes' DynamicPayloads into T or panics.
func DeserializeAttributesPayloadToMust[T any](attr *Attributes) T {
	result, err := DeserializeAttributesPayloadTo[T](attr)
	if err != nil {
		panic(err)
	}
	return result
}

// DeserializeAttributesPayloadToSlice deserializes the Attributes' DynamicPayloads into []T.
func DeserializeAttributesPayloadToSlice[T any](attr *Attributes) ([]T, error) {
	var result []T
	if attr == nil || len(attr.DynamicPayloads) == 0 {
		return []T{}, defaulterr.NilResult
	}

	err := corejson.Deserialize.UsingBytes(attr.DynamicPayloads, &result)
	return result, err
}
