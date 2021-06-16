package stringslice

import "gitlab.com/evatix-go/core/constants"

func SafeIndexAt(slice []string, index int) string {
	if len(slice) == 0 {
		return constants.EmptyString
	}

	return slice[index]
}
