package corestr

// LengthOfStringsOfStrings it doesn't not return all individual items length.
func LengthOfStringsOfStrings(stringsOfStringsItems *[][]string) int {
	if stringsOfStringsItems == nil || *stringsOfStringsItems == nil {
		return 0
	}

	return len(*stringsOfStringsItems)
}
