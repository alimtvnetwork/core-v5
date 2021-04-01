package converters

func StringToByteWithDefault(
	input string, defaultByte byte,
) (value byte, isSuccess bool) {
	vByte, err := StringToByte(input)

	if err != nil {
		return defaultByte, false
	}

	return vByte, true
}
