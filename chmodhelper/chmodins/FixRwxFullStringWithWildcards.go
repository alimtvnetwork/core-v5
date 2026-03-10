package chmodins

import (
	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/internal/strutilinternal"
)

// FixRwxFullStringWithWildcards can be less than 10 and can be
//
//   - "-rwx" will be "-rwx******"
//   - "-rwxr-x" will be "-rwxr-x***"
//   - "-rwxr-x" will be "-rwxr-x***"
func FixRwxFullStringWithWildcards(rwxFull string) (fixedRwx string) {
	length := len(rwxFull)

	if length == RwxFullLength {
		return rwxFull
	} else if length == constants.Zero {
		return AllWildCardsRwxFullString
	}

	return strutilinternal.MaskLine(
		AllWildCardsRwxFullString,
		rwxFull)
}
