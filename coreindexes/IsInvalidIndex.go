package coreindexes

import "github.com/alimtvnetwork/core/constants"

func IsInvalidIndex(index int) bool {
	return index <= constants.InvalidIndex
}
