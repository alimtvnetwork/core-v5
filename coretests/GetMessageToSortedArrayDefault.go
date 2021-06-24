package coretests

// GetMessageToSortedArrayDefault
//
// isPrint: false, isSort: true
func GetMessageToSortedArrayDefault(
	message string,
) []string {
	return GetMessageToSortedArray(
		false,
		true, message)
}
