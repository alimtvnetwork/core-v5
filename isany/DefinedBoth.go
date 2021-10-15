package isany

func DefinedBoth(leftAnyItem, rightAnyItem interface{}) (isBothDefined bool) {
	leftNull := Null(leftAnyItem)

	return !leftNull && leftNull == Null(rightAnyItem)
}
