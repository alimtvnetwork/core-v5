package corestr

func CloneSliceIf(
	isClone bool,
	sourceItems ...string,
) []string {
	if len(sourceItems) == 0 {
		return []string{}
	}

	isSkipClone := !isClone

	if isSkipClone {
		return sourceItems
	}

	destinationSlice := make(
		[]string,
		len(sourceItems))
	copy(destinationSlice, sourceItems)

	return destinationSlice
}
