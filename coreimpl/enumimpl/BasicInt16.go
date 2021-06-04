package enumimpl

type BasicInt16 struct {
	*numberEnumBase
	hashMap        map[string]int16
	minVal, maxVal int16
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

	hashMap := make(map[string]int16, len(actualValueRanges))
	for i, actual := range actualValueRanges {
		key := stringRanges[i]
		hashMap[key] = actual
	}

	return &BasicInt16{
		numberEnumBase: enumBase,
		minVal:         min,
		maxVal:         max,
		hashMap:        hashMap,
	}
}

func (receiver *BasicInt16) Max() int16 {
	return receiver.maxVal
}

func (receiver *BasicInt16) Min() int16 {
	return receiver.minVal
}

func (receiver *BasicInt16) GetValueByString(valueString string) int16 {
	return receiver.hashMap[valueString]
}

func (receiver *BasicInt16) GetStringValue(input int16) string {
	return receiver.StringRanges()[input]
}

func (receiver *BasicInt16) Ranges() []int16 {
	return receiver.actualValueRanges.([]int16)
}

func (receiver *BasicInt16) Hashmap() map[string]int16 {
	return receiver.hashMap
}

func (receiver *BasicInt16) HashmapPtr() *map[string]int16 {
	return &receiver.hashMap
}

func (receiver *BasicInt16) IsValidRange(value int16) bool {
	return value >= receiver.minVal && value <= receiver.maxVal
}
