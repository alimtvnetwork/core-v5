package enumimpl

import "gitlab.com/evatix-go/core/constants"

type BasicByte struct {
	*numberEnumBase
	hashMap        map[string]byte
	minVal, maxVal byte
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

	hashMap := make(map[string]byte, len(actualValueRanges))
	for i, b := range actualValueRanges {
		key := stringRanges[i]
		hashMap[key] = b
	}

	return &BasicByte{
		numberEnumBase: enumBase,
		minVal:         min,
		maxVal:         max,
		hashMap:        hashMap,
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

	enumBase := newNumberEnumBase(
		actualValues,
		indexedSliceWithValues,
		min,
		max)

	hashMap := make(map[string]byte, max)
	for i, key := range indexedSliceWithValues {
		hashMap[key] = actualValues[i]
	}

	return &BasicByte{
		numberEnumBase: enumBase,
		minVal:         byte(min),
		maxVal:         byte(max),
		hashMap:        hashMap,
	}
}

func (receiver *BasicByte) Max() byte {
	return receiver.maxVal
}

func (receiver *BasicByte) Min() byte {
	return receiver.minVal
}

func (receiver *BasicByte) GetValueByString(valueString string) byte {
	return receiver.hashMap[valueString]
}

func (receiver *BasicByte) GetStringValue(input byte) string {
	return receiver.StringRanges()[input]
}

func (receiver *BasicByte) Ranges() []byte {
	return receiver.actualValueRanges.([]byte)
}

func (receiver *BasicByte) Hashmap() map[string]byte {
	return receiver.hashMap
}

func (receiver *BasicByte) HashmapPtr() *map[string]byte {
	return &receiver.hashMap
}

func (receiver *BasicByte) IsValidRange(value byte) bool {
	return value >= receiver.minVal && value <= receiver.maxVal
}
