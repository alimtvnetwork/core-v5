package corejson

func BytesDeepClone(
	inputBytes []byte,
) []byte {
	if len(inputBytes) == 0 {
		return []byte{}
	}

	newBytes := make([]byte, len(inputBytes))
	copy(newBytes, inputBytes)

	return newBytes
}
