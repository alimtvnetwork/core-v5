package stringslice

import "github.com/alimtvnetwork/core/constants"

func Last(slice []string) string {
	return (slice)[len(slice)-constants.One]
}
