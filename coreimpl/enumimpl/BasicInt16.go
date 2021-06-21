package enumimpl

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/core/simplewrap"
)

type BasicInt16 struct {
	*numberEnumBase
	jsonDoubleQuoteNameToValueHashMap        map[string]int16 // contains names double quotes to value
	valueToJsonDoubleQuoteStringBytesHashmap map[int16][]byte // contains value to string bytes with double quotes
	valueNameHashmap                         map[int16]string // contains name without double quotes
	minVal, maxVal                           int16
}

func NewBasicInt16(
	actualValueRanges []int16,
	stringRanges []string,
	min, max int16,
) *BasicInt16 {
	enumBase := newNumberEnumBase(
		actualValueRanges,
		stringRanges,
		min,
		max)

	jsonDoubleQuoteNameToValueHashMap := make(map[string]int16, len(actualValueRanges))
	valueToJsonDoubleQuoteStringBytesHashmap := make(map[int16][]byte, len(actualValueRanges))
	valueNameHashmap := make(map[int16]string, len(actualValueRanges))

	for i, actualVal := range actualValueRanges {
		key := stringRanges[i]
		jsonName := simplewrap.WithDoubleQuote(key)
		jsonDoubleQuoteNameToValueHashMap[jsonName] = actualVal
		valueToJsonDoubleQuoteStringBytesHashmap[actualVal] = []byte(jsonName)
		valueNameHashmap[actualVal] = key
	}

	return &BasicInt16{
		numberEnumBase:                           enumBase,
		minVal:                                   min,
		maxVal:                                   max,
		jsonDoubleQuoteNameToValueHashMap:        jsonDoubleQuoteNameToValueHashMap,
		valueToJsonDoubleQuoteStringBytesHashmap: valueToJsonDoubleQuoteStringBytesHashmap,
		valueNameHashmap:                         valueNameHashmap,
	}
}

func NewBasicInt16UsingIndexedSlice(
	indexedSliceWithValues []string,
) *BasicInt16 {
	min := constants.Zero
	max := len(indexedSliceWithValues)

	actualValues := make([]int16, max)
	for i := range indexedSliceWithValues {
		actualValues[i] = int16(i)
	}

	return NewBasicInt16(
		actualValues,
		indexedSliceWithValues,
		int16(min),
		int16(max))
}

func (receiver *BasicInt16) IsAnyOf(value int16, checkingItems ...int16) bool {
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

func (receiver *BasicInt16) Max() int16 {
	return receiver.maxVal
}

func (receiver *BasicInt16) Min() int16 {
	return receiver.minVal
}

func (receiver *BasicInt16) GetValueByString(valueString string) int16 {
	return receiver.jsonDoubleQuoteNameToValueHashMap[valueString]
}

func (receiver *BasicInt16) GetStringValue(input int16) string {
	return receiver.StringRanges()[input]
}

func (receiver *BasicInt16) Ranges() []int16 {
	return receiver.actualValueRanges.([]int16)
}

func (receiver *BasicInt16) Hashmap() map[string]int16 {
	return receiver.jsonDoubleQuoteNameToValueHashMap
}

func (receiver *BasicInt16) HashmapPtr() *map[string]int16 {
	return &receiver.jsonDoubleQuoteNameToValueHashMap
}

func (receiver *BasicInt16) IsValidRange(value int16) bool {
	return value >= receiver.minVal && value <= receiver.maxVal
}

// ToEnumJsonBytes used for MarshalJSON from map
func (receiver *BasicInt16) ToEnumJsonBytes(value int16) []byte {
	return receiver.valueToJsonDoubleQuoteStringBytesHashmap[value]
}

func (receiver *BasicInt16) ToEnumString(value int16) string {
	return receiver.valueNameHashmap[value]
}

func (receiver *BasicInt16) ToNumberString(valueInRawFormat interface{}) string {
	return fmt.Sprintf(constants.SprintValueFormat, valueInRawFormat)
}

// UnmarshallEnumToValue Mostly used for UnmarshalJSON
//
// Given bytes string enum value and transpile to exact enum raw value using map
func (receiver *BasicInt16) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (int16, error) {
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
