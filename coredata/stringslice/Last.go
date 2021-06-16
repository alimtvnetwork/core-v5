package stringslice

import "gitlab.com/evatix-go/core/constants"

func Last(slice []string) string {
	return (slice)[len(slice)-constants.One]
}
