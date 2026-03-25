package isany

func DefinedLeftRight(leftAnyItem, rightAnyItem any) (isLeftDefined, isRightDefined bool) {
	return !Null(leftAnyItem), !Null(rightAnyItem)
}
