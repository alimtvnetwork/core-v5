package pagingutil

import (
	"math"
)

// GetPagesSize returns ceiling int for pages possible.
// Returns 0 if eachPageSize is zero or negative.
func GetPagesSize(
	eachPageSize,
	totalLength int,
) int {
	if eachPageSize <= 0 {
		return 0
	}

	pagesPossibleFloat := float64(totalLength) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))

	return pagesPossibleCeiling
}
