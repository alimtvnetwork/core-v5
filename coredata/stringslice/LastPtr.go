package stringslice

import "gitlab.com/evatix-go/core/constants"

func LastPtr(slice *[]string) string {
	return (*slice)[len(*slice)-constants.One]
}
