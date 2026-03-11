package converters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coreappend"
	"github.com/alimtvnetwork/core/internal/convertinternal"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

type anyItemConverter struct{}

func (it anyItemConverter) ToString(
	isIncludeFullName bool,
	anyVal any,
) string {
	if anyVal == nil {
		return ""
	}

	if isIncludeFullName {
		return fmt.Sprintf(
			constants.SprintFullPropertyNameValueFormat,
			anyVal,
		)
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		anyVal,
	)
}

func (it anyItemConverter) String(
	anyVal any,
) string {
	if anyVal == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		anyVal,
	)
}

func (it anyItemConverter) FullString(
	anyVal any,
) string {
	if anyVal == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		anyVal,
	)
}

func (it anyItemConverter) StringWithType(
	anyVal any,
) string {
	if anyVal == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintPropertyValueWithTypeFormat,
		anyVal,
		anyVal,
	)
}

// ToSafeSerializedString
//
//	warning : on error swallows it
func (it anyItemConverter) ToSafeSerializedString(
	anyVal any,
) string {
	if anyVal == nil {
		return ""
	}

	switch casted := anyVal.(type) {
	case []byte:
		return BytesTo.String(casted)
	}

	allBytes, _ := json.Marshal(anyVal)

	return BytesTo.String(allBytes)
}

// ToSafeSerializedStringSprintValue
//
//	return value using %v
//
//	warning : on error swallows it
func (it anyItemConverter) ToSafeSerializedStringSprintValue(
	anyVal any,
) string {
	value := it.ToSafeSerializedString(
		anyVal,
	)

	return fmt.Sprintf(
		constants.SprintValueFormat,
		value,
	)
}

func (it anyItemConverter) ToStrings(
	isSkipOnNil bool,
	anyItem any,
) []string {
	if isSkipOnNil && anyItem == nil {
		return []string{}
	}

	reflectVal := reflect.ValueOf(anyItem)

	anyItems := reflectinternal.Converter.ReflectValToInterfaces(
		isSkipOnNil,
		reflectVal,
	)

	return it.ItemsToStringsSkipOnNil(anyItems)
}

func (it anyItemConverter) ToStringsUsingProcessor(
	isSkipOnNil bool,
	processor func(index int, in any) (out string, isTake, isBreak bool),
	anyVal any,
) []string {
	if anyVal == nil {
		return []string{}
	}

	anyItems := it.ToAnyItems(isSkipOnNil, anyVal)
	slice := make([]string, 0, len(anyItems))

	if len(anyItems) == 0 {
		return slice
	}

	for i, item := range anyItems {
		out, isTake, isBreak := processor(i, item)

		if isTake {
			slice = append(slice, out)
		}

		if isBreak {
			return slice
		}
	}

	return slice
}

func (it anyItemConverter) ToStringsUsingSimpleProcessor(
	isSkipOnNil bool,
	simpleProcessor func(index int, in any) (out string),
	anyVal any,
) []string {
	if anyVal == nil {
		return []string{}
	}

	anyItems := it.ToAnyItems(isSkipOnNil, anyVal)
	slice := make([]string, len(anyItems))

	if len(anyItems) == 0 {
		return slice
	}

	for i, item := range anyItems {
		out := simpleProcessor(i, item)

		slice[i] = out
	}

	return slice
}

func (it anyItemConverter) ToValueString(
	anyVal any,
) string {
	if anyVal == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		anyVal,
	)
}

func (it anyItemConverter) ToValueStringWithType(
	anyVal any,
) string {
	if anyVal == nil {
		return fmt.Sprintf(
			constants.SprintNilValueTypeInParenthesisFormat,
			anyVal,
		)
	}

	return fmt.Sprintf(
		constants.SprintValueWithTypeFormat,
		anyVal,
		anyVal,
	)
}

func (it anyItemConverter) ToAnyItems(
	isSkipOnNil bool,
	anyItem any,
) []any {
	if isSkipOnNil && anyItem == nil {
		return []any{}
	}

	reflectVal := reflect.ValueOf(anyItem)

	return reflectinternal.Converter.ReflectValToInterfaces(
		isSkipOnNil,
		reflectVal,
	)
}

func (it anyItemConverter) ToNonNullItems(
	isSkipOnNil bool,
	anyItem any,
) []any {
	if isSkipOnNil && (anyItem == nil || reflectinternal.Is.Null(anyItem)) {
		return []any{}
	}

	reflectVal := reflect.ValueOf(anyItem)

	return reflectinternal.Converter.ReflectValToInterfaces(
		isSkipOnNil,
		reflectVal,
	)
}

func (it anyItemConverter) ItemsToStringsSkipOnNil(
	anyItems ...any,
) []string {
	return coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
		nil,
		nil,
		anyItems...,
	)
}

func (it anyItemConverter) ItemsJoin(
	joiner string,
	anyItems ...any,
) string {
	if anyItems == nil {
		return constants.EmptyString
	}

	anyStrings := it.ItemsToStringsSkipOnNil(anyItems...)

	return strings.Join(anyStrings, joiner)
}

func (it anyItemConverter) ToItemsThenJoin(
	isSkipOnNil bool,
	joiner string,
	anySlice any,
) string {
	if anySlice == nil {
		return constants.EmptyString
	}

	anyStrings := it.ToStrings(
		isSkipOnNil,
		anySlice,
	)

	return strings.Join(
		anyStrings,
		joiner,
	)
}

func (it anyItemConverter) ToFullNameValueString(
	anyVal any,
) string {
	if anyVal == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintFullPropertyNameValueFormat,
		anyVal,
	)
}

// ToPrettyJson
//
// Warning:
//
//	swallows error
func (it anyItemConverter) ToPrettyJson(
	anyItem any,
) string {
	if anyItem == nil {
		return ""
	}

	allBytes, err := json.Marshal(anyItem)

	if err != nil || len(allBytes) == 0 {
		return ""
	}

	var prettyJSON bytes.Buffer

	json.Indent(
		&prettyJSON,
		allBytes,
		constants.EmptyString,
		constants.Tab,
	)

	return prettyJSON.String()
}

// Bytes
//
// ## Steps:
//   - If already in  []byte then return as is.
//   - If already in *[]byte then return as []byte without pointer by checking if not null.
//   - If already in  string then return as []byte(string).
//   - For rest of the cases, convert to json using Marshal and then returns the bytes
//
// Panic if json marshal has error.
func (it anyItemConverter) Bytes(anyItem any) []byte {
	switch expectedAs := anyItem.(type) {
	case []byte:
		if expectedAs == nil {
			return []byte{}
		}

		return expectedAs
	case *[]byte:
		if expectedAs == nil || *expectedAs == nil {
			return []byte{}
		}

		return *expectedAs
	case string:
		return []byte(expectedAs)
	default:
		toBytes, err := json.Marshal(expectedAs)

		if err != nil {
			panic(err)
		}

		return toBytes
	}
}

// ValueString
//
// If nil then returns ""
// Or, returns %v of the value given.
func (it anyItemConverter) ValueString(anyItem any) string {
	if anyItem == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		anyItem,
	)
}

// SmartString
//
//   - If nil return ""
//   - If string return just returns
//   - Or, else return %v of value
func (it anyItemConverter) SmartString(anyItem any) string {
	if anyItem == nil {
		return ""
	}

	return convertinternal.AnyTo.SmartString(anyItem)
}

func (it anyItemConverter) SmartStringsJoiner(
	joiner string,
	anyItems ...any,
) string {
	if len(anyItems) == 0 {
		return ""
	}

	slice := make([]string, len(anyItems))
	converterFunc := convertinternal.AnyTo.SmartJson

	for i, elem := range anyItems {
		slice[i] = converterFunc(elem)
	}

	return strings.Join(slice, joiner)
}

func (it anyItemConverter) SmartStringsOf(
	anyItems ...any,
) string {
	if len(anyItems) == 0 {
		return ""
	}

	return it.SmartStringsJoiner(", ", anyItems...)
}
