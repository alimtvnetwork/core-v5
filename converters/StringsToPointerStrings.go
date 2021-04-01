package converters

// StringsToPointerStrings will give empty or converted results array (not nil)
//
// It doesn't copy but points to same string address in the array
//
// Example code : https://play.golang.org/p/_OkY82E2kO9
func StringsToPointerStrings(ptrStrArray *[]string) *[]*string {
	if ptrStrArray == nil || *ptrStrArray == nil {
		var emptyResult []*string

		return &emptyResult
	}

	newArray := make([]*string, len(*ptrStrArray))

	for i := range *ptrStrArray {
		// direct access important here.
		newArray[i] = &(*ptrStrArray)[i]
	}

	return &newArray
}
