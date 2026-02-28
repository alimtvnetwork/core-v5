package isany

func DefinedBoth(leftAnyItem, rightAnyItem any) (isBothDefined bool) {
	leftNull := Null(leftAnyItem)

	return !leftNull && leftNull == Null(rightAnyItem)
}
