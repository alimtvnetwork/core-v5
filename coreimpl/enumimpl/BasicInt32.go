package enumimpl

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/core/simplewrap"
)

type BasicInt32 struct {
	*numberEnumBase
	jsonDoubleQuoteNameToValueHashMap        map[string]int32 // contains names double quotes to value
	valueToJsonDoubleQuoteStringBytesHashmap map[int32][]byte // contains value to string bytes with double quotes
	valueNameHashmap                         map[int32]string // contains name without double quotes
	minVal, maxVal                           int32
}

func NewBasicInt32(
	actualValueRanges []int32,
	stringRanges []string,
	min, max int32,
) *BasicInt32 {
	enumBase := newNumberEnumBase(
		actualValueRanges,
		stringRanges,
		min,
		max)

	jsonDoubleQuoteNameToValueHashMap := make(map[string]int32, len(actualValueRanges))
	valueToJsonDoubleQuoteStringBytesHashmap := make(map[int32][]byte, len(actualValueRanges))
	valueNameHashmap := make(map[int32]string, len(actualValueRanges))

	for i, actualVal := range actualValueRanges {
		key := stringRanges[i]
		jsonName := simplewrap.WithDoubleQuote(key)
		jsonDoubleQuoteNameToValueHashMap[jsonName] = actualVal
		valueToJsonDoubleQuoteStringBytesHashmap[actualVal] = []byte(jsonName)
		valueNameHashmap[actualVal] = key
	}

	return &BasicInt32{
		numberEnumBase:                           enumBase,
		minVal:                                   min,
		maxVal:                                   max,
		jsonDoubleQuoteNameToValueHashMap:        jsonDoubleQuoteNameToValueHashMap,
		valueToJsonDoubleQuoteStringBytesHashmap: valueToJsonDoubleQuoteStringBytesHashmap,
		valueNameHashmap:                         valueNameHashmap,
	}
}

func (receiver *BasicInt32) IsAnyOf(value int32, checkingItems ...int32) bool {
	if len(checkingItems) == 0 {
		return true
	}

	for _, givenByte := range checkingItems {
		if value == givenByte {
			return true
		}
	}

	return false
}

func (receiver *BasicInt32) Max() int32 {
	return receiver.maxVal
}

func (receiver *BasicInt32) Min() int32 {
	return receiver.minVal
}

func (receiver *BasicInt32) GetValueByString(valueString string) int32 {
	return receiver.jsonDoubleQuoteNameToValueHashMap[valueString]
}

func (receiver *BasicInt32) GetStringValue(input int32) string {
	return receiver.StringRanges()[input]
}

func (receiver *BasicInt32) Ranges() []int32 {
	return receiver.actualValueRanges.([]int32)
}

func (receiver *BasicInt32) Hashmap() map[string]int32 {
	return receiver.jsonDoubleQuoteNameToValueHashMap
}

func (receiver *BasicInt32) HashmapPtr() *map[string]int32 {
	return &receiver.jsonDoubleQuoteNameToValueHashMap
}

func (receiver *BasicInt32) IsValidRange(value int32) bool {
	return value >= receiver.minVal && value <= receiver.maxVal
}

// ToEnumJsonBytes used for MarshalJSON from map
func (receiver *BasicInt32) ToEnumJsonBytes(value int32) []byte {
	return receiver.valueToJsonDoubleQuoteStringBytesHashmap[value]
}

func (receiver *BasicInt32) ToEnumString(value int32) string {
	return receiver.valueNameHashmap[value]
}

func (receiver *BasicInt32) ToNumberString(valueInRawFormat interface{}) string {
	return fmt.Sprintf(constants.SprintValueFormat, valueInRawFormat)
}

// UnmarshallEnumToValue Mostly used for UnmarshalJSON
//
// Given bytes string enum value and transpile to exact enum raw value using map
func (receiver *BasicInt32) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (int32, error) {
	if jsonUnmarshallingValue == nil {
		return constants.Zero,
			defaulterr.UnMarshallingFailedDueToNilOrEmpty
	}

	str := string(jsonUnmarshallingValue)
	v, has := receiver.jsonDoubleQuoteNameToValueHashMap[str]

	if !has {
		return constants.Zero,
			msgtype.MeaningFulErrorWithData(
				msgtype.UnMarshallingFailed,
				"UnmarshallEnumToValue",
				defaulterr.UnMarshallingPlusCannotFindingEnumMap,
				string(jsonUnmarshallingValue))
	}

	return v, nil
}
