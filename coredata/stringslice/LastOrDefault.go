package stringslice

import "gitlab.com/evatix-go/core/constants"

func LastOrDefault(slice []string) string {
	if len(slice) == 0 {
		return constants.EmptyString
	}

	return (slice)[len(slice)-constants.One]
}
