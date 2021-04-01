package converters

func StringsToMap(inputArray *[]string) *map[string]bool {
	length := len(*inputArray)
	hashset := make(map[string]bool, length)

	for _, s := range *inputArray {
		hashset[s] = true
	}

	return &hashset
}
