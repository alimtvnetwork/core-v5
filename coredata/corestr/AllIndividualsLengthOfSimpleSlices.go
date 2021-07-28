package corestr

func AllIndividualsLengthOfSimpleSlices(items ...*SimpleSlice) int {
	length := 0

	if items == nil || len(items) == 0 {
		return length
	}

	for _, simpleSlice := range items {
		length += simpleSlice.Length()
	}

	return length
}
