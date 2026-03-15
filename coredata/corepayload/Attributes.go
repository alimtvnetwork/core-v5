package corepayload

import (
	"bytes"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coreinstruction"
	"github.com/alimtvnetwork/core/coreinterface/errcoreinf"
)

// Attributes holds metadata, key-value pairs, error state, and dynamic payloads
// associated with a PayloadWrapper.
//
// Getters are in AttributesGetters.go.
// Setters and mutating methods are in AttributesSetters.go.
// JSON serialization and deserialization are in AttributesJson.go.
type Attributes struct {
	BasicErrWrapper  errcoreinf.BasicErrWrapper `json:"BasicErrWrapper,omitempty"`
	AuthInfo         *AuthInfo                  `json:"AuthInfo,omitempty"`
	PagingInfo       *PagingInfo                `json:"PagingInfo,omitempty"`
	KeyValuePairs    *corestr.Hashmap           `json:"KeyValuePairs,omitempty"`
	AnyKeyValuePairs *coredynamic.MapAnyItems   `json:"AnyKeyValuePairs,omitempty"`
	FromTo           *coreinstruction.FromTo    `json:"FromTo,omitempty"` // Invoker and Receiver Details
	DynamicPayloads  []byte                     `json:"DynamicPayloads,omitempty"`
}

// =============================================================================
// Clone and equality
// =============================================================================

func (it *Attributes) IsEqual(attributes *Attributes) bool {
	if it == nil && attributes == nil {
		return true
	}

	if it == nil || attributes == nil {
		return false
	}

	if it == attributes {
		return true
	}

	if it.IsErrorDifferent(attributes.BasicErrWrapper) {
		return false
	}

	isPagingDifferent := !it.PagingInfo.IsEqual(attributes.PagingInfo)

	if isPagingDifferent {
		return false
	}

	isKeyValuesDifferent := !it.KeyValuePairs.IsEqualPtr(attributes.KeyValuePairs)

	if isKeyValuesDifferent {
		return false
	}

	isDynamicPayloadsDifferent := !bytes.Equal(
		it.DynamicPayloads,
		attributes.DynamicPayloads)

	if isDynamicPayloadsDifferent {
		return false
	}

	isAnyKeyValuesDifferent := !it.AnyKeyValuePairs.IsEqual(attributes.AnyKeyValuePairs)

	if isAnyKeyValuesDifferent {
		return false
	}

	return true
}

func (it *Attributes) Clone(
	isDeepClone bool,
) (Attributes, error) {
	clonedPtr, err := it.ClonePtr(isDeepClone)

	if err != nil {
		return Attributes{}, err
	}

	if clonedPtr == nil {
		return Attributes{}, nil
	}

	return clonedPtr.NonPtr(), nil
}

func (it *Attributes) ClonePtr(
	isDeepClone bool,
) (*Attributes, error) {
	if it == nil {
		return nil, nil
	}

	if isDeepClone {
		return it.deepClonePtr()
	}

	// NOT deep clone
	return New.
		Attributes.
		All(
			it.AuthInfo,
			it.KeyValuePairs,
			it.AnyKeyValuePairs,
			it.PagingInfo,
			it.DynamicPayloads,
			it.FromTo,
			it.BasicErrWrapper,
		), nil
}

func (it *Attributes) deepClonePtr() (*Attributes, error) {
	var anyMap *coredynamic.MapAnyItems

	if it.AnyKeyValuePairs != nil {
		var err error
		anyMap, err = it.AnyKeyValuePairs.ClonePtr()

		if err != nil {
			return nil, err
		}
	}

	var basicErr errcoreinf.BasicErrWrapper

	if it.HasError() {
		basicErr = it.BasicErrWrapper.CloneInterface()
	}

	return New.
		Attributes.
		All(
			it.AuthInfo.ClonePtr(),
			it.KeyValuePairs.ClonePtr(),
			anyMap,
			it.PagingInfo.ClonePtr(),
			corejson.BytesDeepClone(it.DynamicPayloads),
			it.FromTo.ClonePtr(),
			basicErr), nil
}
