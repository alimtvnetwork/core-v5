package corestr

func AllIndividualItemsStringsOfStringsPointerLength(stringsOfStringsItems *[]*[]string) int {
	if stringsOfStringsItems == nil || *stringsOfStringsItems == nil {
		return 0
	}

	length := 0

	for _, stringsItems := range *stringsOfStringsItems {
		if stringsItems == nil {
			continue
		}

		length += len(*stringsItems)
	}

	return length
}
