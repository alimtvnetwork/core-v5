package enumimpl

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/core/simplewrap"
)

type BasicByte struct {
	*numberEnumBase
	jsonDoubleQuoteNameToValueHashMap        map[string]byte // contains names double quotes to value
	valueToJsonDoubleQuoteStringBytesHashmap map[byte][]byte // contains value to string bytes with double quotes
	valueNameHashmap                         map[byte]string // contains name without double quotes
	minVal, maxVal                           byte
}

func NewBasicByte(
	actualValueRanges []byte,
	stringRanges []string,
	min, max byte,
) *BasicByte {
	enumBase := newNumberEnumBase(
		&actualValueRanges,
		stringRanges,
		min,
		max)

	jsonDoubleQuoteNameToValueHashMap := make(map[string]byte, len(actualValueRanges))
	valueToJsonDoubleQuoteStringBytesHashmap := make(map[byte][]byte, len(actualValueRanges))
	valueNameHashmap := make(map[byte]string, len(actualValueRanges))

	for i, actualVal := range actualValueRanges {
		key := stringRanges[i]
		jsonName := simplewrap.WithDoubleQuote(key)
		jsonDoubleQuoteNameToValueHashMap[jsonName] = actualVal
		valueToJsonDoubleQuoteStringBytesHashmap[actualVal] = []byte(jsonName)
		valueNameHashmap[actualVal] = key
	}

	return &BasicByte{
		numberEnumBase:                           enumBase,
		minVal:                                   min,
		maxVal:                                   max,
		jsonDoubleQuoteNameToValueHashMap:        jsonDoubleQuoteNameToValueHashMap,
		valueToJsonDoubleQuoteStringBytesHashmap: valueToJsonDoubleQuoteStringBytesHashmap,
		valueNameHashmap:                         valueNameHashmap,
	}
}

func NewBasicByteUsingIndexedSlice(
	indexedSliceWithValues []string,
) *BasicByte {
	min := constants.Zero
	max := len(indexedSliceWithValues)

	actualValues := make([]byte, max)
	for i := range indexedSliceWithValues {
		actualValues[i] = byte(i)
	}

	return NewBasicByte(
		actualValues,
		indexedSliceWithValues,
		byte(min),
		byte(max))
}

func (receiver *BasicByte) IsAnyOf(value byte, givenBytes ...byte) bool {
	if len(givenBytes) == 0 {
		return true
	}

	for _, givenByte := range givenBytes {
		if value == givenByte {
			return true
		}
	}

	return false
}

func (receiver *BasicByte) Max() byte {
	return receiver.maxVal
}

func (receiver *BasicByte) Min() byte {
	return receiver.minVal
}

func (receiver *BasicByte) GetValueByString(valueString string) byte {
	return receiver.jsonDoubleQuoteNameToValueHashMap[valueString]
}

func (receiver *BasicByte) GetStringValue(input byte) string {
	return receiver.StringRanges()[input]
}

func (receiver *BasicByte) Ranges() []byte {
	return receiver.actualValueRanges.([]byte)
}

func (receiver *BasicByte) Hashmap() map[string]byte {
	return receiver.jsonDoubleQuoteNameToValueHashMap
}

func (receiver *BasicByte) HashmapPtr() *map[string]byte {
	return &receiver.jsonDoubleQuoteNameToValueHashMap
}

func (receiver *BasicByte) IsValidRange(value byte) bool {
	return value >= receiver.minVal && value <= receiver.maxVal
}

// ToEnumJsonBytes used for MarshalJSON from map
func (receiver *BasicByte) ToEnumJsonBytes(value byte) []byte {
	return receiver.valueToJsonDoubleQuoteStringBytesHashmap[value]
}

func (receiver *BasicByte) ToEnumString(value byte) string {
	return receiver.valueNameHashmap[value]
}

func (receiver *BasicByte) ToNumberString(valueInRawFormat interface{}) string {
	return fmt.Sprintf(constants.SprintValueFormat, valueInRawFormat)
}

// UnmarshallToValue Mostly used for UnmarshalJSON
//
// Given bytes string enum value and transpile to exact enum raw value using map
func (receiver *BasicByte) UnmarshallToValue(
	isMappedToFirstIfEmpty bool,
	jsonUnmarshallingValue []byte,
) (byte, error) {
	if !isMappedToFirstIfEmpty && jsonUnmarshallingValue == nil {
		return constants.Zero,
			defaulterr.UnMarshallingFailedDueToNilOrEmpty
	}

	if isMappedToFirstIfEmpty && jsonUnmarshallingValue == nil {
		return receiver.minVal, nil
	}

	str := string(jsonUnmarshallingValue)
	if isMappedToFirstIfEmpty &&
		(str == constants.EmptyString || str == constants.DoubleQuotationStartEnd) {
		return receiver.minVal, nil
	}

	v, has := receiver.jsonDoubleQuoteNameToValueHashMap[str]

	if !has {
		return constants.Zero,
			msgtype.MeaningfulErrorWithData(
				msgtype.UnMarshallingFailed,
				"UnmarshallToValue",
				defaulterr.UnMarshallingPlusCannotFindingEnumMap,
				string(jsonUnmarshallingValue))
	}

	return v, nil
}
