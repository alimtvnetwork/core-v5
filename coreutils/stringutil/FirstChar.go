package stringutil

// FirstChar returns first char or 0 if empty
func FirstChar(input string) byte {
	if input == "" {
		return 0
	}

	return input[0]
}

// FirstCharOrDefault gives 0 if nothing present
func FirstCharOrDefault(input string) byte {
	if input == "" {
		return 0
	}

	return input[0]
}

// LastChar panics if not char
func LastChar(input string) byte {
	return input[len(input)-1]
}

func LastCharOrDefault(input string) byte {
	if input == "" {
		return 0
	}

	return LastChar(input)
}
