package coredynamic

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/constants/bitsize"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/internal/messages"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
	"github.com/alimtvnetwork/core/internal/strutilinternal"
	"github.com/alimtvnetwork/core/issetter"
)

// DynamicGetters.go — Read-only accessors, type checks, and value extraction
// methods extracted from Dynamic.go.

func (it Dynamic) Data() any {
	return it.innerData
}

func (it Dynamic) Value() any {
	return it.innerData
}

// Length Returns length of a slice, map, array
//
// # It will also reduce from pointer
//
// Reference : https://cutt.ly/PnaWAFn | https://cutt.ly/jnaEig8 | https://play.golang.org/p/UCORoShXlv1
func (it *Dynamic) Length() int {
	if it == nil {
		return constants.Zero
	}

	return it.length.Value()
}

func (it *Dynamic) StructStringPtr() *string {
	if it == nil {
		return nil
	}

	if it.innerDataString != nil {
		return it.innerDataString
	}

	toString := strutilinternal.AnyToString(it.innerData)
	it.innerDataString = &toString

	return it.innerDataString
}

func (it *Dynamic) String() string {
	if it == nil {
		return constants.EmptyString
	}

	return *it.StructStringPtr()
}

func (it *Dynamic) StructString() string {
	if it == nil {
		return constants.EmptyString
	}

	return *it.StructStringPtr()
}

func (it Dynamic) IsNull() bool {
	return it.innerData == nil
}

func (it Dynamic) IsValid() bool {
	return it.isValid
}

func (it Dynamic) IsInvalid() bool {
	return !it.isValid
}

func (it *Dynamic) IsPointer() bool {
	if it == nil {
		return false
	}

	if it.isPointer.IsUninitialized() {
		it.isPointer = issetter.GetBool(
			it.IsReflectKind(reflect.Ptr),
		)
	}

	return it.isPointer.IsTrue()
}

func (it *Dynamic) IsValueType() bool {
	if it == nil {
		return false
	}

	return !it.IsPointer()
}

func (it *Dynamic) IsStructStringNullOrEmpty() bool {
	if it == nil {
		return true
	}

	return it.IsNull() || strutilinternal.IsNullOrEmpty(
		it.StructStringPtr(),
	)
}

func (it *Dynamic) IsStructStringNullOrEmptyOrWhitespace() bool {
	if it == nil {
		return true
	}

	return it.IsNull() || strutilinternal.IsNullOrEmptyOrWhitespace(
		it.StructStringPtr(),
	)
}

func (it *Dynamic) IsPrimitive() bool {
	if it == nil {
		return false
	}

	return reflectinternal.Is.PrimitiveKind(it.ReflectKind())
}

// IsNumber true if float (any), byte, int (any), uint(any)
func (it *Dynamic) IsNumber() bool {
	if it == nil {
		return false
	}

	return reflectinternal.Is.NumberKind(it.ReflectKind())
}

func (it *Dynamic) IsStringType() bool {
	if it == nil {
		return false
	}

	_, isString := it.innerData.(string)

	return isString
}

func (it *Dynamic) IsStruct() bool {
	if it == nil {
		return false
	}

	return it.ReflectKind() == reflect.Struct
}

func (it *Dynamic) IsFunc() bool {
	if it == nil {
		return false
	}

	return it.ReflectKind() == reflect.Func
}

func (it *Dynamic) IsSliceOrArray() bool {
	if it == nil {
		return false
	}

	k := it.ReflectKind()

	return k == reflect.Slice || k == reflect.Array
}

func (it *Dynamic) IsSliceOrArrayOrMap() bool {
	if it == nil {
		return false
	}

	k := it.ReflectKind()

	return k == reflect.Slice ||
		k == reflect.Array ||
		k == reflect.Map
}

func (it *Dynamic) IsMap() bool {
	if it == nil {
		return false
	}

	return it.ReflectKind() == reflect.Map
}

// =============================================================================
// Value extraction
// =============================================================================

func (it *Dynamic) IntDefault(defaultInt int) (val int, isSuccess bool) {
	if it == nil || it.IsNull() {
		return defaultInt, false
	}

	stringVal := it.StructString()
	toInt, err := strconv.Atoi(stringVal)

	if err == nil {
		return toInt, true
	}

	return defaultInt, false
}

func (it *Dynamic) Float64() (val float64, err error) {
	if it == nil || it.IsNull() {
		return constants.Zero, errcore.
			ParsingFailedType.Error(
			messages.DynamicFailedToParseToFloat64BecauseNull,
			it.String(),
		)
	}

	stringVal := it.StructString()
	valFloat, parseErr := strconv.ParseFloat(stringVal, bitsize.Of64)

	if parseErr != nil {
		reference := stringVal +
			constants.NewLineUnix +
			parseErr.Error()

		return constants.Zero, errcore.
			ParsingFailedType.Error(
			errcore.FailedToConvertType.String(),
			reference,
		)
	}

	return valFloat, nil
}

func (it Dynamic) ValueInt() int {
	casted, isSuccess := it.innerData.(int)

	if isSuccess {
		return casted
	}

	return constants.InvalidValue
}

func (it Dynamic) ValueUInt() uint {
	casted, isSuccess := it.innerData.(uint)

	if isSuccess {
		return casted
	}

	return constants.Zero
}

func (it Dynamic) ValueStrings() []string {
	casted, isSuccess := it.innerData.([]string)

	if isSuccess {
		return casted
	}

	return nil
}

func (it Dynamic) ValueBool() bool {
	casted, isSuccess := it.innerData.(bool)

	if isSuccess {
		return casted
	}

	return false
}

func (it Dynamic) ValueInt64() int64 {
	casted, isSuccess := it.innerData.(int64)

	if isSuccess {
		return casted
	}

	return constants.InvalidValue
}

func (it *Dynamic) ValueNullErr() error {
	if it == nil {
		return errcore.
			CannotBeNilOrEmptyType.
			ErrorNoRefs("Dynamic is nil or null")
	}

	if reflectinternal.Is.Null(it.innerData) {
		return errcore.
			CannotBeNilOrEmptyType.
			ErrorNoRefs("Dynamic internal data is nil.")
	}

	return nil
}

func (it *Dynamic) ValueString() string {
	if it == nil || it.innerData == nil {
		return constants.EmptyString
	}

	currentString, isString := it.innerData.(string)

	if isString {
		return currentString
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.innerData,
	)
}

func (it *Dynamic) Bytes() (rawBytes []byte, isSuccess bool) {
	if it == nil {
		return nil, false
	}

	rawBytes, isSuccess = it.innerData.([]byte)

	return rawBytes, isSuccess
}
