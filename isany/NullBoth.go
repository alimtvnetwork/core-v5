package isany

func NullBoth(leftAnyItem, rightAnyItem any) (isBothNull bool) {
	leftNull := Null(leftAnyItem)

	return leftNull && leftNull == Null(rightAnyItem)
}
