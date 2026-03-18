package corejson

func BytesCloneIf(
	isDeepClone bool,
	inputBytes []byte,
) []byte {
	if !isDeepClone || len(inputBytes) == 0 {
		return []byte{}
	}

	newBytes := make([]byte, len(inputBytes))
	copy(newBytes, inputBytes)

	return newBytes
}
