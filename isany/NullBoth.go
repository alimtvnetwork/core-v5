package isany

func NullBoth(leftAnyItem, rightAnyItem interface{}) (isBothNull bool) {
	leftNull := Null(leftAnyItem)

	return leftNull && leftNull == Null(rightAnyItem)
}
