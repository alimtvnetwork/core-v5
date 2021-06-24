package corestr

func AllIndividualStringsOfStringsLength(stringOfStringsItems *[][]string) int {
	if stringOfStringsItems == nil || *stringOfStringsItems == nil {
		return 0
	}

	length := 0

	for _, stringsItems := range *stringOfStringsItems {
		length += len(stringsItems)
	}

	return length
}
