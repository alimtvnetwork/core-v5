package anycmp

import (
	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/isany"
)

// Cmp
//
//	Usages quick determination to say Equal, NotEqual or Inconclusive.
//	Don't use any deep reflect or bytes comparison or string comparison.
//
// Steps:
//   - Returns Equal if both are same pointer
//   - Returns Equal if both are nil
//   - Returns NotEqual if both one nil and another is not
//   - Else, returns Inconclusive.
func Cmp(left, right any) corecomparator.Compare {
	if left == right {
		return corecomparator.Equal
	}

	if left == nil && right == nil {
		return corecomparator.Equal
	}

	if left == nil || right == nil {
		return corecomparator.NotEqual
	}

	isLeftNull, isRightNull := isany.NullLeftRight(left, right)
	isBothEqual := isLeftNull == isRightNull
	isBothDifferent := !isBothEqual

	if isLeftNull && isBothEqual {
		// both null
		return corecomparator.Equal
	} else if isBothDifferent && (isLeftNull || isRightNull) {
		// any null but the other is not
		return corecomparator.NotEqual
	}

	return corecomparator.Inconclusive
}
