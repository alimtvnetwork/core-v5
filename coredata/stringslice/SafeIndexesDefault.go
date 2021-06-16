package stringslice

import (
	"gitlab.com/evatix-go/core/constants"
)

// SafeIndexesDefault Only indexes which are present values will be included.
//
// Warning : Not found indexes will not be included in the values.
func SafeIndexesDefault(slice []string, indexes ...int) (values []string) {
	length := len(slice)

	if length == 0 {
		return []string{}
	}

	values = make(
		[]string,
		constants.Zero,
		length)

	inputIndex := 0
	lastIndex := length - 1
	for _, index := range indexes {
		if index > lastIndex {
			// don't include
			continue
		}

		values = append(values, slice[index])
		inputIndex++
	}

	return values
}
