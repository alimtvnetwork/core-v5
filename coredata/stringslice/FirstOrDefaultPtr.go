package stringslice

import "gitlab.com/evatix-go/core/constants"

func FirstOrDefaultPtr(slice *[]string) string {
	if slice == nil || len(*slice) == 0 {
		return constants.EmptyString
	}

	return (*slice)[0]
}
