package stringslice

func AnyItemsCloneIf(
	isClone bool,
	additionalCap int,
	slice []any,
) (newSlice []any) {
	if slice == nil && !isClone {
		return []any{}
	}

	if !isClone {
		return slice
	}

	return AnyItemsCloneUsingCap(additionalCap, slice)
}
