package isany

func DefinedLeftRight(leftAnyItem, rightAnyItem interface{}) (isLeftDefined, isRightDefined bool) {
	return !Null(leftAnyItem), !Null(rightAnyItem)
}
