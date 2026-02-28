package isany

func NullLeftRight(
	leftAnyItem,
	rightAnyItem any,
) (
	isLeftNull, isRightNull bool,
) {
	return Null(leftAnyItem),
		Null(rightAnyItem)
}
