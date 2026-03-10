package stringslice

import "github.com/alimtvnetwork/core/constants"

func LastOrDefault(slice []string) string {
	if len(slice) == 0 {
		return constants.EmptyString
	}

	return (slice)[len(slice)-constants.One]
}
