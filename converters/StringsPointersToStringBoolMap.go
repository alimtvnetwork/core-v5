package converters

func StringsPointersToStringBoolMap(inputArray *[]*string) *map[string]bool {
	length := len(*inputArray)
	hashset := make(map[string]bool, length)

	for _, s := range *inputArray {
		sC := *s
		hashset[sC] = true
	}

	return &hashset
}
