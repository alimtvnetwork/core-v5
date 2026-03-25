package stringslice

func AnyItemsCloneIf(
	isClone bool,
	additionalCap int,
	slice []any,
) (newSlice []any) {
	isSkipClone := !isClone

	if slice == nil && isSkipClone {
		return []any{}
	}

	if isSkipClone {
		return slice
	}

	return AnyItemsCloneUsingCap(additionalCap, slice)
}
