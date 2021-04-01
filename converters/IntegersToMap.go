package converters

func IntegersToMap(inputArray *[]int) *map[int]bool {
	length := len(*inputArray)
	hashset := make(map[int]bool, length)

	for _, s := range *inputArray {
		hashset[s] = true
	}

	return &hashset
}
