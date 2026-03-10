package corepayload

import (
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coreinterface/payloadinf"
	"github.com/alimtvnetwork/core/defaulterr"
)

// AttributesJson.go — JSON serialization, deserialization, and string methods extracted from Attributes.go

func (it *Attributes) PayloadsPrettyString() string {
	if it.IsEmpty() || len(it.DynamicPayloads) == 0 {
		return ""
	}

	return corejson.BytesToPrettyString(it.DynamicPayloads)
}

func (it *Attributes) PayloadsJsonResult() *corejson.Result {
	if it.IsEmpty() || len(it.DynamicPayloads) == 0 {
		return corejson.Empty.ResultPtr()
	}

	return corejson.NewResult.UsingTypeBytesPtr(
		attributesTypeName,
		it.DynamicPayloads)
}

func (it Attributes) JsonString() string {
	return it.JsonPtr().JsonString()
}

func (it Attributes) JsonStringMust() string {
	jsonResult := it.JsonPtr()
	jsonResult.MustBeSafe()

	return jsonResult.JsonString()
}

func (it Attributes) String() string {
	return it.JsonString()
}

func (it Attributes) PrettyJsonString() string {
	return it.JsonPtr().PrettyJsonString()
}

func (it Attributes) Json() corejson.Result {
	return corejson.New(it)
}

func (it Attributes) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it Attributes) JsonModel() Attributes {
	return it
}

func (it Attributes) JsonModelAny() any {
	return it.JsonModel()
}

//goland:noinspection GoLinterLocal
func (it *Attributes) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Attributes, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return &Attributes{}, err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
//
//goland:noinspection GoLinterLocal
func (it *Attributes) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Attributes {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *Attributes) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it Attributes) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}

// BasicErrorDeserializedTo
//
// Expectation Attributes.ErrorMessage needs to
// be in json format and toPtr
// should match reflection types
func (it *Attributes) BasicErrorDeserializedTo(
	toPtr any,
) error {
	if it.IsEmptyError() {
		return nil
	}

	return corejson.
		Deserialize.
		UsingBytes(
			it.BasicErrWrapper.SerializeMust(),
			toPtr)
}

func (it *Attributes) DeserializeDynamicPayloads(
	toPtr any,
) error {
	return corejson.
		Deserialize.
		UsingBytes(
			it.DynamicPayloads,
			toPtr)
}

func (it *Attributes) DeserializeDynamicPayloadsToAttributes() (
	newAttr *Attributes, err error,
) {
	newAttr = &Attributes{}
	err = corejson.Deserialize.UsingBytes(
		it.DynamicPayloads,
		newAttr)

	return newAttr, err
}

func (it *Attributes) DeserializeDynamicPayloadsToPayloadWrapper() (
	payloadWrapper *PayloadWrapper, err error,
) {
	payloadWrapper = New.PayloadWrapper.Empty()
	err = corejson.Deserialize.UsingBytes(
		it.DynamicPayloads,
		payloadWrapper)

	return payloadWrapper, err
}

func (it *Attributes) DeserializeDynamicPayloadsToPayloadWrappersCollection() (
	payloadsCollection *PayloadsCollection, err error,
) {
	return New.
		PayloadsCollection.
		Deserialize(
			it.DynamicPayloads)
}

func (it *Attributes) DeserializeDynamicPayloadsMust(
	toPtr any,
) {
	corejson.Deserialize.
		UsingBytesMust(
			it.DynamicPayloads,
			toPtr)
}

func (it *Attributes) DynamicPayloadsDeserialize(
	unmarshallingPointer any,
) error {
	if it == nil {
		return defaulterr.AttributeNull
	}

	return corejson.Deserialize.UsingBytes(
		it.DynamicPayloads,
		unmarshallingPointer)
}

func (it *Attributes) DynamicPayloadsDeserializeMust(
	unmarshallingPointer any,
) {
	err := corejson.Deserialize.UsingBytes(
		it.DynamicPayloads,
		unmarshallingPointer)

	if err != nil {
		panic(err)
	}
}

func (it Attributes) NonPtr() Attributes {
	return it
}

func (it Attributes) AsAttributesBinder() payloadinf.AttributesBinder {
	return &it
}
