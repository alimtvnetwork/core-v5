package coredynamic

import (
	"fmt"
	"reflect"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

type KeyVal struct {
	Key   any
	Value any
}

func (it KeyVal) KeyDynamic() Dynamic {
	return NewDynamic(it.Key, true)
}

func (it KeyVal) ValueDynamic() Dynamic {
	return NewDynamic(it.Value, true)
}

func (it *KeyVal) KeyDynamicPtr() *Dynamic {
	if it == nil {
		return nil
	}

	return NewDynamicPtr(it.Key, true)
}

func (it *KeyVal) ValueDynamicPtr() *Dynamic {
	if it == nil {
		return nil
	}

	return NewDynamicPtr(it.Value, true)
}

func (it KeyVal) IsKeyNull() bool {
	return reflectinternal.Is.Null(it.Key)
}

func (it KeyVal) IsKeyNullOrEmptyString() bool {
	return reflectinternal.Is.Null(it.Key) || it.Key.(string) == ""
}

func (it KeyVal) IsValueNull() bool {
	return reflectinternal.Is.Null(it.Value)
}

func (it *KeyVal) String() string {
	if it == nil {
		return constants.EmptyString
	}

	return fmt.Sprintf(
		constants.KeyValuePariSimpleFormat,
		it.Key,
		it.Key,
		it.Value,
		it.Value,
	)
}

func (it KeyVal) ValueReflectValue() reflect.Value {
	return reflect.ValueOf(it.Value)
}

func (it KeyVal) ValueInt() int {
	casted, isSuccess := it.Value.(int)

	if isSuccess {
		return casted
	}

	return constants.InvalidValue
}

func (it KeyVal) ValueUInt() uint {
	casted, isSuccess := it.Value.(uint)

	if isSuccess {
		return casted
	}

	return constants.Zero
}

func (it KeyVal) ValueStrings() []string {
	casted, isSuccess := it.Value.([]string)

	if isSuccess {
		return casted
	}

	return nil
}

func (it KeyVal) ValueBool() bool {
	casted, isSuccess := it.Value.(bool)

	if isSuccess {
		return casted
	}

	return false
}

func (it KeyVal) ValueInt64() int64 {
	casted, isSuccess := it.Value.(int64)

	if isSuccess {
		return casted
	}

	return constants.InvalidValue
}

func (it *KeyVal) CastKeyVal(
	keyToPointer,
	valueToPointer any,
) error {
	if it == nil {
		return errcore.
			CannotBeNilOrEmptyType.
			ErrorNoRefs("KeyVal is nil or null")
	}

	err := ReflectSetFromTo(it.Key, keyToPointer)

	if err != nil {
		return nil
	}

	return ReflectSetFromTo(it.Value, valueToPointer)
}

func (it *KeyVal) ReflectSetKey(
	keyToPointer any,
) error {
	if it == nil {
		return errcore.
			CannotBeNilOrEmptyType.
			ErrorNoRefs("KeyVal is nil or null")
	}

	return ReflectSetFromTo(it.Key, keyToPointer)
}

func (it *KeyVal) ValueNullErr() error {
	if it == nil {
		return errcore.
			CannotBeNilOrEmptyType.
			ErrorNoRefs("KeyVal is nil or null")
	}

	if reflectinternal.Is.Null(it.Value) {
		return errcore.
			CannotBeNilOrEmptyType.
			Error("KeyVal.Value is nil or null, doesn't expect to be null.", "Key : "+it.KeyString())
	}

	return nil
}

func (it *KeyVal) KeyNullErr() error {
	if it == nil {
		return errcore.
			CannotBeNilOrEmptyType.
			ErrorNoRefs("KeyVal is nil or null")
	}

	if reflectinternal.Is.Null(it.Key) {
		return errcore.
			CannotBeNilOrEmptyType.
			Error("KeyVal.Key is nil or null, doesn't expect to be null.", "Value : "+it.ValueString())
	}

	return nil
}

func (it *KeyVal) KeyString() string {
	if it == nil || it.Key == nil {
		return constants.EmptyString
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.Key,
	)
}

func (it *KeyVal) ValueString() string {
	if it == nil || it.Value == nil {
		return constants.EmptyString
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.Value,
	)
}

func (it *KeyVal) KeyReflectSet(toPointer any) error {
	if it == nil {
		return errcore.
			CannotBeNilOrEmptyType.
			ErrorNoRefs("KeyVal is nil or null")
	}

	return ReflectSetFromTo(it.Key, toPointer)
}

func (it *KeyVal) ValueReflectSet(toPointer any) error {
	if it == nil {
		return errcore.
			CannotBeNilOrEmptyType.
			ErrorNoRefs("KeyVal is nil or null")
	}

	return ReflectSetFromTo(it.Value, toPointer)
}

func (it *KeyVal) ReflectSetTo(toPointer any) error {
	if it == nil {
		return errcore.
			CannotBeNilOrEmptyType.
			ErrorNoRefs("KeyVal is nil or null")
	}

	return ReflectSetFromTo(it.Value, toPointer)
}

func (it *KeyVal) ReflectSetToMust(toPointer any) {
	err := it.ReflectSetTo(toPointer)
	errcore.MustBeEmpty(err)
}

func (it KeyVal) JsonModel() any {
	return it
}

func (it KeyVal) JsonModelAny() any {
	return it.JsonModel()
}

func (it KeyVal) Json() corejson.Result {
	return corejson.New(it)
}

func (it KeyVal) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

//goland:noinspection GoLinterLocal
func (it *KeyVal) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*KeyVal, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return nil, err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
//
//goland:noinspection GoLinterLocal
func (it *KeyVal) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *KeyVal {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *KeyVal) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *KeyVal) Serialize() (jsonBytesPtr []byte, err error) {
	jsonResult := it.Json()

	if jsonResult.HasError() {
		return []byte{}, jsonResult.MeaningfulError()
	}

	return jsonResult.SafeBytes(), nil
}
