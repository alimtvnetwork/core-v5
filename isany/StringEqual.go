package isany

import "github.com/alimtvnetwork/core/internal/strutilinternal"

func StringEqual(
	left, right any,
) bool {
	leftString := strutilinternal.AnyToFieldNameString(left)
	rightString := strutilinternal.AnyToFieldNameString(right)

	return leftString == rightString
}
