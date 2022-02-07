package enumimpl

import (
	"gitlab.com/evatix-go/core/converters"
)

type newBasicStringCreator struct{}

// CreateUsingAliasMap
//
// Length : must match stringRanges and actualRangesAnyType
func (it newBasicStringCreator) CreateUsingAliasMap(
	typeName string,
	stringRanges []string,
	aliasingMap map[string]string,
	min, max string,
) *BasicString {
	enumBase := newNumberEnumBase(
		typeName,
		stringRanges,
		stringRanges,
		min,
		max)

	jsonDoubleQuoteNameToValueHashMap := make(
		map[string]string,
		len(stringRanges))
	valueToJsonDoubleQuoteStringBytesHashmap := make(
		map[string][]byte,
		len(stringRanges))

	for i, actualVal := range stringRanges {
		key := stringRanges[i]
		jsonName := toJsonName(key)
		jsonDoubleQuoteNameToValueHashMap[jsonName] = actualVal
		jsonDoubleQuoteNameToValueHashMap[key] = actualVal
		valueToJsonDoubleQuoteStringBytesHashmap[key] = []byte(jsonName)
	}

	if len(aliasingMap) > 0 {
		for aliasName, aliasValue := range aliasingMap {
			aliasJsonName := toJsonName(aliasName)
			jsonDoubleQuoteNameToValueHashMap[aliasName] = aliasValue
			jsonDoubleQuoteNameToValueHashMap[aliasJsonName] = aliasValue
		}
	}

	return &BasicString{
		numberEnumBase: enumBase,
		minVal:         min,
		maxVal:         max,
		jsonDoubleQuoteNameToValueHashMap: *converters.
			StringsToMap(&stringRanges),
		valueToJsonDoubleQuoteStringBytesHashmap: valueToJsonDoubleQuoteStringBytesHashmap,
	}
}
