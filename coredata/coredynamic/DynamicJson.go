package coredynamic

import (
	"encoding/json"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/defaulterr"
	"github.com/alimtvnetwork/core/errcore"
)

// DynamicJson.go — JSON serialization, deserialization, and string conversion
// methods extracted from Dynamic.go.

func (it *Dynamic) Deserialize(jsonBytes []byte) (deserialized *Dynamic, err error) {
	if it == nil {
		return InvalidDynamicPtr(), defaulterr.UnmarshallingFailedDueToNilOrEmpty
	}

	err = corejson.
		Deserialize.
		UsingBytes(jsonBytes, it.innerData)

	it.isValid = err == nil

	return it, err
}

func (it *Dynamic) ValueMarshal() (jsonBytes []byte, err error) {
	if it == nil {
		return nil, defaulterr.NilResult
	}

	return corejson.
		Serialize.
		ToBytesErr(it.innerData)
}

func (it *Dynamic) JsonPayloadMust() (jsonBytes []byte) {
	return corejson.
		Serialize.
		ToSafeBytesMust(it.innerData)
}

// JsonBytesPtr returns empty bytes on nil.
// no error on nil.
func (it *Dynamic) JsonBytesPtr() (jsonBytes []byte, err error) {
	if it.IsNull() {
		return []byte{}, nil
	}

	marshalledBytes, marshalErr := json.Marshal(it.innerData)

	if marshalErr != nil {
		return []byte{}, marshalErr
	}

	return marshalledBytes, nil
}

func (it *Dynamic) MarshalJSON() ([]byte, error) {
	return corejson.
		Serialize.
		ToBytesErr(it.innerData)
}

func (it *Dynamic) UnmarshalJSON(data []byte) error {
	if it == nil {
		return defaulterr.UnmarshallingFailedDueToNilOrEmpty
	}

	err := corejson.
		Deserialize.
		UsingBytes(data, it.innerData)

	it.isValid = err == nil

	return err
}

func (it Dynamic) JsonModel() any {
	return it.innerData
}

func (it Dynamic) JsonModelAny() any {
	return it.JsonModel()
}

func (it Dynamic) Json() corejson.Result {
	return corejson.New(it)
}

func (it Dynamic) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

//goland:noinspection GoLinterLocal
func (it *Dynamic) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Dynamic, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return nil, err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
//
//goland:noinspection GoLinterLocal
func (it *Dynamic) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Dynamic {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *Dynamic) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *Dynamic) JsonBytes() (jsonBytesPtr []byte, err error) {
	allBytes, err := it.JsonBytesPtr()

	if err != nil {
		return []byte{}, err
	}

	return allBytes, err
}

func (it *Dynamic) JsonString() (jsonString string, err error) {
	marshalledBytes, err := it.JsonBytes()

	if err != nil {
		return constants.EmptyString, err
	}

	return string(marshalledBytes), err
}

func (it *Dynamic) JsonStringMust() string {
	marshalledBytes, err := it.JsonBytes()

	if err != nil {
		errcore.
			MarshallingFailedType.
			HandleUsingPanic(err.Error(), it.innerDataString)
	}

	return string(marshalledBytes)
}
